package benchmark

import (
	"fmt"
	"testing"

	"github.com/landaiqing/go-pixelnebula"
	"github.com/landaiqing/go-pixelnebula/style"
)

// BenchmarkBatchGeneration 测试不同批量大小的SVG生成性能
func BenchmarkBatchGeneration(b *testing.B) {
	batchSizes := []int{10, 50, 100, 500}

	for _, size := range batchSizes {
		// 顺序生成的测试
		b.Run(fmt.Sprintf("Sequential_Size_%d", size), func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				pn := pixelnebula.NewPixelNebula()
				pn.WithStyle(style.GirlStyle)
				pn.WithSize(231, 231)
				pn.WithParallelRender(false) // 禁用并行渲染

				// 准备ID列表
				ids := generateIDs(size)

				// 执行批量生成
				_, err := pn.GenerateBatch(ids, false, nil)
				if err != nil {
					b.Fatal(err)
				}
			}
		})

		// 并行生成的测试
		b.Run(fmt.Sprintf("Parallel_Size_%d", size), func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				pn := pixelnebula.NewPixelNebula()
				pn.WithStyle(style.GirlStyle)
				pn.WithSize(231, 231)
				pn.WithParallelRender(true) // 启用并行渲染

				// 准备ID列表
				ids := generateIDs(size)

				// 执行批量生成
				_, err := pn.GenerateBatch(ids, false, nil)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkBatchWithDifferentConcurrencyLevels 测试不同并发级别的批量生成性能
func BenchmarkBatchWithDifferentConcurrencyLevels(b *testing.B) {
	batchSize := 100
	concurrencyLevels := []int{2, 4, 8, 16, 32}

	for _, level := range concurrencyLevels {
		b.Run(fmt.Sprintf("ConcurrencyLevel_%d", level), func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				pn := pixelnebula.NewPixelNebula()
				pn.WithStyle(style.GirlStyle)
				pn.WithSize(231, 231)
				pn.WithParallelRender(true)   // 启用并行渲染
				pn.WithConcurrencyPool(level) // 设置并发级别

				// 准备ID列表
				ids := generateIDs(batchSize)

				// 执行批量生成
				_, err := pn.GenerateBatch(ids, false, nil)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkBatchToBase64 测试批量转换为Base64的性能
func BenchmarkBatchToBase64(b *testing.B) {
	batchSizes := []int{10, 50, 100}

	for _, size := range batchSizes {
		// 顺序转换
		b.Run(fmt.Sprintf("Sequential_Size_%d", size), func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				pn := pixelnebula.NewPixelNebula()
				pn.WithStyle(style.GirlStyle)
				pn.WithSize(231, 231)
				pn.WithParallelRender(false) // 禁用并行渲染

				// 准备ID列表
				ids := generateIDs(size)

				// 执行批量生成和转换
				_, err := pn.GenerateBatchBase64(ids, false, nil, 231, 231)
				if err != nil {
					b.Fatal(err)
				}
			}
		})

		// 并行转换
		b.Run(fmt.Sprintf("Parallel_Size_%d", size), func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				pn := pixelnebula.NewPixelNebula()
				pn.WithStyle(style.GirlStyle)
				pn.WithSize(231, 231)
				pn.WithParallelRender(true) // 启用并行渲染

				// 准备ID列表
				ids := generateIDs(size)

				// 执行批量生成和转换
				_, err := pn.GenerateBatchBase64(ids, false, nil, 231, 231)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkBatchWithCache 测试使用缓存的批量生成性能
func BenchmarkBatchWithCache(b *testing.B) {
	batchSize := 100

	b.Run("WithoutCache", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			pn := pixelnebula.NewPixelNebula()
			pn.WithStyle(style.GirlStyle)
			pn.WithSize(231, 231)
			pn.WithParallelRender(true) // 启用并行渲染

			// 每次使用固定的ID列表
			ids := generateFixedIDs(batchSize)

			// 执行批量生成
			_, err := pn.GenerateBatch(ids, false, nil)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("WithCache", func(b *testing.B) {
		// 创建一个使用缓存的实例
		pn := pixelnebula.NewPixelNebula()
		pn.WithStyle(style.GirlStyle)
		pn.WithSize(231, 231)
		pn.WithParallelRender(true) // 启用并行渲染
		pn.WithDefaultCache()       // 启用缓存

		// 使用固定ID列表
		ids := generateFixedIDs(batchSize)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			// 执行批量生成（由于固定ID，将使用缓存）
			_, err := pn.GenerateBatch(ids, false, nil)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// BenchmarkSaveToFiles 测试批量保存文件的性能
func BenchmarkSaveToFiles(b *testing.B) {
	// 跳过实际文件写入，仅测试内存操作部分
	b.Skip("Skipping file I/O benchmark to avoid disk writes")

	batchSizes := []int{10, 50, 100}

	for _, size := range batchSizes {
		// 顺序保存
		b.Run(fmt.Sprintf("Sequential_Size_%d", size), func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				pn := pixelnebula.NewPixelNebula()
				pn.WithStyle(style.GirlStyle)
				pn.WithSize(231, 231)
				pn.WithParallelRender(false) // 禁用并行渲染

				// 准备ID列表
				ids := generateIDs(size)

				// 执行批量保存（使用临时目录）
				_, err := pn.SaveBatchToFiles(ids, false, nil, "/tmp/benchmark/avatar_%s.svg")
				if err != nil {
					b.Fatal(err)
				}
			}
		})

		// 并行保存
		b.Run(fmt.Sprintf("Parallel_Size_%d", size), func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				pn := pixelnebula.NewPixelNebula()
				pn.WithStyle(style.GirlStyle)
				pn.WithSize(231, 231)
				pn.WithParallelRender(true) // 启用并行渲染

				// 准备ID列表
				ids := generateIDs(size)

				// 执行批量保存（使用临时目录）
				_, err := pn.SaveBatchToFiles(ids, false, nil, "/tmp/benchmark/avatar_%s.svg")
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkMemoryUsageBatch 测试批量生成的内存使用情况
func BenchmarkMemoryUsageBatch(b *testing.B) {
	batchSize := 100

	// 测试顺序生成的内存使用
	b.Run("SequentialBatch", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			pn := pixelnebula.NewPixelNebula()
			pn.WithStyle(style.GirlStyle)
			pn.WithSize(231, 231)
			pn.WithParallelRender(false) // 禁用并行渲染

			// 准备ID列表
			ids := generateIDs(batchSize)

			// 执行批量生成
			_, err := pn.GenerateBatch(ids, false, nil)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	// 测试并行生成的内存使用
	b.Run("ParallelBatch", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			pn := pixelnebula.NewPixelNebula()
			pn.WithStyle(style.GirlStyle)
			pn.WithSize(231, 231)
			pn.WithParallelRender(true) // 启用并行渲染

			// 准备ID列表
			ids := generateIDs(batchSize)

			// 执行批量生成
			_, err := pn.GenerateBatch(ids, false, nil)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// 生成指定数量的随机ID
func generateIDs(count int) []string {
	ids := make([]string, count)
	for i := 0; i < count; i++ {
		ids[i] = fmt.Sprintf("batch-id-%d", i)
	}
	return ids
}

// 生成固定的ID集合，用于缓存测试
func generateFixedIDs(count int) []string {
	ids := make([]string, count)
	for i := 0; i < count; i++ {
		ids[i] = fmt.Sprintf("fixed-id-%d", i)
	}
	return ids
}
