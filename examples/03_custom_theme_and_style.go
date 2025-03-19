package main

import (
	"fmt"
	"os"

	"github.com/landaiqing/go-pixelnebula"
	"github.com/landaiqing/go-pixelnebula/style"
	"github.com/landaiqing/go-pixelnebula/theme"
)

// 自定义主题和风格示例
// 展示如何创建自定义主题和风格
func main() {
	// 创建一个新的PixelNebula实例
	pn := pixelnebula.NewPixelNebula()

	// 1. 自定义主题示例
	// 创建自定义主题 - 每个主题包含各部分的颜色设置
	customThemes := []theme.Theme{
		{
			theme.ThemePart{
				// 环境部分颜色
				"env": []string{"#FF5733", "#C70039"},
				// 头部颜色
				"head": []string{"#FFC300", "#FF5733"},
				// 衣服颜色
				"clo": []string{"#2E86C1", "#1A5276"},
				// 眼睛颜色
				"eyes": []string{"#000000", "#FFFFFF"},
				// 嘴巴颜色
				"mouth": []string{"#E74C3C"},
				// 头顶装饰颜色
				"top": []string{"#884EA0", "#7D3C98"},
			},
			theme.ThemePart{
				// 另一个主题配色
				"env":   []string{"#3498DB", "#2874A6"},
				"head":  []string{"#F5CBA7", "#F0B27A"},
				"clo":   []string{"#27AE60", "#196F3D"},
				"eyes":  []string{"#2C3E50", "#FDFEFE"},
				"mouth": []string{"#CB4335"},
				"top":   []string{"#D35400", "#BA4A00"},
			},
		},
	}

	// 应用自定义主题
	pn.WithCustomizeTheme(customThemes)

	// 生成使用自定义主题的头像
	pn.WithSize(250, 250)
	pn.WithTheme(0)

	// 生成第一个自定义主题的头像
	svg1, err := pn.Generate("custom-theme-1", false).SetTheme(0).ToSVG()
	if err != nil {
		fmt.Printf("生成自定义主题1的SVG失败: %v\n", err)
	} else {
		// 保存到文件
		err = os.WriteFile("custom_theme_1.svg", []byte(svg1), 0644)
		if err != nil {
			fmt.Printf("保存自定义主题1文件失败: %v\n", err)
		} else {
			fmt.Println("成功生成自定义主题1头像: custom_theme_1.svg")
		}
	}

	// 生成第二个自定义主题的头像
	svg2, err := pn.Generate("custom-theme-2", false).SetTheme(1).ToSVG()
	if err != nil {
		fmt.Printf("生成自定义主题2的SVG失败: %v\n", err)
	} else {
		// 保存到文件
		err = os.WriteFile("custom_theme_2.svg", []byte(svg2), 0644)
		if err != nil {
			fmt.Printf("保存自定义主题2文件失败: %v\n", err)
		} else {
			fmt.Println("成功生成自定义主题2头像: custom_theme_2.svg")
		}
	}

	// 2. 自定义风格示例
	// 创建一个新的PixelNebula实例，用于自定义风格
	pn2 := pixelnebula.NewPixelNebula()

	// 创建自定义风格 - 每种风格包含不同形状部件的SVG路径
	// 注意：这里仅作示例，实际使用中需要提供完整的SVG路径数据
	customStyles := []style.StyleSet{
		{
			// 第一种自定义风格
			style.TypeEnv:   `<circle cx="50%" cy="50%" r="48%" fill="#FILL0;"></circle>`,
			style.TypeHead:  `<circle cx="50%" cy="50%" r="35%" fill="#FILL0;"></circle>`,
			style.TypeClo:   `<rect x="25%" y="65%" width="50%" height="30%" fill="#FILL0;"></rect>`,
			style.TypeEyes:  `<circle cx="40%" cy="45%" r="5%" fill="#FILL0;"></circle><circle cx="60%" cy="45%" r="5%" fill="#FILL1;"></circle>`,
			style.TypeMouth: `<path d="M 40% 60% Q 50% 70% 60% 60%" stroke="#FILL0;" stroke-width="2" fill="none"></path>`,
			style.TypeTop:   `<path d="M 30% 30% L 50% 10% L 70% 30%" stroke="#FILL0;" stroke-width="4" fill="#FILL1;"></path>`,
		},
	}

	// 应用自定义风格
	pn2.WithCustomizeStyle(customStyles)
	pn2.WithSize(250, 250)

	// 使用自定义风格生成头像
	svg3, err := pn2.Generate("custom-style", false).SetStyleByIndex(0).ToSVG()
	if err != nil {
		fmt.Printf("生成自定义风格的SVG失败: %v\n", err)
	} else {
		// 保存到文件
		err = os.WriteFile("custom_style.svg", []byte(svg3), 0644)
		if err != nil {
			fmt.Printf("保存自定义风格文件失败: %v\n", err)
		} else {
			fmt.Println("成功生成自定义风格头像: custom_style.svg")
		}
	}

	fmt.Println("自定义主题和风格示例完成！")
}
