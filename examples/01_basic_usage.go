package main

import (
	"fmt"
	"os"

	"github.com/landaiqing/go-pixelnebula"
	"github.com/landaiqing/go-pixelnebula/style"
)

// 基本用法示例
// 展示如何创建简单的PixelNebula头像
func main() {
	// 创建一个新的PixelNebula实例
	pn := pixelnebula.NewPixelNebula()

	// 设置风格 - 这里使用默认的AfrohairStyle风格
	pn.WithStyle(style.AfrohairStyle)

	// 设置主题索引 - 每种风格有多个主题可选
	pn.WithTheme(0)

	// 设置头像尺寸 (宽度, 高度)
	pn.WithSize(300, 300)

	// 生成SVG - 需要提供唯一ID和是否生成无环境模式的参数
	// 第一个参数：唯一标识符，用于生成不同的头像
	// 第二个参数：是否为无环境模式，true表示不生成背景环境
	svg, err := pn.Generate("my-unique-id-123", false).ToSVG()
	if err != nil {
		fmt.Printf("生成SVG失败: %v\n", err)
		os.Exit(1)
	}

	// 保存到文件
	err = os.WriteFile("basic_avatar.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Printf("保存文件失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("成功生成基本头像: basic_avatar.svg")

	// 再生成一个无环境模式的头像
	svgNoEnv, err := pn.Generate("my-unique-id-123", true).ToSVG()
	if err != nil {
		fmt.Printf("生成无环境SVG失败: %v\n", err)
		os.Exit(1)
	}

	// 保存到文件
	err = os.WriteFile("basic_avatar_no_env.svg", []byte(svgNoEnv), 0644)
	if err != nil {
		fmt.Printf("保存文件失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("成功生成无环境头像: basic_avatar_no_env.svg")
}
