package animation

import (
	"fmt"
	"strings"
)

// BlinkAnimation 闪烁动画
type BlinkAnimation struct {
	BaseAnimation
	BlinkCount int     // 闪烁次数
	MinOpacity float64 // 最小透明度
	MaxOpacity float64 // 最大透明度
}

// NewBlinkAnimation 创建一个闪烁动画
func NewBlinkAnimation(targetID string, minOpacity, maxOpacity float64, blinkCount int, duration float64, repeatCount int) *BlinkAnimation {
	return &BlinkAnimation{
		BaseAnimation: BaseAnimation{
			Type:        Blink, // 闪烁动画类型
			Duration:    duration,
			RepeatCount: repeatCount,
			Delay:       0,
			TargetID:    targetID,
			Attributes:  make(map[string]string),
		},
		BlinkCount: blinkCount,
		MinOpacity: minOpacity,
		MaxOpacity: maxOpacity,
	}
}

// GenerateSVG 生成闪烁动画的SVG代码
func (a *BlinkAnimation) GenerateSVG() string {
	var sb strings.Builder

	// 创建一个animate元素
	sb.WriteString(fmt.Sprintf("<animate href=\"#%s\" attributeName=\"opacity\" ", a.TargetID))

	// 根据闪烁次数生成关键帧
	var keyTimes, values []string

	// 计算关键帧
	for i := 0; i <= a.BlinkCount*2; i++ {
		// 计算关键帧时间点
		keyTime := float64(i) / float64(a.BlinkCount*2)
		keyTimes = append(keyTimes, fmt.Sprintf("%.2f", keyTime))

		// 计算关键帧值，交替使用最大和最小透明度
		if i%2 == 0 {
			values = append(values, fmt.Sprintf("%.1f", a.MaxOpacity))
		} else {
			values = append(values, fmt.Sprintf("%.1f", a.MinOpacity))
		}
	}

	// 添加关键帧属性
	sb.WriteString(fmt.Sprintf("keyTimes=\"%s\" ", strings.Join(keyTimes, ";")))
	sb.WriteString(fmt.Sprintf("values=\"%s\" ", strings.Join(values, ";")))
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
