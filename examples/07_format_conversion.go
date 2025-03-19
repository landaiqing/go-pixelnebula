package main

import (
	"fmt"
	"os"

	"github.com/landaiqing/go-pixelnebula"
	"github.com/landaiqing/go-pixelnebula/style"
)

// 格式转换示例
// 展示如何将SVG转换为其他格式
func main() {
	// 创建一个新的PixelNebula实例
	pn := pixelnebula.NewPixelNebula()

	// 设置基本属性
	pn.WithStyle(style.AfrohairStyle)
	pn.WithTheme(0)
	pn.WithSize(500, 500) // 使用较大尺寸以便转换后的图像清晰

	// 添加一些动画效果
	pn.WithRotateAnimation("env", 0, 360, 10, -1)
	pn.WithGradientAnimation("env", []string{"#3498db", "#2ecc71", "#f1c40f"}, 5, -1, true)

	// 1. 生成并保存SVG文件
	svgData, err := pn.Generate("format-conversion", false).ToSVG()
	if err != nil {
		fmt.Printf("生成SVG失败: %v\n", err)
		os.Exit(1)
	}

	// 保存SVG文件
	err = os.WriteFile("format_conversion.svg", []byte(svgData), 0644)
	if err != nil {
		fmt.Printf("保存SVG文件失败: %v\n", err)
	} else {
		fmt.Println("成功生成SVG文件: format_conversion.svg")
	}

	// 2. 转换为Base64格式
	base64Data, err := pn.Generate("format-conversion", false).ToBase64()
	if err != nil {
		fmt.Printf("转换为Base64失败: %v\n", err)
	} else {
		// 保存Base64数据到文件
		err = os.WriteFile("format_conversion.base64.txt", []byte(base64Data), 0644)
		if err != nil {
			fmt.Printf("保存Base64文件失败: %v\n", err)
		} else {
			fmt.Println("成功生成Base64编码文件: format_conversion.base64.txt")
		}
	}
	fmt.Println("格式转换示例完成！")
}
