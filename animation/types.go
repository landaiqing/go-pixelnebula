package animation

import (
	"strings"
)

// AnimationType 表示动画类型
type AnimationType string

// 预定义动画类型常量
const (
	Rotate    AnimationType = "rotate"    // 旋转动画
	Gradient  AnimationType = "gradient"  // 渐变动画
	Transform AnimationType = "transform" // 变换动画
	Fade      AnimationType = "fade"      // 淡入淡出动画
	Path      AnimationType = "path"      // 路径动画
	Color     AnimationType = "color"     // 颜色变换动画
	Bounce    AnimationType = "bounce"    // 弹跳动画
	Wave      AnimationType = "wave"      // 波浪动画
	Blink     AnimationType = "blink"     // 闪烁动画
)

// Animation 表示一个SVG动画接口
type Animation interface {
	// GenerateSVG 生成动画的SVG代码
	GenerateSVG() string
	// GetTargetID 获取目标元素ID
	GetTargetID() string
}

// BaseAnimation 基础动画结构，包含所有动画共有的属性
type BaseAnimation struct {
	Type        AnimationType     // 动画类型
	Duration    float64           // 动画持续时间（秒）
	RepeatCount int               // 重复次数，-1表示无限重复
	Delay       float64           // 延迟时间（秒）
	TargetID    string            // 目标元素ID
	Attributes  map[string]string // 动画属性
}

// GetTargetID 获取目标元素ID
func (a *BaseAnimation) GetTargetID() string {
	return a.TargetID
}

// Manager 动画管理器，负责管理所有动画
type Manager struct {
	animations []Animation
}

// GetAnimations 获取所有动画
func (m *Manager) GetAnimations() []Animation {
	return m.animations
}

// NewAnimationManager 创建一个新的动画管理器
func NewAnimationManager() *Manager {
	return &Manager{
		animations: make([]Animation, 0),
	}
}

// AddAnimation 添加一个动画
func (m *Manager) AddAnimation(animation Animation) {
	m.animations = append(m.animations, animation)
}

// GenerateSVGAnimations 生成SVG动画代码
func (m *Manager) GenerateSVGAnimations() string {
	if len(m.animations) == 0 {
		return ""
	}

	var sb strings.Builder

	// 添加SVG命名空间声明
	sb.WriteString("<defs>\n")

	// 用于存储需要放在defs中的定义
	var defsContent strings.Builder
	// 用于存储需要直接添加到SVG中的动画元素
	var animationsContent strings.Builder
	// 用于存储旋转动画的映射，键为目标元素ID
	rotateAnimations := make(map[string]string)

	// 处理所有动画
	for _, anim := range m.animations {
		svgCode := anim.GenerateSVG()
		if svgCode == "" {
			continue
		}

		// 根据动画类型决定放置位置
		switch a := anim.(type) {
		case *GradientAnimation:
			// 渐变定义需要放在defs中
			defsContent.WriteString(svgCode)
		case *RotateAnimation:
			// 旋转动画需要包裹目标元素，先存储起来
			// 提取animateTransform标签
			if start := strings.Index(svgCode, "<animateTransform"); start != -1 {
				if end := strings.Index(svgCode[start:], "/>"); end != -1 {
					rotateAnimations[a.GetTargetID()] = svgCode[start : start+end+2]
				}
			}
		default:
			// 其他动画元素直接添加到SVG中
			animationsContent.WriteString(svgCode)
		}
	}

	// 只有当存在需要放在defs中的内容时才添加defs标签
	if defsContent.Len() > 0 {
		sb.WriteString(defsContent.String())
		sb.WriteString("</defs>\n")
	} else {
		// 如果没有需要放在defs中的内容，则不添加defs标签
		sb.Reset()
	}

	// 添加直接放置的动画元素
	sb.WriteString(animationsContent.String())

	return sb.String()
}
