package animation

import (
	"fmt"
)

// RotateAnimation 旋转动画
type RotateAnimation struct {
	BaseAnimation
	FromAngle float64 // 起始角度
	ToAngle   float64 // 结束角度
	CenterX   float64 // 旋转中心X坐标
	CenterY   float64 // 旋转中心Y坐标
}

// NewRotateAnimation 创建一个旋转动画
func NewRotateAnimation(targetID string, fromAngle, toAngle float64, duration float64, repeatCount int) *RotateAnimation {
	return &RotateAnimation{
		BaseAnimation: BaseAnimation{
			Type:        Rotate,
			Duration:    duration,
			RepeatCount: repeatCount,
			Delay:       0,
			TargetID:    targetID,
			Attributes:  make(map[string]string),
		},
		FromAngle: fromAngle,
		ToAngle:   toAngle,
		CenterX:   0,
		CenterY:   0,
	}
}

// GenerateSVG 生成旋转动画的SVG代码
func (a *RotateAnimation) GenerateSVG() string {
	// 创建一个带有transform-box和transform-origin样式的g元素
	// 这个g元素将包裹目标元素及其相关元素
	svg := fmt.Sprintf("<g style=\"transform-box: fill-box; transform-origin: center;\">\n")

	// 这里只添加animateTransform元素
	svg += fmt.Sprintf("  <animateTransform attributeName=\"transform\" attributeType=\"XML\" type=\"rotate\" ")
	// 不再需要指定中心点坐标，直接使用角度值
	svg += fmt.Sprintf("from=\"%g\" to=\"%g\" ", a.FromAngle, a.ToAngle)
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

	svg += "</g>\n"

	return svg
}
