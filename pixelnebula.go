package pixelnebula

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/landaiqing/go-pixelnebula/animation"
	"github.com/landaiqing/go-pixelnebula/cache"
	"github.com/landaiqing/go-pixelnebula/converter"
	"github.com/landaiqing/go-pixelnebula/errors"
	"github.com/landaiqing/go-pixelnebula/style"
	"github.com/landaiqing/go-pixelnebula/theme"
)

const (
	hashLength = 12
)

var (
	// 优化正则表达式，使用更高效的模式
	numberRegex = regexp.MustCompile(`[0-9]`)
	// 使用非贪婪模式并优化颜色匹配模式
	colorRegex = regexp.MustCompile(`#([^;]*);`)
	// 使用字节池减少内存分配
	hashBufPool = sync.Pool{
		New: func() interface{} {
			buf := make([]byte, 64)
			return &buf
		},
	}
	// 缓存键计算结果缓存
	keyCache    = make(map[string][2]int)
	keyCacheMux sync.RWMutex
	// 使用对象池减少内存分配
	builderPool = sync.Pool{
		New: func() interface{} {
			return &strings.Builder{}
		},
	}
	mapPool = sync.Pool{
		New: func() interface{} {
			return make(map[string]string, 6)
		},
	}
	keyMapPool = sync.Pool{
		New: func() interface{} {
			return make(map[string][2]int, 6)
		},
	}
)

type PNOptions struct {
	ThemeIndex int // 主题索引，
	StyleIndex int // 风格索引，
}

type PixelNebula struct {
	svgEnd       string
	themeManager *theme.Manager
	styleManager *style.Manager
	animManager  *animation.Manager
	cache        *cache.PNCache
	hasher       hash.Hash
	options      *PNOptions
	width        int
	height       int
	imgData      []byte
}

// NewPixelNebula 创建一个PixelNebula实例
func NewPixelNebula() *PixelNebula {
	return &PixelNebula{
		svgEnd:       "</svg>",
		themeManager: theme.NewThemeManager(),
		styleManager: style.NewShapeManager(),
		animManager:  animation.NewAnimationManager(),
		hasher:       sha256.New(),
		options:      &PNOptions{ThemeIndex: -1, StyleIndex: -1}, // 初始化为 -1 表示未设置
		width:        231,
		height:       231,
	}
}

// getSvgStart 根据当前宽高生成SVG开始标签
func (pn *PixelNebula) getSvgStart() string {
	return fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 %d %d\">", pn.width, pn.height)
}

// WithTheme 设置固定主题
func (pn *PixelNebula) WithTheme(themeIndex int) *PixelNebula {
	// 如果已设置 style,则验证主题索引是否有效
	if styleIndex := pn.options.StyleIndex; styleIndex >= 0 {
		// 获取该风格下的主题数量
		themeCount := pn.themeManager.ThemeCount(styleIndex)
		if themeIndex < 0 || themeIndex >= themeCount {
			log.Printf("pixelnebula: theme index range is:[0, %d), but got %d", themeCount, themeIndex)
			panic(errors.ErrInvalidTheme)
		}
	}
	pn.options.ThemeIndex = themeIndex
	return pn
}

// WithStyle 设置固定风格
func (pn *PixelNebula) WithStyle(style style.StyleType) *PixelNebula {
	styleIndex, err := pn.styleManager.GetStyleIndex(style)
	if err != nil {
		panic(err)
	}
	pn.options.StyleIndex = styleIndex
	return pn
}

// WithSize 设置尺寸
func (pn *PixelNebula) WithSize(width, height int) *PixelNebula {
	pn.width = width
	pn.height = height
	return pn
}

// WithCustomizeTheme 设置自定义主题
func (pn *PixelNebula) WithCustomizeTheme(theme []theme.Theme) *PixelNebula {
	pn.themeManager.CustomizeTheme(theme)
	return pn
}

// WithCustomizeStyle 设置自定义风格
func (pn *PixelNebula) WithCustomizeStyle(style []style.StyleSet) *PixelNebula {
	pn.styleManager.CustomizeStyle(style)
	return pn
}

// hashToNum 将哈希字符串转换为数字
func (pn *PixelNebula) hashToNum(hash []string) int64 {
	if len(hash) == 0 {
		return 0
	}

	// 使用位运算直接计算哈希值，减少内存分配和计算复杂度
	var result int64
	for _, h := range hash {
		// 优化：直接处理字符，避免ParseInt的开销
		for i := 0; i < len(h); i++ {
			if h[i] >= '0' && h[i] <= '9' {
				// 使用位运算进行计算: result = result*10 + (h[i] - '0')
				result = (result << 3) + (result << 1) + int64(h[i]-'0')
			}
		}
	}

	// 确保结果为正数
	if result < 0 {
		result = -result
	}

	return result
}

// getCacheKey 生成缓存键的哈希表示
func (pn *PixelNebula) getCacheKey(id string, hash []string, index int, opts *PNOptions) string {
	// 构造一个唯一的键字符串
	var key string
	if opts != nil && opts.StyleIndex >= 0 && opts.ThemeIndex >= 0 {
		key = fmt.Sprintf("%s_%d_%d_%d", id, index, opts.StyleIndex, opts.ThemeIndex)
	} else if len(hash) > 0 {
		key = fmt.Sprintf("%s_%d_%s", id, index, strings.Join(hash, ""))
	} else {
		key = fmt.Sprintf("%s_%d", id, index)
	}
	return key
}

// calcKey 计算主题和部分的键值
func (pn *PixelNebula) calcKey(hash []string, opts *PNOptions) [2]int {
	// 检查是否使用固定值
	if opts != nil && opts.StyleIndex >= 0 && opts.ThemeIndex >= 0 {
		return [2]int{opts.StyleIndex, opts.ThemeIndex}
	}

	// 尝试从缓存中获取结果
	cacheKey := strings.Join(hash, "")
	keyCacheMux.RLock()
	if result, ok := keyCache[cacheKey]; ok {
		keyCacheMux.RUnlock()
		return result
	}
	keyCacheMux.RUnlock()

	// 计算哈希值
	hashNum := pn.hashToNum(hash)

	// 获取可用的风格数量
	styleCount := pn.themeManager.StyleCount()
	if styleCount == 0 {
		return [2]int{0, 0}
	}

	// 使用位运算优化取模操作
	styleIndex := int(hashNum % int64(styleCount))
	if styleIndex < 0 {
		styleIndex = -styleIndex
	}

	// 获取该风格下的主题数量
	themeCount := pn.themeManager.ThemeCount(styleIndex)
	if themeCount == 0 {
		return [2]int{styleIndex, 0}
	}

	// 使用哈希值的不同部分计算主题索引
	themeIndex := int(hashNum % int64(themeCount))
	if themeIndex < 0 {
		themeIndex = -themeIndex
	}

	// 将结果存入缓存
	result := [2]int{styleIndex, themeIndex}
	keyCacheMux.Lock()
	keyCache[cacheKey] = result
	keyCacheMux.Unlock()

	return result
}

// WithCache 设置缓存选项
func (pn *PixelNebula) WithCache(options cache.CacheOptions) *PixelNebula {
	pn.cache = cache.NewCache(options)
	// 确保启动监控器
	if pn.cache != nil && options.Monitoring.Enabled && pn.cache.GetMonitor() == nil {
		pn.cache.Monitor = cache.NewMonitor(pn.cache, options.Monitoring)
		pn.cache.Monitor.Start()
	}
	return pn
}

// WithDefaultCache 设置默认缓存选项
func (pn *PixelNebula) WithDefaultCache() *PixelNebula {
	pn.cache = cache.NewDefaultCache()
	// 确保启动监控器
	if pn.cache != nil && pn.cache.GetOptions().Monitoring.Enabled && pn.cache.GetMonitor() == nil {
		pn.cache.Monitor = cache.NewMonitor(pn.cache, pn.cache.GetOptions().Monitoring)
		pn.cache.Monitor.Start()
	}
	return pn
}

// WithCompression 设置压缩选项
func (pn *PixelNebula) WithCompression(options cache.CompressOptions) *PixelNebula {
	if pn.cache != nil {
		cacheOptions := pn.cache.GetOptions()
		cacheOptions.Compression = options
		pn.cache.UpdateOptions(cacheOptions)
	}
	return pn
}

// WithMonitoring 设置监控选项
func (pn *PixelNebula) WithMonitoring(options cache.MonitorOptions) *PixelNebula {
	if pn.cache != nil {
		cacheOptions := pn.cache.GetOptions()
		cacheOptions.Monitoring = options
		pn.cache.UpdateOptions(cacheOptions)

		// 如果启用了监控但监控器尚未创建，则创建并启动监控器
		if options.Enabled && pn.cache.Monitor == nil {
			pn.cache.Monitor = cache.NewMonitor(pn.cache, options)
			pn.cache.Monitor.Start()
		}
	}
	return pn
}

// WithAnimation 添加动画效果
func (pn *PixelNebula) WithAnimation(animation animation.Animation) *PixelNebula {
	pn.animManager.AddAnimation(animation)
	return pn
}

// WithRotateAnimation 添加旋转动画
func (pn *PixelNebula) WithRotateAnimation(targetID string, fromAngle, toAngle float64, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewRotateAnimation(targetID, fromAngle, toAngle, duration, repeatCount)
	pn.animManager.AddAnimation(anim)
	return pn
}

// WithGradientAnimation 添加渐变动画
func (pn *PixelNebula) WithGradientAnimation(targetID string, colors []string, duration float64, repeatCount int, animate bool) *PixelNebula {
	anim := animation.NewGradientAnimation(targetID, colors, duration, repeatCount, animate)
	pn.animManager.AddAnimation(anim)
	return pn
}

// WithTransformAnimation 添加变换动画
func (pn *PixelNebula) WithTransformAnimation(targetID string, transformType string, from, to string, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewTransformAnimation(targetID, transformType, from, to, duration, repeatCount)
	pn.animManager.AddAnimation(anim)
	return pn
}

// WithFadeAnimation 添加淡入淡出动画
func (pn *PixelNebula) WithFadeAnimation(targetID string, from, to string, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewFadeAnimation(targetID, from, to, duration, repeatCount)
	pn.animManager.AddAnimation(anim)
	return pn
}

// WithPathAnimation 添加路径动画
func (pn *PixelNebula) WithPathAnimation(targetID string, path string, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewPathAnimation(targetID, path, duration, repeatCount)
	pn.animManager.AddAnimation(anim)
	return pn
}

// WithPathAnimationRotate 添加带旋转的路径动画
func (pn *PixelNebula) WithPathAnimationRotate(targetID string, path string, rotate string, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewPathAnimation(targetID, path, duration, repeatCount)
	anim.WithRotate(rotate)
	pn.animManager.AddAnimation(anim)
	return pn
}

// WithColorAnimation 添加颜色变换动画
func (pn *PixelNebula) WithColorAnimation(targetID string, property string, fromColor, toColor string, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewColorAnimation(targetID, property, fromColor, toColor, duration, repeatCount)
	pn.animManager.AddAnimation(anim)
	return pn
}

// WithBounceAnimation 添加弹跳动画
func (pn *PixelNebula) WithBounceAnimation(targetID string, property string, from, to string, bounceCount int, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewBounceAnimation(targetID, property, from, to, bounceCount, duration, repeatCount)
	pn.animManager.AddAnimation(anim)
	return pn
}

// WithWaveAnimation 添加波浪动画
func (pn *PixelNebula) WithWaveAnimation(targetID string, amplitude, frequency float64, direction string, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewWaveAnimation(targetID, amplitude, frequency, direction, duration, repeatCount)
	pn.animManager.AddAnimation(anim)
	return pn
}

// WithBlinkAnimation 添加闪烁动画
func (pn *PixelNebula) WithBlinkAnimation(targetID string, minOpacity, maxOpacity float64, blinkCount int, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewBlinkAnimation(targetID, minOpacity, maxOpacity, blinkCount, duration, repeatCount)
	pn.animManager.AddAnimation(anim)
	return pn
}

// SVGBuilder 用于处理SVG生成后的链式操作
type SVGBuilder struct {
	pn         *PixelNebula
	svg        string
	id         string
	sansEnv    bool
	themeIndex int
	styleIndex int
	width      int
	height     int
	hasError   error
}

// Generate 现在返回 SVGBuilder
func (pn *PixelNebula) Generate(id string, sansEnv bool) *SVGBuilder {
	return &SVGBuilder{
		pn:         pn,
		id:         id,
		sansEnv:    sansEnv,
		width:      pn.width,
		height:     pn.height,
		themeIndex: pn.options.ThemeIndex,
		styleIndex: pn.options.StyleIndex,
	}
}

// SetTheme 设置主题
func (sb *SVGBuilder) SetTheme(theme int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	themeCount := sb.pn.themeManager.ThemeCount(sb.styleIndex)
	if theme < 0 || theme >= themeCount {
		log.Printf("pixelnebula: theme index range is:[0, %d), but got %d", themeCount, theme)
		sb.hasError = errors.ErrInvalidTheme
		return sb
	}
	sb.themeIndex = theme
	return sb
}

// SetStyle 设置风格
// 注意：当使用WithCustomizeStyle设置自定义风格后，此方法将无法正常工作，应使用SetStyleByIndex代替
func (sb *SVGBuilder) SetStyle(style style.StyleType) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	index, err := sb.pn.styleManager.GetStyleIndex(style)
	if err != nil {
		sb.hasError = err
		return sb
	}
	sb.styleIndex = index
	return sb
}

// SetStyleByIndex 设置风格索引
// 此方法可用于设置自定义风格的索引，特别是在使用WithCustomizeStyle后
func (sb *SVGBuilder) SetStyleByIndex(index int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	themeCount := sb.pn.themeManager.StyleCount()
	if index < 0 || index >= themeCount {
		log.Printf("pixelnebula: style index range is:[0, %d), but got %d", themeCount, index)
		sb.hasError = errors.ErrInvalidStyleName
		return sb
	}
	sb.styleIndex = index
	return sb
}

// SetSize 设置尺寸
func (sb *SVGBuilder) SetSize(width, height int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	sb.width = width
	sb.height = height
	return sb
}

// SetAnimation 添加动画效果
func (sb *SVGBuilder) SetAnimation(anim animation.Animation) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	sb.pn.animManager.AddAnimation(anim)
	return sb
}

// SetRotateAnimation 添加旋转动画
func (sb *SVGBuilder) SetRotateAnimation(targetID string, fromAngle, toAngle float64, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewRotateAnimation(targetID, fromAngle, toAngle, duration, repeatCount)
	sb.pn.animManager.AddAnimation(anim)
	return sb
}

// SetGradientAnimation 添加渐变动画
func (sb *SVGBuilder) SetGradientAnimation(targetID string, colors []string, duration float64, repeatCount int, animate bool) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewGradientAnimation(targetID, colors, duration, repeatCount, animate)
	sb.pn.animManager.AddAnimation(anim)
	return sb
}

// SetTransformAnimation 添加变换动画
func (sb *SVGBuilder) SetTransformAnimation(targetID string, transformType string, from, to string, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewTransformAnimation(targetID, transformType, from, to, duration, repeatCount)
	sb.pn.animManager.AddAnimation(anim)
	return sb
}

// SetFadeAnimation 添加淡入淡出动画
func (sb *SVGBuilder) SetFadeAnimation(targetID string, from, to string, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewFadeAnimation(targetID, from, to, duration, repeatCount)
	sb.pn.animManager.AddAnimation(anim)
	return sb
}

// SetPathAnimation 添加路径动画
func (sb *SVGBuilder) SetPathAnimation(targetID string, path string, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewPathAnimation(targetID, path, duration, repeatCount)
	sb.pn.animManager.AddAnimation(anim)
	return sb
}

// SetPathAnimationRotate 添加带旋转的路径动画
func (sb *SVGBuilder) SetPathAnimationRotate(targetID string, path string, rotate string, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewPathAnimation(targetID, path, duration, repeatCount)
	anim.WithRotate(rotate)
	sb.pn.animManager.AddAnimation(anim)
	return sb
}

// SetColorAnimation 添加颜色变换动画
func (sb *SVGBuilder) SetColorAnimation(targetID string, property string, fromColor, toColor string, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewColorAnimation(targetID, property, fromColor, toColor, duration, repeatCount)
	sb.pn.animManager.AddAnimation(anim)
	return sb
}

// SetBounceAnimation 添加弹跳动画
func (sb *SVGBuilder) SetBounceAnimation(targetID string, property string, from, to string, bounceCount int, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewBounceAnimation(targetID, property, from, to, bounceCount, duration, repeatCount)
	sb.pn.animManager.AddAnimation(anim)
	return sb
}

// SetWaveAnimation 添加波浪动画
func (sb *SVGBuilder) SetWaveAnimation(targetID string, amplitude, frequency float64, direction string, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewWaveAnimation(targetID, amplitude, frequency, direction, duration, repeatCount)
	sb.pn.animManager.AddAnimation(anim)
	return sb
}

// SetBlinkAnimation 添加闪烁动画
func (sb *SVGBuilder) SetBlinkAnimation(targetID string, minOpacity, maxOpacity float64, blinkCount int, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewBlinkAnimation(targetID, minOpacity, maxOpacity, blinkCount, duration, repeatCount)
	sb.pn.animManager.AddAnimation(anim)
	return sb
}

// Build 生成最终的SVG
func (sb *SVGBuilder) Build() *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}

	opts := &PNOptions{
		ThemeIndex: sb.themeIndex,
		StyleIndex: sb.styleIndex,
	}

	svg, err := sb.pn.generateSVG(sb.id, sb.sansEnv, opts)
	if err != nil {
		sb.hasError = err
		return sb
	}

	sb.svg = svg
	sb.pn.imgData = []byte(svg)
	sb.pn.width = sb.width
	sb.pn.height = sb.height
	return sb
}

// ToSVG 获取SVG字符串
func (sb *SVGBuilder) ToSVG() (string, error) {
	if sb.svg == "" {
		sb = sb.Build()
	}
	if sb.hasError != nil {
		return "", sb.hasError
	}
	return sb.svg, nil
}

// ToBase64 获取Base64编码的SVG字符串 注意：这个设置宽高无效
func (sb *SVGBuilder) ToBase64() (string, error) {
	if sb.svg == "" {
		sb = sb.Build()
	}
	if sb.hasError != nil {
		return "", sb.hasError
	}
	conv := converter.NewSVGConverter([]byte(sb.svg), sb.width, sb.height)
	return conv.ToBase64()
}

// ToFile 将SVG代码保存到文件
func (sb *SVGBuilder) ToFile(filePath string) error {
	if sb.svg == "" {
		sb = sb.Build()
	}
	if sb.hasError != nil {
		return sb.hasError
	}
	return os.WriteFile(filePath, []byte(sb.svg), 0644)
}

// 将原来的 GenerateSVG 重命名为 generateSVG，作为内部方法
func (pn *PixelNebula) generateSVG(id string, sansEnv bool, opts *PNOptions) (svg string, err error) {
	if opts == nil {
		opts = pn.options
	}
	// 验证参数
	if id == "" {
		return "", errors.ErrAvatarIDRequired
	}

	// 如果启用了缓存，先尝试从缓存获取
	if pn.cache != nil {
		cacheKey := cache.CacheKey{
			Id:      id,
			SansEnv: sansEnv,
		}

		if opts != nil {
			cacheKey.Theme = opts.ThemeIndex
			cacheKey.Part = opts.StyleIndex
		}

		if cachedSVG, found := pn.cache.Get(cacheKey); found {
			return cachedSVG, nil
		}
	}

	// 使用对象池获取缓冲区
	hashBuf := hashBufPool.Get().(*[]byte)
	defer hashBufPool.Put(hashBuf)

	// 计算avatarId的哈希值 - 优化版本
	pn.hasher.Reset()
	pn.hasher.Write([]byte(id))
	sum := pn.hasher.Sum((*hashBuf)[:0])
	s := hex.EncodeToString(sum)
	hashStr := numberRegex.FindAllString(s, -1)
	if len(hashStr) < hashLength {
		return "", errors.ErrInsufficientHash
	}
	hashStr = hashStr[0:hashLength]

	// 从对象池获取映射
	p := keyMapPool.Get().(map[string][2]int)
	defer func() {
		// 清空并归还对象池
		for k := range p {
			delete(p, k)
		}
		keyMapPool.Put(p)
	}()

	// 计算各部分的键值
	p[string(style.TypeEnv)] = pn.calcKey(hashStr[:2], opts)
	p[string(style.TypeClo)] = pn.calcKey(hashStr[2:4], opts)
	p[string(style.TypeHead)] = pn.calcKey(hashStr[4:6], opts)
	p[string(style.TypeMouth)] = pn.calcKey(hashStr[6:8], opts)
	p[string(style.TypeEyes)] = pn.calcKey(hashStr[8:10], opts)
	p[string(style.TypeTop)] = pn.calcKey(hashStr[10:], opts)

	// 获取结果映射
	final := mapPool.Get().(map[string]string)
	defer func() {
		// 清空并归还对象池
		for k := range final {
			delete(final, k)
		}
		mapPool.Put(final)
	}()

	// 获取并处理各部分的SVG
	for k, v := range p {
		// 获取主题颜色
		themePart, err := pn.themeManager.GetTheme(v[0], v[1])
		if err != nil {
			return "", err
		}

		colors, ok := themePart[k]
		if !ok {
			return "", errors.ErrInvalidColor
		}

		// 获取形状SVG
		shapeType := style.ShapeType(k)
		svgPart, err := pn.styleManager.GetShape(v[0], shapeType)
		if err != nil {
			return "", err
		}

		match := colorRegex.FindAllStringSubmatch(svgPart, -1)

		// 从对象池获取Builder
		sb := builderPool.Get().(*strings.Builder)
		sb.Reset()
		sb.Grow(len(svgPart) + 50) // 预分配足够的容量

		lastIndex := 0
		for i, m := range match {
			if i < len(colors) {
				// 找到完整匹配的位置
				index := strings.Index(svgPart[lastIndex:], m[0]) + lastIndex
				// 添加匹配前的部分
				sb.WriteString(svgPart[lastIndex:index])
				// 添加替换后的颜色
				// 检查颜色值是否已经包含#前缀
				if strings.HasPrefix(colors[i], "#") {
					sb.WriteString(colors[i])
				} else {
					sb.WriteString("#")
					sb.WriteString(colors[i])
				}
				sb.WriteString(";")
				// 更新lastIndex
				lastIndex = index + len(m[0])
			}
		}
		// 添加剩余部分
		sb.WriteString(svgPart[lastIndex:])

		final[k] = sb.String()

		// 归还Builder到对象池
		builderPool.Put(sb)
	}

	// 使用对象池获取主Builder来构建最终SVG
	builder := builderPool.Get().(*strings.Builder)
	builder.Reset()
	// 预估SVG大小，避免多次内存分配
	builder.Grow(2048) // 2KB 应该足够容纳大多数SVG

	// 添加SVG开始标签
	builder.WriteString(pn.getSvgStart())

	// 获取动画定义
	animations := pn.animManager.GenerateSVGAnimations()
	if animations != "" {
		builder.WriteString(animations)
	}

	// 构建和处理旋转动画 - 使用对象池
	rotateAnimations := make(map[string]bool)
	rotateAnimationSVGs := make(map[string]string)

	// 收集旋转动画
	for _, anim := range pn.animManager.GetAnimations() {
		if rotateAnim, ok := anim.(*animation.RotateAnimation); ok {
			targetID := anim.GetTargetID()
			rotateAnimations[targetID] = true

			// 提取animateTransform部分
			svgCode := rotateAnim.GenerateSVG()
			if start := strings.Index(svgCode, "<animateTransform"); start != -1 {
				if end := strings.Index(svgCode[start:], "/>"); end != -1 {
					rotateAnimationSVGs[targetID] = svgCode[start : start+end+2]
				}
			}
		}
	}

	// 处理环境
	if !sansEnv {
		if _, hasRotate := rotateAnimations["env"]; hasRotate {
			builder.WriteString("<g style=\"transform-box: fill-box; transform-origin: center;\">\n")
			builder.WriteString(final["env"])
			if animSVG, ok := rotateAnimationSVGs["env"]; ok {
				builder.WriteString(animSVG)
			}
			builder.WriteString("</g>\n")
		} else {
			builder.WriteString(final["env"])
		}
	}

	// 处理其他元素
	elements := []string{"head", "clo", "top", "eyes", "mouth"}
	for _, elem := range elements {
		if _, hasRotate := rotateAnimations[elem]; hasRotate {
			builder.WriteString("<g style=\"transform-box: fill-box; transform-origin: center;\">\n")
			builder.WriteString(final[elem])
			if animSVG, ok := rotateAnimationSVGs[elem]; ok {
				builder.WriteString(animSVG)
			}
			builder.WriteString("</g>\n")
		} else {
			builder.WriteString(final[elem])
		}
	}

	builder.WriteString(pn.svgEnd)
	svg = builder.String()

	// 将生成的SVG存储到实例中
	pn.imgData = []byte(svg)

	// 归还Builder到对象池
	builderPool.Put(builder)

	// 如果启用了缓存，将结果存入缓存
	if pn.cache != nil {
		cacheKey := cache.CacheKey{
			Id:      id,
			SansEnv: sansEnv,
		}

		if opts != nil {
			cacheKey.Theme = opts.ThemeIndex
			cacheKey.Part = opts.StyleIndex
		}

		pn.cache.Set(cacheKey, svg)
	}

	return svg, nil
}

// GetCacheStats 获取缓存统计信息
func (pn *PixelNebula) GetCacheStats() (size, hits, misses int, hitRate float64, enabled bool, maxSize int, expiration time.Duration, evictionType string) {
	if pn.cache == nil {
		log.Println("pixelnebula: 缓存未初始化，请先调用WithCache或WithDefaultCache")
		return 0, 0, 0, 0, false, 0, 0, ""
	}

	hits, misses, hitRate = pn.cache.Stats()
	options := pn.cache.GetOptions()

	return pn.cache.Size(), hits, misses, hitRate, options.Enabled, options.Size, options.Expiration, options.EvictionType
}

// CacheItemInfo 缓存项信息结构体
type CacheItemInfo struct {
	Key          cache.CacheKey
	SVG          string
	Compressed   []byte
	IsCompressed bool
	CreatedAt    time.Time
	LastUsed     time.Time
}

// GetCacheItems 获取所有缓存项
func (pn *PixelNebula) GetCacheItems() []CacheItemInfo {
	result := []CacheItemInfo{}

	if pn.cache == nil {
		log.Println("pixelnebula: 缓存未初始化，请先调用WithCache或WithDefaultCache")
		return result
	}

	// 获取内部缓存项
	items := pn.cache.GetAllItems()
	if len(items) == 0 {
		log.Println("pixelnebula: 缓存中没有数据，请先生成一些SVG")
	}

	for key, item := range items {
		cacheItem := CacheItemInfo{
			Key:          key,
			SVG:          item.SVG,
			Compressed:   item.Compressed,
			IsCompressed: item.IsCompressed,
			CreatedAt:    item.CreatedAt,
			LastUsed:     item.LastUsed,
		}

		result = append(result, cacheItem)
	}

	return result
}

// MonitorSampleInfo 监控样本信息
type MonitorSampleInfo struct {
	Timestamp   time.Time
	Size        int
	Hits        int
	Misses      int
	HitRate     float64
	MemoryUsage int64
}

// GetMonitorStats 获取监控统计信息
func (pn *PixelNebula) GetMonitorStats() (enabled bool, sampleInterval, adjustInterval time.Duration,
	targetHitRate float64, lastAdjusted time.Time, samples []MonitorSampleInfo) {

	if pn.cache == nil {
		log.Println("pixelnebula: 缓存未初始化，请先调用WithCache或WithDefaultCache")
		return false, 0, 0, 0, time.Time{}, nil
	}

	options := pn.cache.GetOptions().Monitoring

	// 确保监控器已启用并初始化
	if !options.Enabled {
		log.Println("pixelnebula: 监控未启用，请设置Monitoring.Enabled=true")
		return options.Enabled, options.SampleInterval, options.AdjustInterval, options.TargetHitRate, time.Time{}, nil
	}

	if pn.cache.GetMonitor() == nil {
		log.Println("pixelnebula: 监控器未初始化，正在初始化...")
		pn.cache.Monitor = cache.NewMonitor(pn.cache, options)
		pn.cache.Monitor.Start()
	}

	monitor := pn.cache.GetMonitor()
	stats := monitor.GetStats()
	sampleHistory := monitor.GetSampleHistory()

	if len(sampleHistory) == 0 {
		log.Println("pixelnebula: 监控样本为空，请等待采样完成")
	}

	// 转换样本历史
	samplesInfo := make([]MonitorSampleInfo, 0, len(sampleHistory))
	for _, sample := range sampleHistory {
		sampleInfo := MonitorSampleInfo{
			Timestamp:   sample.LastAdjusted,
			Size:        sample.Size,
			Hits:        sample.Hits,
			Misses:      sample.Misses,
			HitRate:     sample.HitRate,
			MemoryUsage: sample.MemoryUsage,
		}
		samplesInfo = append(samplesInfo, sampleInfo)
	}

	return options.Enabled, options.SampleInterval, options.AdjustInterval,
		options.TargetHitRate, stats.LastAdjusted, samplesInfo
}

// DeleteCacheItem 删除指定的缓存项
func (pn *PixelNebula) DeleteCacheItem(key string) bool {
	if pn.cache == nil {
		log.Println("pixelnebula: 缓存未初始化，请先调用WithCache或WithDefaultCache")
		return false
	}

	// 解析key字符串，格式为"id_sansEnv_theme_part"
	parts := strings.Split(key, "_")
	if len(parts) < 4 {
		log.Printf("pixelnebula: 无效的key格式: %s，应为'id_sansEnv_theme_part'", key)
		return false
	}

	id := parts[0]
	sansEnv := parts[1] == "true"

	themeItem, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Printf("pixelnebula: 解析theme失败: %v", err)
		return false
	}

	part, err := strconv.Atoi(parts[3])
	if err != nil {
		log.Printf("pixelnebula: 解析part失败: %v", err)
		return false
	}

	// 构造CacheKey
	cacheKey := cache.CacheKey{
		Id:      id,
		SansEnv: sansEnv,
		Theme:   themeItem,
		Part:    part,
	}

	result := pn.cache.DeleteItem(cacheKey)
	if !result {
		log.Printf("pixelnebula: 未找到缓存项: %s", key)
	}
	return result
}

// ClearCache 清空缓存
func (pn *PixelNebula) ClearCache() {
	if pn.cache == nil {
		log.Println("pixelnebula: 缓存未初始化，请先调用WithCache或WithDefaultCache")
		return
	}

	pn.cache.Clear()
	log.Println("pixelnebula: 缓存已清空")
}
