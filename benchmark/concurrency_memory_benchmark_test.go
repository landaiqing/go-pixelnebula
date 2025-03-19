package benchmark

import (
	"sync"
	"testing"

	"github.com/landaiqing/go-pixelnebula"
	"github.com/landaiqing/go-pixelnebula/style"
)

// BenchmarkConcurrentGeneration 测试并发生成头像的性能
func BenchmarkConcurrentGeneration(b *testing.B) {
	concurrencyCounts := []int{1, 2, 4, 8, 16}

	for _, count := range concurrencyCounts {
		b.Run("Concurrent_"+Itoa(count), func(b *testing.B) {
			b.ResetTimer()

			// 将总迭代次数调整为b.N，确保可比较性
			b.SetParallelism(count)
			b.RunParallel(func(pb *testing.PB) {
				counter := 0
				for pb.Next() {
					counter++
					pn := pixelnebula.NewPixelNebula()
					pn.WithStyle(style.GirlStyle)
					pn.WithSize(231, 231)
					pn.WithDefaultCache()

					_, err := pn.Generate("benchmark-concurrent-"+Itoa(counter), false).ToSVG()
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		})
	}
}

// BenchmarkConcurrentWithSharedInstance 测试使用共享实例进行并发生成的性能
func BenchmarkConcurrentWithSharedInstance(b *testing.B) {
	concurrencyCounts := []int{1, 2, 4, 8, 16}

	for _, count := range concurrencyCounts {
		b.Run("SharedInstance_"+Itoa(count), func(b *testing.B) {
			// 创建一个共享实例
			pn := pixelnebula.NewPixelNebula()
			pn.WithStyle(style.GirlStyle)
			pn.WithSize(231, 231)
			pn.WithDefaultCache()

			// 创建互斥锁保护共享实例
			var mu sync.Mutex

			b.ResetTimer()
			b.SetParallelism(count)
			b.RunParallel(func(pb *testing.PB) {
				counter := 0
				for pb.Next() {
					counter++
					// 锁定共享实例
					mu.Lock()
					_, err := pn.Generate("benchmark-shared-"+Itoa(counter), false).ToSVG()
					mu.Unlock()
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		})
	}
}

// BenchmarkMemoryUsage 测试不同操作的内存使用情况
func BenchmarkMemoryUsage(b *testing.B) {
	// 注意: 这个基准测试主要关注内存分配统计，
	// Go 的基准测试框架会自动收集并报告内存统计数据

	// 测试基本头像生成的内存使用
	b.Run("BasicGeneration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			pn := pixelnebula.NewPixelNebula()
			pn.WithStyle(style.GirlStyle)
			pn.WithSize(231, 231)
			_, err := pn.Generate("memory-basic", false).ToSVG()
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	// 测试添加动画的内存使用
	b.Run("WithAnimations", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			pn := pixelnebula.NewPixelNebula()
			pn.WithStyle(style.GirlStyle)
			pn.WithSize(231, 231)
			pn.WithRotateAnimation("env", 0, 360, 10, 1)
			pn.WithFadeAnimation("eyes", "1", "0.3", 2, 1)
			_, err := pn.Generate("memory-animations", false).ToSVG()
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	// 测试缓存的内存使用
	b.Run("WithCache", func(b *testing.B) {
		pn := pixelnebula.NewPixelNebula()
		pn.WithStyle(style.GirlStyle)
		pn.WithSize(231, 231)
		pn.WithDefaultCache()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := pn.Generate("memory-cache", false).ToSVG()
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	// 测试大尺寸头像的内存使用
	b.Run("LargeSize", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			pn := pixelnebula.NewPixelNebula()
			pn.WithStyle(style.GirlStyle)
			pn.WithSize(1000, 1000)
			_, err := pn.Generate("memory-large", false).ToSVG()
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
