package animation

import (
	"fmt"
	"math"
	"strings"
)

// WaveAnimation 波浪动画
type WaveAnimation struct {
	BaseAnimation
	Amplitude float64 // 波浪振幅
	Frequency float64 // 波浪频率
	Direction string  // 波浪方向 ("horizontal" 或 "vertical")
}

// NewWaveAnimation 创建一个波浪动画
func NewWaveAnimation(targetID string, amplitude, frequency float64, direction string, duration float64, repeatCount int) *WaveAnimation {
	return &WaveAnimation{
		BaseAnimation: BaseAnimation{
			Type:        Wave, // 波浪动画类型
			Duration:    duration,
			RepeatCount: repeatCount,
			Delay:       0,
			TargetID:    targetID,
			Attributes:  make(map[string]string),
		},
		Amplitude: amplitude,
		Frequency: frequency,
		Direction: direction,
	}
}

// GenerateSVG 生成波浪动画的SVG代码
func (a *WaveAnimation) GenerateSVG() string {
	var sb strings.Builder

	// 生成波浪路径
	path := a.generateWavePath()

	// 创建一个animateMotion元素
	sb.WriteString(fmt.Sprintf("<animateMotion href=\"#%s\" ", a.TargetID))
	sb.WriteString(fmt.Sprintf("path=\"%s\" ", path))
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

// generateWavePath 生成波浪路径
func (a *WaveAnimation) generateWavePath() string {
	var path strings.Builder
	path.WriteString("M0,0 ")

	// 生成正弦波路径
	points := 20 // 路径点数量
	for i := 0; i <= points; i++ {
		x := float64(i) / float64(points) * 100 // 0-100 范围
		// 计算正弦波 y 值
		y := a.Amplitude * math.Sin(a.Frequency*x*math.Pi/180)

		if a.Direction == "horizontal" {
			path.WriteString(fmt.Sprintf("L%g,%g ", x, y))
		} else { // vertical
			path.WriteString(fmt.Sprintf("L%g,%g ", y, x))
		}
	}

	return path.String()
}
