package benchmark

import (
	"strconv"
	"testing"

	"github.com/landaiqing/go-pixelnebula"
	"github.com/landaiqing/go-pixelnebula/style"
)

// BenchmarkBasicAvatarGeneration 测试基本头像生成性能
func BenchmarkBasicAvatarGeneration(b *testing.B) {
	// 重置计时器
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pn := pixelnebula.NewPixelNebula()
		pn.WithStyle(style.GirlStyle)
		pn.WithSize(231, 231)

		// 生成SVG
		_, err := pn.Generate("benchmark-id", false).ToSVG()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkAvatarWithNoEnvironment 测试无环境头像生成性能
func BenchmarkAvatarWithNoEnvironment(b *testing.B) {
	// 重置计时器
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pn := pixelnebula.NewPixelNebula()
		pn.WithStyle(style.GirlStyle)
		pn.WithSize(231, 231)

		// 生成无环境SVG
		_, err := pn.Generate("benchmark-id", true).ToSVG()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkDifferentSizes 测试不同大小头像生成性能
func BenchmarkDifferentSizes(b *testing.B) {
	sizes := []int{100, 200, 400, 800}

	for _, size := range sizes {
		b.Run("Size_"+Itoa(size), func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				pn := pixelnebula.NewPixelNebula()
				pn.WithStyle(style.GirlStyle)
				pn.WithSize(size, size)

				_, err := pn.Generate("benchmark-size-"+Itoa(size), false).ToSVG()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkIDReuse 测试多次使用相同ID生成头像的性能（不使用缓存）
func BenchmarkIDReuse(b *testing.B) {
	pn := pixelnebula.NewPixelNebula()
	pn.WithStyle(style.GirlStyle)
	pn.WithSize(231, 231)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := pn.Generate("fixed-benchmark-id", false).ToSVG()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Itoa 简单的整数转字符串函数
func Itoa(n int) string {
	return strconv.Itoa(n)
}
