package main

import (
	"fmt"
	"os"

	"github.com/landaiqing/go-pixelnebula"
	"github.com/landaiqing/go-pixelnebula/style"
)

// 所有动画效果示例
// 展示PixelNebula支持的所有动画类型
func main() {
	// 创建一个新的PixelNebula实例
	pn := pixelnebula.NewPixelNebula()

	// 设置风格和尺寸
	pn.WithStyle(style.AfrohairStyle)
	pn.WithTheme(0)
	pn.WithSize(300, 300)

	// 1. 旋转动画 - 让环境和头部旋转
	pn.WithRotateAnimation("env", 0, 360, 10, -1)  // 无限循环旋转环境
	pn.WithRotateAnimation("head", 0, 360, 15, -1) // 无限循环旋转头部

	// 2. 渐变动画 - 给环境添加渐变色
	pn.WithGradientAnimation("env", []string{"#3498db", "#2ecc71", "#f1c40f", "#e74c3c", "#9b59b6"}, 8, -1, true)

	// 3. 淡入淡出动画 - 让眼睛闪烁
	pn.WithFadeAnimation("eyes", "1", "0.3", 2, -1)

	// 4. 变换动画 - 让嘴巴缩放
	pn.WithTransformAnimation("mouth", "scale", "1 1", "1.2 1.2", 1, -1)

	// 5. 颜色变换动画 - 让头顶装饰变色
	pn.WithColorAnimation("top", "fill", "#9b59b6", "#e74c3c", 3, -1)

	// 6. 弹跳动画 - 让整个头像上下弹跳
	pn.WithBounceAnimation("head", "translateY", "0", "-10", 3, 5, -1)

	// 7. 波浪动画 - 让衣服产生波浪效果
	pn.WithWaveAnimation("clo", 5, 0.2, "horizontal", 4, -1)

	// 8. 闪烁动画 - 让头顶装饰闪烁
	pn.WithBlinkAnimation("top", 0.3, 1.0, 4, 6, -1)

	// 9. 路径动画 - 让眼睛沿着路径移动
	pn.WithPathAnimation("eyes", "M 0,0 C 10,-10 -10,-10 0,0", 3, -1)

	// 10. 带旋转的路径动画 - 让眼睛在移动的同时旋转
	pn.WithPathAnimationRotate("mouth", "M 0,0 C 5,5 -5,5 0,0", "auto", 4, -1)

	// 生成SVG
	svg, err := pn.Generate("all-animations-example", false).ToSVG()
	if err != nil {
		fmt.Printf("生成SVG失败: %v\n", err)
		os.Exit(1)
	}

	// 保存到文件
	err = os.WriteFile("all_animations.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Printf("保存文件失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("成功生成包含所有动画效果的头像: all_animations.svg")
}
