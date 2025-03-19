package benchmark

import (
	"testing"

	"github.com/landaiqing/go-pixelnebula"
	"github.com/landaiqing/go-pixelnebula/style"
	"github.com/landaiqing/go-pixelnebula/theme"
)

// BenchmarkDifferentStyles 测试不同风格的生成性能
func BenchmarkDifferentStyles(b *testing.B) {
	styles := []struct {
		name  string
		style style.StyleType
	}{
		{"GirlStyle", style.GirlStyle},
		{"AteamStyle", style.AteamStyle},
		{"BlondStyle", style.BlondStyle},
		{"FirehairStyle", style.FirehairStyle},
	}

	for _, s := range styles {
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				pn := pixelnebula.NewPixelNebula()
				pn.WithStyle(s.style)
				pn.WithSize(231, 231)

				_, err := pn.Generate("benchmark-style-"+s.name, false).ToSVG()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkDifferentThemes 测试不同主题的生成性能
func BenchmarkDifferentThemes(b *testing.B) {
	// 假设有5个内置主题索引
	themeCount := 5

	for i := 0; i < themeCount; i++ {
		b.Run("Theme_"+Itoa(i), func(b *testing.B) {
			b.ResetTimer()

			for j := 0; j < b.N; j++ {
				pn := pixelnebula.NewPixelNebula()
				pn.WithStyle(style.GirlStyle)
				pn.WithSize(231, 231)

				_, err := pn.Generate("benchmark-theme-"+Itoa(i), false).SetTheme(i).Build().ToSVG()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkCustomTheme 测试自定义主题的性能
func BenchmarkCustomTheme(b *testing.B) {
	// 创建一个自定义主题
	customTheme := []theme.Theme{
		{
			theme.ThemePart{
				"env":   []string{"#f0f0f0", "#e0e0e0"},
				"head":  []string{"#ffd699"},
				"eyes":  []string{"#555555", "#ffffff"},
				"mouth": []string{"#ff6b6b"},
				"top":   []string{"#6b5b95", "#6b5b95"},
				"clo":   []string{"#88b04b"},
			},
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pn := pixelnebula.NewPixelNebula()
		pn.WithStyle(style.GirlStyle)
		pn.WithSize(231, 231)
		pn.WithCustomizeTheme(customTheme)

		_, err := pn.Generate("benchmark-custom-theme", false).ToSVG()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkStyleThemeCombinations 测试不同风格和主题组合的性能
func BenchmarkStyleThemeCombinations(b *testing.B) {
	styles := []style.StyleType{style.GirlStyle, style.AsianStyle}
	themes := []int{0, 1, 2}

	for _, s := range styles {
		for _, t := range themes {
			styleName := "Unknown"
			switch s {
			case style.GirlStyle:
				styleName = "Girl"
			case style.AsianStyle:
				styleName = "Asian"
			}

			b.Run(styleName+"_Theme"+Itoa(t), func(b *testing.B) {
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					pn := pixelnebula.NewPixelNebula()
					pn.WithStyle(s)
					pn.WithSize(231, 231)

					_, err := pn.Generate("benchmark-combo", false).SetTheme(t).Build().ToSVG()
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}
