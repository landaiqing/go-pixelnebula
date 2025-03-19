package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/landaiqing/go-pixelnebula"
	"github.com/landaiqing/go-pixelnebula/style"
)

// 风格和主题示例
// 展示如何使用不同的风格和主题生成多个头像
func main() {
	// 创建一个新的PixelNebula实例
	pn := pixelnebula.NewPixelNebula()

	// 定义要展示的风格数组
	styles := []style.StyleType{
		style.AteamStyle,
		style.GirlStyle,
		style.CountryStyle,
		style.GeeknotStyle,
		style.PunkStyle,
		// 可以添加更多内置风格
	}

	// 为每种风格生成不同主题的头像
	for styleIndex, styleType := range styles {
		// 设置当前风格
		pn.WithStyle(styleType)

		// 获取风格名称用于文件命名
		var styleName string
		switch styleType {
		case style.AteamStyle:
			styleName = "ateam"
		case style.GirlStyle:
			styleName = "girl"
		case style.CountryStyle:
			styleName = "country"
		case style.GeeknotStyle:
			styleName = "geeknot"
		case style.PunkStyle:
			styleName = "punk"
		default:
			styleName = "unknown"
		}

		// 对每种风格，生成3个不同主题的头像
		for themeIndex := 0; themeIndex < 3; themeIndex++ {
			// 设置主题
			pn.WithTheme(themeIndex)

			// 设置尺寸
			pn.WithSize(200, 200)

			// 生成唯一ID - 这里使用风格和主题索引组合
			uniqueID := "style-" + strconv.Itoa(styleIndex) + "-theme-" + strconv.Itoa(themeIndex)

			// 生成SVG
			svg, err := pn.Generate(uniqueID, false).ToSVG()
			if err != nil {
				fmt.Printf("生成风格%s主题%d的SVG失败: %v\n", styleName, themeIndex, err)
				continue
			}

			// 文件名
			filename := fmt.Sprintf("%s_theme_%d.svg", styleName, themeIndex)

			// 保存到文件
			err = os.WriteFile(filename, []byte(svg), 0644)
			if err != nil {
				fmt.Printf("保存文件%s失败: %v\n", filename, err)
				continue
			}

			fmt.Printf("成功生成头像: %s\n", filename)
		}
	}

	fmt.Println("所有风格和主题头像生成完成！")
}
