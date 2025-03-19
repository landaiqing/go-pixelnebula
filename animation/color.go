package animation

import (
	"fmt"
	"strings"
)

// ColorAnimation 颜色变换动画
type ColorAnimation struct {
	BaseAnimation
	FromColor string // 起始颜色
	ToColor   string // 结束颜色
	Property  string // 要变换的属性（fill 或 stroke）
}

// NewColorAnimation 创建一个颜色变换动画
func NewColorAnimation(targetID string, property string, fromColor, toColor string, duration float64, repeatCount int) *ColorAnimation {
	return &ColorAnimation{
		BaseAnimation: BaseAnimation{
			Type:        Color, // 颜色变换动画类型
			Duration:    duration,
			RepeatCount: repeatCount,
			Delay:       0,
			TargetID:    targetID,
			Attributes:  make(map[string]string),
		},
		FromColor: fromColor,
		ToColor:   toColor,
		Property:  property,
	}
}

// GenerateSVG 生成颜色变换动画的SVG代码
func (a *ColorAnimation) GenerateSVG() string {
	var sb strings.Builder

	// 创建一个animate元素
	sb.WriteString(fmt.Sprintf("<animate href=\"#%s\" attributeName=\"%s\" ", a.TargetID, a.Property))
	sb.WriteString(fmt.Sprintf("from=\"%s\" to=\"%s\" ", a.FromColor, a.ToColor))
	sb.WriteString(fmt.Sprintf("dur=\"%gs\" ", a.Duration))

	if a.RepeatCount < 0 {
		sb.WriteString("repeatCount=\"indefinite\" ")
	} else if a.RepeatCount > 0 {
		sb.WriteString(fmt.Sprintf("repeatCount=\"%d\" ", a.RepeatCount))
	}

	if a.Delay > 0 {
		sb.WriteString(fmt.Sprintf("begin=\"%gs\" ", a.Delay))
	}

	sb.WriteString("/>\n")

	return sb.String()
}
