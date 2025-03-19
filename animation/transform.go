package animation

import (
	"fmt"
)

// TransformAnimation 变换动画
type TransformAnimation struct {
	BaseAnimation
	TransformType string // 变换类型（scale, translate等）
	From          string // 起始变换值
	To            string // 结束变换值
}

// NewTransformAnimation 创建一个变换动画
func NewTransformAnimation(targetID string, transformType string, from, to string, duration float64, repeatCount int) *TransformAnimation {
	return &TransformAnimation{
		BaseAnimation: BaseAnimation{
			Type:        Transform,
			Duration:    duration,
			RepeatCount: repeatCount,
			Delay:       0,
			TargetID:    targetID,
			Attributes:  make(map[string]string),
		},
		TransformType: transformType,
		From:          from,
		To:            to,
	}
}

// GenerateSVG 生成变换动画的SVG代码
func (a *TransformAnimation) GenerateSVG() string {
	// 创建一个animateTransform元素，并将其添加到目标元素中
	svg := fmt.Sprintf("<animateTransform href=\"#%s\" attributeName=\"transform\" attributeType=\"XML\" type=\"%s\" ", a.TargetID, a.TransformType)
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

	svg += "additive=\"sum\" />\n"

	return svg
}
