package benchmark

import (
	"testing"
	"time"

	"github.com/landaiqing/go-pixelnebula"
	"github.com/landaiqing/go-pixelnebula/cache"
	"github.com/landaiqing/go-pixelnebula/style"
)

// BenchmarkDefaultCacheVsNoCache 对比有无默认缓存的性能差异
func BenchmarkDefaultCacheVsNoCache(b *testing.B) {
	// 不使用缓存的基准测试
	b.Run("NoCache", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			pn := pixelnebula.NewPixelNebula()
			pn.WithStyle(style.GirlStyle)
			pn.WithSize(231, 231)

			_, err := pn.Generate("benchmark-cache-test", false).ToSVG()
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	// 使用默认缓存的基准测试
	b.Run("DefaultCache", func(b *testing.B) {
		// 创建一个带默认缓存的实例
		pn := pixelnebula.NewPixelNebula()
		pn.WithStyle(style.GirlStyle)
		pn.WithSize(231, 231)
		pn.WithDefaultCache()

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, err := pn.Generate("benchmark-cache-test", false).ToSVG()
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// BenchmarkCacheSizes 测试不同缓存大小对性能的影响
func BenchmarkCacheSizes(b *testing.B) {
	cacheSizes := []int{10, 100, 1000}

	for _, size := range cacheSizes {
		b.Run("CacheSize_"+Itoa(size), func(b *testing.B) {
			// 创建自定义缓存配置
			pn := pixelnebula.NewPixelNebula()
			pn.WithStyle(style.GirlStyle)
			pn.WithSize(231, 231)
			pn.WithCache(cache.CacheOptions{
				Size:       size,
				Expiration: 3600 * time.Second,
			})

			// 预热缓存，生成一些不同的头像
			for i := 0; i < size/2; i++ {
				_, _ = pn.Generate("preload-"+Itoa(i), false).ToSVG()
			}

			b.ResetTimer()

			// 测试缓存命中和未命中的混合场景
			for i := 0; i < b.N; i++ {
				id := "benchmark-" + Itoa(i%size) // 循环使用ID，确保部分缓存命中
				_, err := pn.Generate(id, false).ToSVG()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkCacheCompression 测试缓存压缩对性能的影响
func BenchmarkCacheCompression(b *testing.B) {
	compressionLevels := []struct {
		name  string
		level int
	}{
		{"NoCompression", 0},
		{"LowCompression", 3},
		{"MediumCompression", 6},
		{"HighCompression", 9},
	}

	for _, cl := range compressionLevels {
		b.Run(cl.name, func(b *testing.B) {
			// 创建带压缩缓存的实例
			pn := pixelnebula.NewPixelNebula()
			pn.WithStyle(style.GirlStyle)
			pn.WithSize(231, 231)
			pn.WithDefaultCache()

			if cl.level > 0 {
				pn.WithCompression(cache.CompressOptions{
					Enabled:      true,
					Level:        cl.level,
					MinSizeBytes: 100,
				})
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_, err := pn.Generate("benchmark-compress", false).ToSVG()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkCacheExpiry 测试不同缓存过期时间的性能影响
func BenchmarkCacheExpiry(b *testing.B) {
	expiryTimes := []struct {
		name string
		time time.Duration
	}{
		{"Short_1m", 1 * time.Minute},
		{"Medium_1h", 1 * time.Hour},
		{"Long_24h", 24 * time.Hour},
	}

	for _, et := range expiryTimes {
		b.Run(et.name, func(b *testing.B) {
			// 创建带自定义过期时间的缓存实例
			pn := pixelnebula.NewPixelNebula()
			pn.WithStyle(style.GirlStyle)
			pn.WithSize(231, 231)
			pn.WithCache(cache.CacheOptions{
				Size:       100,
				Expiration: et.time,
			})

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_, err := pn.Generate("benchmark-expiry", false).ToSVG()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
