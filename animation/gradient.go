package animation

import (
	"fmt"
	"strings"
)

// GradientAnimation 渐变动画
type GradientAnimation struct {
	BaseAnimation
	Colors  []string // 渐变颜色列表
	Animate bool     // 是否添加动画效果
}

// NewGradientAnimation 创建一个渐变动画
func NewGradientAnimation(targetID string, colors []string, duration float64, repeatCount int, animate bool) *GradientAnimation {
	return &GradientAnimation{
		BaseAnimation: BaseAnimation{
			Type:        Gradient,
			Duration:    duration,
			RepeatCount: repeatCount,
			Delay:       0,
			TargetID:    targetID,
			Attributes:  make(map[string]string),
		},
		Colors:  colors,
		Animate: animate,
	}
}

// GenerateSVG 生成渐变动画的SVG代码
func (a *GradientAnimation) GenerateSVG() string {
	var sb strings.Builder

	// 创建渐变定义
	gradientID := fmt.Sprintf("%s-gradient", a.TargetID)
	sb.WriteString(fmt.Sprintf("<linearGradient id=\"%s\" x1=\"0%%\" y1=\"0%%\" x2=\"100%%\" y2=\"0%%\">\n", gradientID))

	// 添加渐变颜色
	for i, color := range a.Colors {
		offset := float64(i) / float64(len(a.Colors)-1) * 100
		sb.WriteString(fmt.Sprintf("  <stop offset=\"%g%%\" stop-color=\"%s\" />\n", offset, color))
	}
	sb.WriteString("</linearGradient>\n")

	// 为目标元素添加样式引用
	sb.WriteString(fmt.Sprintf("<style type=\"text/css\">\n  #%s { fill: url(#%s) !important; }\n</style>\n", a.TargetID, gradientID))

	// 添加动画
	if a.Animate {
		// x1 动画
		sb.WriteString(fmt.Sprintf("<animate href=\"#%s\" attributeName=\"x1\" from=\"0%%\" to=\"100%%\" ", gradientID))
		sb.WriteString(fmt.Sprintf("dur=\"%gs\" ", a.Duration))
		if a.RepeatCount < 0 {
			sb.WriteString("repeatCount=\"indefinite\" ")
		} else if a.RepeatCount > 0 {
			sb.WriteString(fmt.Sprintf("repeatCount=\"%d\" ", a.RepeatCount))
		}
		sb.WriteString("/>\n")

		// x2 动画
		sb.WriteString(fmt.Sprintf("<animate href=\"#%s\" attributeName=\"x2\" from=\"100%%\" to=\"200%%\" ", gradientID))
		sb.WriteString(fmt.Sprintf("dur=\"%gs\" ", a.Duration))
		if a.RepeatCount < 0 {
			sb.WriteString("repeatCount=\"indefinite\" ")
		} else if a.RepeatCount > 0 {
			sb.WriteString(fmt.Sprintf("repeatCount=\"%d\" ", a.RepeatCount))
		}
		sb.WriteString("/>\n")
	}

	return sb.String()
}
