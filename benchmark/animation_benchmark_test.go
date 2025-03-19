package benchmark

import (
	"testing"

	"github.com/landaiqing/go-pixelnebula"
	"github.com/landaiqing/go-pixelnebula/style"
)

// BenchmarkRotateAnimation 测试旋转动画的性能
func BenchmarkRotateAnimation(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pn := pixelnebula.NewPixelNebula()
		pn.WithStyle(style.GirlStyle)
		pn.WithSize(231, 231)
		pn.WithRotateAnimation("env", 0, 360, 10, 1) // 单次旋转

		_, err := pn.Generate("benchmark-rotate", false).ToSVG()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkGradientAnimation 测试渐变动画的性能
func BenchmarkGradientAnimation(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pn := pixelnebula.NewPixelNebula()
		pn.WithStyle(style.GirlStyle)
		pn.WithSize(231, 231)
		pn.WithGradientAnimation("head", []string{"#ff0000", "#00ff00", "#0000ff"}, 5, 1, true)

		_, err := pn.Generate("benchmark-gradient", false).ToSVG()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkFadeAnimation 测试淡入淡出动画的性能
func BenchmarkFadeAnimation(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pn := pixelnebula.NewPixelNebula()
		pn.WithStyle(style.GirlStyle)
		pn.WithSize(231, 231)
		pn.WithFadeAnimation("eyes", "1", "0.3", 2, 1)

		_, err := pn.Generate("benchmark-fade", false).ToSVG()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkTransformAnimation 测试变换动画的性能
func BenchmarkTransformAnimation(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pn := pixelnebula.NewPixelNebula()
		pn.WithStyle(style.GirlStyle)
		pn.WithSize(231, 231)
		pn.WithTransformAnimation("mouth", "scale", "1 1", "1.1 1.1", 1.5, 1)

		_, err := pn.Generate("benchmark-transform", false).ToSVG()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkColorAnimation 测试颜色变换动画的性能
func BenchmarkColorAnimation(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pn := pixelnebula.NewPixelNebula()
		pn.WithStyle(style.GirlStyle)
		pn.WithSize(231, 231)
		pn.WithColorAnimation("clo", "fill", "#ff0000", "#0000ff", 3, 1)

		_, err := pn.Generate("benchmark-color", false).ToSVG()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkMultipleAnimations 测试多个动画组合的性能
func BenchmarkMultipleAnimations(b *testing.B) {
	animationCounts := []int{1, 2, 3, 5}

	for _, count := range animationCounts {
		b.Run("Animations_"+Itoa(count), func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				pn := pixelnebula.NewPixelNebula()
				pn.WithStyle(style.GirlStyle)
				pn.WithSize(231, 231)

				// 根据数量添加不同的动画
				if count >= 1 {
					pn.WithRotateAnimation("env", 0, 360, 10, 1)
				}
				if count >= 2 {
					pn.WithFadeAnimation("eyes", "1", "0.3", 2, 1)
				}
				if count >= 3 {
					pn.WithTransformAnimation("mouth", "scale", "1 1", "1.1 1.1", 1.5, 1)
				}
				if count >= 4 {
					pn.WithColorAnimation("clo", "fill", "#ff0000", "#0000ff", 3, 1)
				}
				if count >= 5 {
					pn.WithGradientAnimation("head", []string{"#ff0000", "#00ff00", "#0000ff"}, 5, 1, true)
				}

				_, err := pn.Generate("benchmark-multi-"+Itoa(count), false).ToSVG()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
