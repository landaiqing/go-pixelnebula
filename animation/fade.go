package animation

import (
	"fmt"
)

// FadeAnimation 淡入淡出动画
type FadeAnimation struct {
	BaseAnimation
	From string // 起始透明度
	To   string // 结束透明度
}

// NewFadeAnimation 创建一个淡入淡出动画
func NewFadeAnimation(targetID string, from, to string, duration float64, repeatCount int) *FadeAnimation {
	return &FadeAnimation{
		BaseAnimation: BaseAnimation{
			Type:        Fade,
			Duration:    duration,
			RepeatCount: repeatCount,
			Delay:       0,
			TargetID:    targetID,
			Attributes:  make(map[string]string),
		},
		From: from,
		To:   to,
	}
}

// GenerateSVG 生成淡入淡出动画的SVG代码
func (a *FadeAnimation) GenerateSVG() string {
	// 创建一个animate元素，并将其添加到目标元素中
	svg := fmt.Sprintf("<animate href=\"#%s\" attributeName=\"opacity\" ", a.TargetID)
	svg += fmt.Sprintf("from=\"%s\" to=\"%s\" ", a.From, a.To)
	svg += fmt.Sprintf("dur=\"%gs\" ", a.Duration)

	if a.RepeatCount < 0 {
		svg += "repeatCount=\"indefinite\" "
	} else if a.RepeatCount > 0 {
		svg += fmt.Sprintf("repeatCount=\"%d\" ", a.RepeatCount)
	}

	if a.Delay > 0 {
		svg += fmt.Sprintf("begin=\"%gs\" ", a.Delay)
	}

	svg += "/>\n"

	return svg
}
