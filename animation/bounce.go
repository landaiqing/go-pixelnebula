package animation

import (
	"fmt"
	"strings"
)

// BounceAnimation 弹跳动画
type BounceAnimation struct {
	BaseAnimation
	Property    string // 要变换的属性（如 y, transform 等）
	From        string // 起始值
	To          string // 结束值
	BounceCount int    // 弹跳次数
}

// NewBounceAnimation 创建一个弹跳动画
func NewBounceAnimation(targetID string, property string, from, to string, bounceCount int, duration float64, repeatCount int) *BounceAnimation {
	return &BounceAnimation{
		BaseAnimation: BaseAnimation{
			Type:        Bounce, // 弹跳动画类型
			Duration:    duration,
			RepeatCount: repeatCount,
			Delay:       0,
			TargetID:    targetID,
			Attributes:  make(map[string]string),
		},
		Property:    property,
		From:        from,
		To:          to,
		BounceCount: bounceCount,
	}
}

// GenerateSVG 生成弹跳动画的SVG代码
func (a *BounceAnimation) GenerateSVG() string {
	var sb strings.Builder

	// 创建animateTransform元素，使用transform属性
	if a.Property == "transform" {
		sb.WriteString(fmt.Sprintf("<animateTransform href=\"#%s\" attributeName=\"transform\" type=\"translate\" ", a.TargetID))
		sb.WriteString(fmt.Sprintf("from=\"%s\" to=\"%s\" ", a.From, a.To))
	} else {
		sb.WriteString(fmt.Sprintf("<animate href=\"#%s\" attributeName=\"%s\" ", a.TargetID, a.Property))
		sb.WriteString(fmt.Sprintf("from=\"%s\" to=\"%s\" ", a.From, a.To))
	}

	sb.WriteString(fmt.Sprintf("dur=\"%gs\" ", a.Duration))

	// 生成弹跳效果的关键帧
	var keyTimes, values []string
	step := 1.0 / float64(a.BounceCount*2)

	for i := 0; i <= a.BounceCount*2; i++ {
		// 调整关键帧时间，使动画更加平滑
		keyTime := float64(i) * step
		// 为每个弹跳周期添加额外的中间帧
		if i > 0 && i < a.BounceCount*2 {
			keyTime = keyTime + (step * 0.1) // 稍微延长每次弹跳的时间
		}
		keyTimes = append(keyTimes, fmt.Sprintf("%.3f", keyTime))

		if i%2 == 0 {
			values = append(values, a.From)
		} else {
			values = append(values, a.To)
		}
	}

	// 添加关键帧属性
	sb.WriteString(fmt.Sprintf("values=\"%s\" ", strings.Join(values, ";")))
	sb.WriteString(fmt.Sprintf("keyTimes=\"%s\" ", strings.Join(keyTimes, ";")))

	// 添加缓动函数
	sb.WriteString("calcMode=\"spline\" ")
	sb.WriteString("keySplines=\"")
	for i := 0; i < len(values)-1; i++ {
		if i > 0 {
			sb.WriteString(";")
		}
		if i%2 == 0 {
			// 快速上升
			sb.WriteString("0.2 0 0.8 1")
		} else {
			// 缓慢下落
			sb.WriteString("0.2 0.8 0.8 1")
		}
	}
	sb.WriteString("\" ")

	if a.RepeatCount < 0 {
		sb.WriteString("repeatCount=\"indefinite\" ")
	} else if a.RepeatCount > 0 {
		sb.WriteString(fmt.Sprintf("repeatCount=\"%d\" ", a.RepeatCount))
	}

	if a.Delay > 0 {
		sb.WriteString(fmt.Sprintf("begin=\"%gs\" ", a.Delay))
	}

	// 添加fill属性
	sb.WriteString("fill=\"freeze\" />")

	return sb.String()
}
