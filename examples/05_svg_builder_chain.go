package main

import (
	"fmt"
	"os"

	"github.com/landaiqing/go-pixelnebula"
	"github.com/landaiqing/go-pixelnebula/style"
)

// SVG构建器链式调用示例
// 展示如何使用链式调用API创建头像
func main() {
	// 创建一个新的PixelNebula实例
	pn := pixelnebula.NewPixelNebula().WithDefaultCache()

	// 示例1: 基本链式调用
	// 使用链式调用创建并保存头像
	svg1, err := pn.Generate("chain-example-1", false).
		SetStyle(style.AfrohairStyle).
		SetTheme(0).
		SetSize(200, 200).
		ToSVG()

	if err != nil {
		fmt.Printf("生成基本链式调用SVG失败: %v\n", err)
	} else {
		// 保存到文件
		err = os.WriteFile("basic_chain.svg", []byte(svg1), 0644)
		if err != nil {
			fmt.Printf("保存基本链式调用SVG文件失败: %v\n", err)
		} else {
			fmt.Println("成功生成基本链式调用头像: basic_chain.svg")
		}
	}

	// 示例2: 带动画的链式调用
	// 使用链式调用添加多种动画效果
	svg2, err := pn.Generate("chain-example-2", false).
		SetStyle(style.GirlStyle).
		SetTheme(1).
		SetSize(300, 300).
		// 添加旋转动画
		SetRotateAnimation("env", 0, 360, 10, -1).
		// 添加淡入淡出动画
		SetFadeAnimation("eyes", "1", "0.3", 2, -1).
		// 添加变换动画
		SetTransformAnimation("mouth", "scale", "1 1", "1.2 1.2", 1, -1).
		// 添加颜色变换动画
		SetColorAnimation("top", "fill", "#9b59b6", "#e74c3c", 3, -1).
		// 构建并获取SVG
		ToSVG()

	if err != nil {
		fmt.Printf("生成带动画的链式调用SVG失败: %v\n", err)
	} else {
		// 保存到文件
		err = os.WriteFile("animated_chain.svg", []byte(svg2), 0644)
		if err != nil {
			fmt.Printf("保存带动画的链式调用SVG文件失败: %v\n", err)
		} else {
			fmt.Println("成功生成带动画的链式调用头像: animated_chain.svg")
		}
	}

	// 示例3: 直接保存到文件的链式调用
	err = pn.Generate("chain-example-3", false).
		SetStyle(style.BlondStyle).
		SetTheme(2).
		SetSize(250, 250).
		// 添加波浪动画
		SetWaveAnimation("clo", 5, 0.2, "horizontal", 4, -1).
		// 添加闪烁动画
		SetBlinkAnimation("top", 0.3, 1.0, 4, 6, -1).
		// 构建并直接保存到文件
		Build().
		ToFile("direct_file_chain.svg")

	if err != nil {
		fmt.Printf("直接保存到文件的链式调用失败: %v\n", err)
	} else {
		fmt.Println("成功生成并直接保存头像到文件: direct_file_chain.svg")
	}

	// 示例4: 转换为Base64的链式调用
	base64, err := pn.Generate("chain-example-4", false).
		SetStyle(style.BlondStyle).
		SetTheme(0).
		SetSize(200, 200).
		// 添加旋转动画
		SetRotateAnimation("head", 0, 360, 15, -1).
		// 构建并转换为Base64
		ToBase64()

	if err != nil {
		fmt.Printf("转换为Base64的链式调用失败: %v\n", err)
	} else {
		// 保存Base64编码到文件
		err = os.WriteFile("base64_avatar.txt", []byte(base64), 0644)
		if err != nil {
			fmt.Printf("保存Base64编码到文件失败: %v\n", err)
		} else {
			fmt.Println("成功生成Base64编码头像并保存到文件: base64_avatar.txt")
		}
	}

	fmt.Println("SVG构建器链式调用示例完成！")
}
