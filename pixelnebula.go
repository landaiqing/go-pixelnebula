package pixelnebula

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
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
	// 使用分片锁减少锁竞争
	keyCacheShards    = 16 // 分片数量
	keyCacheLocks     = make([]sync.RWMutex, keyCacheShards)
	keyCacheShardData = make([]map[string][2]int, keyCacheShards)
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

// init 初始化分片缓存
func init() {
	// 初始化分片缓存
	for i := 0; i < keyCacheShards; i++ {
		keyCacheShardData[i] = make(map[string][2]int)
	}
}

// 计算字符串哈希获取分片索引
func getShardIndex(key string) int {
	var hashKey uint32
	for i := 0; i < len(key); i++ {
		hashKey = hashKey*31 + uint32(key[i])
	}
	return int(hashKey % uint32(keyCacheShards))
}

type PNOptions struct {
	ThemeIndex      int  // 主题索引
	StyleIndex      int  // 风格索引
	ParallelRender  bool // 是否启用并行渲染
	ConcurrencyPool int  // 并发池大小，默认为CPU核心数
}

type PixelNebula struct {
	SvgEnd       string
	ThemeManager *theme.Manager
	StyleManager *style.Manager
	AnimManager  *animation.Manager
	Cache        *cache.PNCache
	Hasher       hash.Hash
	Options      *PNOptions
	Width        int
	Height       int
	ImgData      []byte
}

// NewPixelNebula 创建一个PixelNebula实例
func NewPixelNebula() *PixelNebula {
	return &PixelNebula{
		SvgEnd:       "</svg>",
		ThemeManager: theme.NewThemeManager(),
		StyleManager: style.NewShapeManager(),
		AnimManager:  animation.NewAnimationManager(),
		Hasher:       sha256.New(),
		Options:      &PNOptions{ThemeIndex: -1, StyleIndex: -1, ParallelRender: false, ConcurrencyPool: runtime.NumCPU()}, // 初始化为 -1 表示未设置
		Width:        231,
		Height:       231,
	}
}

// getSvgStart 根据当前宽高生成SVG开始标签
func (pn *PixelNebula) getSvgStart() string {
	return fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 %d %d\">", pn.Width, pn.Height)
}

// WithTheme 设置固定主题
func (pn *PixelNebula) WithTheme(themeIndex int) *PixelNebula {
	// 如果已设置 style,则验证主题索引是否有效
	if styleIndex := pn.Options.StyleIndex; styleIndex >= 0 {
		// 获取该风格下的主题数量
		themeCount := pn.ThemeManager.ThemeCount(styleIndex)
		if themeIndex < 0 || themeIndex >= themeCount {
			log.Printf("pixelnebula: theme index range is:[0, %d), but got %d", themeCount, themeIndex)
			panic(errors.ErrInvalidTheme)
		}
	}
	pn.Options.ThemeIndex = themeIndex
	return pn
}

// WithStyle 设置固定风格
func (pn *PixelNebula) WithStyle(style style.StyleType) *PixelNebula {
	styleIndex, err := pn.StyleManager.GetStyleIndex(style)
	if err != nil {
		panic(err)
	}
	pn.Options.StyleIndex = styleIndex
	return pn
}

// WithSize 设置尺寸
func (pn *PixelNebula) WithSize(width, height int) *PixelNebula {
	pn.Width = width
	pn.Height = height
	return pn
}

// WithCustomizeTheme 设置自定义主题
func (pn *PixelNebula) WithCustomizeTheme(theme []theme.Theme) *PixelNebula {
	pn.ThemeManager.CustomizeTheme(theme)
	return pn
}

// WithCustomizeStyle 设置自定义风格
func (pn *PixelNebula) WithCustomizeStyle(style []style.StyleSet) *PixelNebula {
	pn.StyleManager.CustomizeStyle(style)
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

	// 计算缓存键
	cacheKey := strings.Join(hash, "")

	// 计算分片索引
	shardIndex := getShardIndex(cacheKey)

	// 尝试从缓存中获取结果，使用读锁
	keyCacheLocks[shardIndex].RLock()
	if result, ok := keyCacheShardData[shardIndex][cacheKey]; ok {
		keyCacheLocks[shardIndex].RUnlock()
		return result
	}
	keyCacheLocks[shardIndex].RUnlock()

	// 计算哈希值
	hashNum := pn.hashToNum(hash)

	// 获取可用的风格数量
	styleCount := pn.ThemeManager.StyleCount()
	if styleCount == 0 {
		return [2]int{0, 0}
	}

	// 使用位运算优化取模操作
	styleIndex := int(hashNum % int64(styleCount))
	if styleIndex < 0 {
		styleIndex = -styleIndex
	}

	// 获取该风格下的主题数量
	themeCount := pn.ThemeManager.ThemeCount(styleIndex)
	if themeCount == 0 {
		return [2]int{styleIndex, 0}
	}

	// 使用哈希值的不同部分计算主题索引
	themeIndex := int(hashNum % int64(themeCount))
	if themeIndex < 0 {
		themeIndex = -themeIndex
	}

	// 将结果存入缓存，使用写锁
	result := [2]int{styleIndex, themeIndex}
	keyCacheLocks[shardIndex].Lock()
	keyCacheShardData[shardIndex][cacheKey] = result
	keyCacheLocks[shardIndex].Unlock()

	return result
}

// WithCache 设置缓存选项
func (pn *PixelNebula) WithCache(options cache.CacheOptions) *PixelNebula {
	pn.Cache = cache.NewCache(options)
	// 确保启动监控器
	if pn.Cache != nil && options.Monitoring.Enabled && pn.Cache.GetMonitor() == nil {
		pn.Cache.Monitor = cache.NewMonitor(pn.Cache, options.Monitoring)
		pn.Cache.Monitor.Start()
	}
	return pn
}

// WithDefaultCache 设置默认缓存选项
func (pn *PixelNebula) WithDefaultCache() *PixelNebula {
	pn.Cache = cache.NewDefaultCache()
	// 确保启动监控器
	if pn.Cache != nil && pn.Cache.GetOptions().Monitoring.Enabled && pn.Cache.GetMonitor() == nil {
		pn.Cache.Monitor = cache.NewMonitor(pn.Cache, pn.Cache.GetOptions().Monitoring)
		pn.Cache.Monitor.Start()
	}
	return pn
}

// WithCompression 设置压缩选项
func (pn *PixelNebula) WithCompression(options cache.CompressOptions) *PixelNebula {
	if pn.Cache != nil {
		cacheOptions := pn.Cache.GetOptions()
		cacheOptions.Compression = options
		pn.Cache.UpdateOptions(cacheOptions)
	}
	return pn
}

// WithMonitoring 设置监控选项
func (pn *PixelNebula) WithMonitoring(options cache.MonitorOptions) *PixelNebula {
	if pn.Cache != nil {
		cacheOptions := pn.Cache.GetOptions()
		cacheOptions.Monitoring = options
		pn.Cache.UpdateOptions(cacheOptions)

		// 如果启用了监控但监控器尚未创建，则创建并启动监控器
		if options.Enabled && pn.Cache.Monitor == nil {
			pn.Cache.Monitor = cache.NewMonitor(pn.Cache, options)
			pn.Cache.Monitor.Start()
		}
	}
	return pn
}

// WithAnimation 添加动画效果
func (pn *PixelNebula) WithAnimation(animation animation.Animation) *PixelNebula {
	pn.AnimManager.AddAnimation(animation)
	return pn
}

// WithRotateAnimation 添加旋转动画
func (pn *PixelNebula) WithRotateAnimation(targetID string, fromAngle, toAngle float64, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewRotateAnimation(targetID, fromAngle, toAngle, duration, repeatCount)
	pn.AnimManager.AddAnimation(anim)
	return pn
}

// WithGradientAnimation 添加渐变动画
func (pn *PixelNebula) WithGradientAnimation(targetID string, colors []string, duration float64, repeatCount int, animate bool) *PixelNebula {
	anim := animation.NewGradientAnimation(targetID, colors, duration, repeatCount, animate)
	pn.AnimManager.AddAnimation(anim)
	return pn
}

// WithTransformAnimation 添加变换动画
func (pn *PixelNebula) WithTransformAnimation(targetID string, transformType string, from, to string, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewTransformAnimation(targetID, transformType, from, to, duration, repeatCount)
	pn.AnimManager.AddAnimation(anim)
	return pn
}

// WithFadeAnimation 添加淡入淡出动画
func (pn *PixelNebula) WithFadeAnimation(targetID string, from, to string, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewFadeAnimation(targetID, from, to, duration, repeatCount)
	pn.AnimManager.AddAnimation(anim)
	return pn
}

// WithPathAnimation 添加路径动画
func (pn *PixelNebula) WithPathAnimation(targetID string, path string, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewPathAnimation(targetID, path, duration, repeatCount)
	pn.AnimManager.AddAnimation(anim)
	return pn
}

// WithPathAnimationRotate 添加带旋转的路径动画
func (pn *PixelNebula) WithPathAnimationRotate(targetID string, path string, rotate string, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewPathAnimation(targetID, path, duration, repeatCount)
	anim.WithRotate(rotate)
	pn.AnimManager.AddAnimation(anim)
	return pn
}

// WithColorAnimation 添加颜色变换动画
func (pn *PixelNebula) WithColorAnimation(targetID string, property string, fromColor, toColor string, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewColorAnimation(targetID, property, fromColor, toColor, duration, repeatCount)
	pn.AnimManager.AddAnimation(anim)
	return pn
}

// WithBounceAnimation 添加弹跳动画
func (pn *PixelNebula) WithBounceAnimation(targetID string, property string, from, to string, bounceCount int, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewBounceAnimation(targetID, property, from, to, bounceCount, duration, repeatCount)
	pn.AnimManager.AddAnimation(anim)
	return pn
}

// WithWaveAnimation 添加波浪动画
func (pn *PixelNebula) WithWaveAnimation(targetID string, amplitude, frequency float64, direction string, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewWaveAnimation(targetID, amplitude, frequency, direction, duration, repeatCount)
	pn.AnimManager.AddAnimation(anim)
	return pn
}

// WithBlinkAnimation 添加闪烁动画
func (pn *PixelNebula) WithBlinkAnimation(targetID string, minOpacity, maxOpacity float64, blinkCount int, duration float64, repeatCount int) *PixelNebula {
	anim := animation.NewBlinkAnimation(targetID, minOpacity, maxOpacity, blinkCount, duration, repeatCount)
	pn.AnimManager.AddAnimation(anim)
	return pn
}

// WithParallelRender 启用并行渲染
func (pn *PixelNebula) WithParallelRender(enabled bool) *PixelNebula {
	pn.Options.ParallelRender = enabled
	return pn
}

// WithConcurrencyPool 设置并发池大小
func (pn *PixelNebula) WithConcurrencyPool(size int) *PixelNebula {
	if size <= 0 {
		size = runtime.NumCPU()
	}
	pn.Options.ConcurrencyPool = size
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
		width:      pn.Width,
		height:     pn.Height,
		themeIndex: pn.Options.ThemeIndex,
		styleIndex: pn.Options.StyleIndex,
	}
}

// SetTheme 设置主题
func (sb *SVGBuilder) SetTheme(theme int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	themeCount := sb.pn.ThemeManager.ThemeCount(sb.styleIndex)
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
	index, err := sb.pn.StyleManager.GetStyleIndex(style)
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
	themeCount := sb.pn.ThemeManager.StyleCount()
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
	sb.pn.AnimManager.AddAnimation(anim)
	return sb
}

// SetRotateAnimation 添加旋转动画
func (sb *SVGBuilder) SetRotateAnimation(targetID string, fromAngle, toAngle float64, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewRotateAnimation(targetID, fromAngle, toAngle, duration, repeatCount)
	sb.pn.AnimManager.AddAnimation(anim)
	return sb
}

// SetGradientAnimation 添加渐变动画
func (sb *SVGBuilder) SetGradientAnimation(targetID string, colors []string, duration float64, repeatCount int, animate bool) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewGradientAnimation(targetID, colors, duration, repeatCount, animate)
	sb.pn.AnimManager.AddAnimation(anim)
	return sb
}

// SetTransformAnimation 添加变换动画
func (sb *SVGBuilder) SetTransformAnimation(targetID string, transformType string, from, to string, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewTransformAnimation(targetID, transformType, from, to, duration, repeatCount)
	sb.pn.AnimManager.AddAnimation(anim)
	return sb
}

// SetFadeAnimation 添加淡入淡出动画
func (sb *SVGBuilder) SetFadeAnimation(targetID string, from, to string, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewFadeAnimation(targetID, from, to, duration, repeatCount)
	sb.pn.AnimManager.AddAnimation(anim)
	return sb
}

// SetPathAnimation 添加路径动画
func (sb *SVGBuilder) SetPathAnimation(targetID string, path string, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewPathAnimation(targetID, path, duration, repeatCount)
	sb.pn.AnimManager.AddAnimation(anim)
	return sb
}

// SetPathAnimationRotate 添加带旋转的路径动画
func (sb *SVGBuilder) SetPathAnimationRotate(targetID string, path string, rotate string, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewPathAnimation(targetID, path, duration, repeatCount)
	anim.WithRotate(rotate)
	sb.pn.AnimManager.AddAnimation(anim)
	return sb
}

// SetColorAnimation 添加颜色变换动画
func (sb *SVGBuilder) SetColorAnimation(targetID string, property string, fromColor, toColor string, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewColorAnimation(targetID, property, fromColor, toColor, duration, repeatCount)
	sb.pn.AnimManager.AddAnimation(anim)
	return sb
}

// SetBounceAnimation 添加弹跳动画
func (sb *SVGBuilder) SetBounceAnimation(targetID string, property string, from, to string, bounceCount int, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewBounceAnimation(targetID, property, from, to, bounceCount, duration, repeatCount)
	sb.pn.AnimManager.AddAnimation(anim)
	return sb
}

// SetWaveAnimation 添加波浪动画
func (sb *SVGBuilder) SetWaveAnimation(targetID string, amplitude, frequency float64, direction string, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewWaveAnimation(targetID, amplitude, frequency, direction, duration, repeatCount)
	sb.pn.AnimManager.AddAnimation(anim)
	return sb
}

// SetBlinkAnimation 添加闪烁动画
func (sb *SVGBuilder) SetBlinkAnimation(targetID string, minOpacity, maxOpacity float64, blinkCount int, duration float64, repeatCount int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	anim := animation.NewBlinkAnimation(targetID, minOpacity, maxOpacity, blinkCount, duration, repeatCount)
	sb.pn.AnimManager.AddAnimation(anim)
	return sb
}

// SetParallelRender 设置是否启用并行渲染
func (sb *SVGBuilder) SetParallelRender(enabled bool) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	sb.pn.Options.ParallelRender = enabled
	return sb
}

// SetConcurrencyPool 设置并发池大小
func (sb *SVGBuilder) SetConcurrencyPool(size int) *SVGBuilder {
	if sb.hasError != nil {
		return sb
	}
	if size <= 0 {
		size = runtime.NumCPU()
	}
	sb.pn.Options.ConcurrencyPool = size
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
	sb.pn.ImgData = []byte(svg)
	sb.pn.Width = sb.width
	sb.pn.Height = sb.height
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
		opts = pn.Options
	}
	// 验证参数
	if id == "" {
		return "", errors.ErrAvatarIDRequired
	}

	// 如果启用了缓存，先尝试从缓存获取
	if pn.Cache != nil {
		cacheKey := cache.CacheKey{
			Id:      id,
			SansEnv: sansEnv,
		}

		if opts != nil {
			cacheKey.Theme = opts.ThemeIndex
			cacheKey.Part = opts.StyleIndex
		}

		if cachedSVG, found := pn.Cache.Get(cacheKey); found {
			return cachedSVG, nil
		}
	}

	// 使用对象池获取缓冲区
	hashBuf := hashBufPool.Get().(*[]byte)
	defer hashBufPool.Put(hashBuf)

	// 计算avatarId的哈希值 - 优化版本
	pn.Hasher.Reset()
	pn.Hasher.Write([]byte(id))
	sum := pn.Hasher.Sum((*hashBuf)[:0])
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

	// 根据是否启用并行渲染选择处理方式
	if opts.ParallelRender {
		// 并行处理
		var wg sync.WaitGroup
		errChan := make(chan error, len(p))
		// 创建互斥锁来保护 final map
		var finalMux sync.Mutex

		// 对每个部分启动一个 goroutine
		for k, v := range p {
			wg.Add(1)
			go func(key string, val [2]int) {
				defer wg.Done()

				// 使用临时变量处理这个部分
				tempResult := ""

				// 获取主题颜色
				themePart, err := pn.ThemeManager.GetTheme(val[0], val[1])
				if err != nil {
					errChan <- err
					return
				}

				colors, ok := themePart[key]
				if !ok {
					errChan <- errors.ErrInvalidColor
					return
				}

				// 获取形状SVG
				shapeType := style.ShapeType(key)
				svgPart, err := pn.StyleManager.GetShape(val[0], shapeType)
				if err != nil {
					errChan <- err
					return
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

				tempResult = sb.String()

				// 归还Builder到对象池
				builderPool.Put(sb)

				// 使用互斥锁保护对 final map 的写入
				finalMux.Lock()
				final[key] = tempResult
				finalMux.Unlock()
			}(k, v)
		}

		// 等待所有部分处理完成
		wg.Wait()

		// 检查是否有错误
		select {
		case err := <-errChan:
			return "", err
		default:
			// 没有错误，继续处理
		}
	} else {
		// 串行处理
		for k, v := range p {
			if err := pn.processSVGPart(k, v, final); err != nil {
				return "", err
			}
		}
	}

	// 使用对象池获取主Builder来构建最终SVG
	builder := builderPool.Get().(*strings.Builder)
	builder.Reset()
	// 预估SVG大小，避免多次内存分配
	builder.Grow(2048) // 2KB 应该足够容纳大多数SVG

	// 添加SVG开始标签
	builder.WriteString(pn.getSvgStart())

	// 获取动画定义
	animations := pn.AnimManager.GenerateSVGAnimations()
	if animations != "" {
		builder.WriteString(animations)
	}

	// 构建和处理旋转动画 - 使用对象池
	rotateAnimations := make(map[string]bool)
	rotateAnimationSVGs := make(map[string]string)

	// 收集旋转动画
	for _, anim := range pn.AnimManager.GetAnimations() {
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

	builder.WriteString(pn.SvgEnd)
	svg = builder.String()

	// 将生成的SVG存储到实例中
	pn.ImgData = []byte(svg)

	// 归还Builder到对象池
	builderPool.Put(builder)

	// 如果启用了缓存，将结果存入缓存
	if pn.Cache != nil {
		cacheKey := cache.CacheKey{
			Id:      id,
			SansEnv: sansEnv,
		}

		if opts != nil {
			cacheKey.Theme = opts.ThemeIndex
			cacheKey.Part = opts.StyleIndex
		}

		pn.Cache.Set(cacheKey, svg)
	}

	return svg, nil
}

// GetCacheStats 获取缓存统计信息
func (pn *PixelNebula) GetCacheStats() (size, hits, misses int, hitRate float64, enabled bool, maxSize int, expiration time.Duration, evictionType string) {
	if pn.Cache == nil {
		log.Println("pixelnebula: 缓存未初始化，请先调用WithCache或WithDefaultCache")
		return 0, 0, 0, 0, false, 0, 0, ""
	}

	hits, misses, hitRate = pn.Cache.Stats()
	options := pn.Cache.GetOptions()

	return pn.Cache.Size(), hits, misses, hitRate, options.Enabled, options.Size, options.Expiration, options.EvictionType
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
	var result []CacheItemInfo

	if pn.Cache == nil {
		log.Println("pixelnebula: 缓存未初始化，请先调用WithCache或WithDefaultCache")
		return result
	}

	// 获取内部缓存项
	items := pn.Cache.GetAllItems()
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

	if pn.Cache == nil {
		log.Println("pixelnebula: 缓存未初始化，请先调用WithCache或WithDefaultCache")
		return false, 0, 0, 0, time.Time{}, nil
	}

	options := pn.Cache.GetOptions().Monitoring

	// 确保监控器已启用并初始化
	if !options.Enabled {
		log.Println("pixelnebula: 监控未启用，请设置Monitoring.Enabled=true")
		return options.Enabled, options.SampleInterval, options.AdjustInterval, options.TargetHitRate, time.Time{}, nil
	}

	if pn.Cache.GetMonitor() == nil {
		log.Println("pixelnebula: 监控器未初始化，正在初始化...")
		pn.Cache.Monitor = cache.NewMonitor(pn.Cache, options)
		pn.Cache.Monitor.Start()
	}

	monitor := pn.Cache.GetMonitor()
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
	if pn.Cache == nil {
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

	result := pn.Cache.DeleteItem(cacheKey)
	if !result {
		log.Printf("pixelnebula: 未找到缓存项: %s", key)
	}
	return result
}

// ClearCache 清空缓存
func (pn *PixelNebula) ClearCache() {
	if pn.Cache == nil {
		log.Println("pixelnebula: 缓存未初始化，请先调用WithCache或WithDefaultCache")
		return
	}

	pn.Cache.Clear()
	log.Println("pixelnebula: 缓存已清空")
}

// 将原来的 generateSVG 方法中的部分代码提取为独立函数，方便并行处理
func (pn *PixelNebula) processSVGPart(k string, v [2]int, final map[string]string) error {
	// 获取主题颜色
	themePart, err := pn.ThemeManager.GetTheme(v[0], v[1])
	if err != nil {
		return err
	}

	colors, ok := themePart[k]
	if !ok {
		return errors.ErrInvalidColor
	}

	// 获取形状SVG
	shapeType := style.ShapeType(k)
	svgPart, err := pn.StyleManager.GetShape(v[0], shapeType)
	if err != nil {
		return err
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

	return nil
}

// GenerateBatch 批量生成SVG图像
// ids：要生成的ID列表
// sansEnv：是否不包含环境
// opts：生成选项
// 返回：SVG映射表（id -> svg）和错误
func (pn *PixelNebula) GenerateBatch(ids []string, sansEnv bool, opts *PNOptions) (map[string]string, error) {
	if opts == nil {
		opts = pn.Options
	}

	// 验证参数
	if len(ids) == 0 {
		return nil, errors.ErrAvatarIDRequired
	}

	// 创建结果映射
	result := make(map[string]string, len(ids))

	// 如果没有启用并发处理，则串行生成
	if !opts.ParallelRender {
		for _, id := range ids {
			svg, err := pn.generateSVG(id, sansEnv, opts)
			if err != nil {
				return result, err
			}
			result[id] = svg
		}
		return result, nil
	}

	// 并发生成
	// 创建一个工作池限制并发数
	workerCount := opts.ConcurrencyPool
	if workerCount <= 0 {
		workerCount = runtime.NumCPU()
	}

	// 创建任务通道
	tasks := make(chan string, len(ids))
	for _, id := range ids {
		tasks <- id
	}
	close(tasks)

	// 创建结果通道
	type resultPair struct {
		id  string
		svg string
		err error
	}
	resultChan := make(chan resultPair, len(ids))

	// 启动工作池
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 创建一个新的PixelNebula实例，但只复制必要的字段，避免资源竞争
			// 特别是创建独立的哈希实例
			workerPN := &PixelNebula{
				SvgEnd:       pn.SvgEnd,
				ThemeManager: pn.ThemeManager, // 这些管理器是安全的，因为它们的方法是并发安全的或只读的
				StyleManager: pn.StyleManager,
				AnimManager:  pn.AnimManager,
				Cache:        pn.Cache,     // 缓存有自己的锁机制
				Hasher:       sha256.New(), // 创建新的哈希实例，避免并发访问冲突
				Options:      opts,
				Width:        pn.Width,
				Height:       pn.Height,
			}

			for id := range tasks {
				svg, err := workerPN.generateSVG(id, sansEnv, opts)
				resultChan <- resultPair{id, svg, err}
			}
		}()
	}

	// 等待所有工作完成
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集结果
	for r := range resultChan {
		if r.err != nil {
			return result, r.err
		}
		result[r.id] = r.svg
	}

	return result, nil
}

// SaveBatchToFiles 批量保存SVG到文件
// ids：要生成的ID列表
// sansEnv：是否不包含环境
// opts：生成选项
// filePathPattern：文件路径模式，如 "output/avatar_%s.svg"，其中 %s 将被替换为ID
// 返回：成功保存的文件数量和错误
func (pn *PixelNebula) SaveBatchToFiles(ids []string, sansEnv bool, opts *PNOptions, filePathPattern string) (int, error) {
	if opts == nil {
		opts = pn.Options
	}

	// 验证参数
	if len(ids) == 0 {
		return 0, errors.ErrAvatarIDRequired
	}

	if filePathPattern == "" {
		return 0, fmt.Errorf("文件路径模式不能为空")
	}

	// 批量生成SVG
	svgs, err := pn.GenerateBatch(ids, sansEnv, opts)
	if err != nil {
		return 0, err
	}

	// 保存文件的处理函数
	saveFile := func(id, svg, pathPattern string) error {
		filePath := fmt.Sprintf(pathPattern, id)

		// 确保目录存在
		dir := filepath.Dir(filePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}

		return os.WriteFile(filePath, []byte(svg), 0644)
	}

	// 如果不启用并行处理，串行保存
	if !opts.ParallelRender {
		count := 0
		for id, svg := range svgs {
			if err := saveFile(id, svg, filePathPattern); err != nil {
				return count, err
			}
			count++
		}
		return count, nil
	}

	// 并行保存文件
	workerCount := opts.ConcurrencyPool
	if workerCount <= 0 {
		workerCount = runtime.NumCPU()
	}

	// 创建任务通道
	type saveTask struct {
		id  string
		svg string
	}
	tasks := make(chan saveTask, len(svgs))
	for id, svg := range svgs {
		tasks <- saveTask{id, svg}
	}
	close(tasks)

	// 创建结果通道和错误通道
	type saveResult struct {
		success bool
		err     error
	}
	resultChan := make(chan saveResult, len(svgs))

	// 启动工作池
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasks {
				err := saveFile(task.id, task.svg, filePathPattern)
				resultChan <- saveResult{err == nil, err}
			}
		}()
	}

	// 等待所有工作完成
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集结果
	successCount := 0
	var firstError error

	for result := range resultChan {
		if result.success {
			successCount++
		} else if firstError == nil {
			// 保存第一个遇到的错误
			firstError = result.err
		}
	}

	return successCount, firstError
}

// BatchToBase64 批量转换SVG到Base64
// svgs: SVG字符串映射，key为ID，value为SVG内容
// width, height: SVG尺寸
// 返回: Base64编码的映射表（id -> base64）和错误
func (pn *PixelNebula) BatchToBase64(svgs map[string]string, width, height int) (map[string]string, error) {
	if len(svgs) == 0 {
		return nil, fmt.Errorf("没有需要转换的SVG")
	}

	// 创建结果映射
	result := make(map[string]string, len(svgs))

	// 如果没有启用并发处理，则串行转换
	if !pn.Options.ParallelRender {
		for id, svg := range svgs {
			conv := converter.NewSVGConverter([]byte(svg), width, height)
			base64Str, err := conv.ToBase64()
			if err != nil {
				return result, err
			}
			result[id] = base64Str
		}
		return result, nil
	}

	// 并发转换
	// 创建一个工作池限制并发数
	workerCount := pn.Options.ConcurrencyPool
	if workerCount <= 0 {
		workerCount = runtime.NumCPU()
	}

	// 创建任务通道
	type convTask struct {
		id  string
		svg string
	}
	tasks := make(chan convTask, len(svgs))
	for id, svg := range svgs {
		tasks <- convTask{id, svg}
	}
	close(tasks)

	// 创建结果通道
	type resultPair struct {
		id     string
		base64 string
		err    error
	}
	resultChan := make(chan resultPair, len(svgs))

	// 启动工作池
	var wg sync.WaitGroup
	var resultMutex sync.Mutex // 添加互斥锁保护结果映射

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasks {
				// 每个任务创建一个新的转换器实例
				conv := converter.NewSVGConverter([]byte(task.svg), width, height)
				base64Str, err := conv.ToBase64()

				// 使用通道传递结果，避免直接操作共享map
				resultChan <- resultPair{task.id, base64Str, err}
			}
		}()
	}

	// 等待所有工作完成
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集结果
	for r := range resultChan {
		if r.err != nil {
			return result, r.err
		}
		// 使用互斥锁保护map写入
		resultMutex.Lock()
		result[r.id] = r.base64
		resultMutex.Unlock()
	}

	return result, nil
}

// GenerateBatchBase64 添加批量生成Base64方法，直接从ID生成
func (pn *PixelNebula) GenerateBatchBase64(ids []string, sansEnv bool, opts *PNOptions, width, height int) (map[string]string, error) {
	if opts == nil {
		opts = pn.Options
	}

	// 先批量生成SVG
	svgs, err := pn.GenerateBatch(ids, sansEnv, opts)
	if err != nil {
		return nil, err
	}

	// 批量转换为Base64
	return pn.BatchToBase64(svgs, width, height)
}
