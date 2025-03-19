package animation

import (
	"fmt"
	"strings"
)

// PathAnimation 路径动画
type PathAnimation struct {
	BaseAnimation
	Path   string // SVG路径数据
	Rotate string // 是否旋转元素以跟随路径方向 ("auto", "auto-reverse", 或 "0")
}

// NewPathAnimation 创建一个路径动画
func NewPathAnimation(targetID string, path string, duration float64, repeatCount int) *PathAnimation {
	return &PathAnimation{
		BaseAnimation: BaseAnimation{
			Type:        Path, // 路径动画类型
			Duration:    duration,
			RepeatCount: repeatCount,
			Delay:       0,
			TargetID:    targetID,
			Attributes:  make(map[string]string),
		},
		Path:   path,
		Rotate: "0", // 默认不旋转
	}
}

// WithRotate 设置是否旋转元素以跟随路径方向
func (a *PathAnimation) WithRotate(rotate string) *PathAnimation {
	a.Rotate = rotate
	return a
}

// GenerateSVG 生成路径动画的SVG代码
func (a *PathAnimation) GenerateSVG() string {
	var sb strings.Builder

	// 创建一个animateMotion元素
	sb.WriteString(fmt.Sprintf("<animateMotion href=\"#%s\" ", a.TargetID))
	sb.WriteString(fmt.Sprintf("path=\"%s\" ", a.Path))
	sb.WriteString(fmt.Sprintf("dur=\"%gs\" ", a.Duration))

	if a.RepeatCount < 0 {
		sb.WriteString("repeatCount=\"indefinite\" ")
	} else if a.RepeatCount > 0 {
		sb.WriteString(fmt.Sprintf("repeatCount=\"%d\" ", a.RepeatCount))
	}

	if a.Delay > 0 {
		sb.WriteString(fmt.Sprintf("begin=\"%gs\" ", a.Delay))
	}

	// 设置旋转属性
	sb.WriteString(fmt.Sprintf("rotate=\"%s\" ", a.Rotate))

	sb.WriteString("/>\n")

	return sb.String()
}
