package main

import (
	"fmt"
	"os"
	"time"

	"github.com/landaiqing/go-pixelnebula"
	"github.com/landaiqing/go-pixelnebula/cache"
	"github.com/landaiqing/go-pixelnebula/style"
)

// 缓存系统示例
// 展示如何使用PixelNebula的缓存功能
func main() {
	// 1. 使用默认缓存
	fmt.Println("=== 使用默认缓存示例 ===")
	defaultCacheExample()

	// 2. 使用自定义缓存
	fmt.Println("\n=== 使用自定义缓存示例 ===")
	customCacheExample()

	// 3. 使用带监控的缓存
	fmt.Println("\n=== 使用带监控的缓存示例 ===")
	monitoredCacheExample()

	// 4. 使用压缩缓存
	fmt.Println("\n=== 使用压缩缓存示例 ===")
	compressedCacheExample()
}

// 使用默认缓存示例
func defaultCacheExample() {
	// 创建一个带默认缓存的PixelNebula实例
	pn := pixelnebula.NewPixelNebula().WithDefaultCache()

	// 设置基本属性
	pn.WithStyle(style.AfrohairStyle)
	pn.WithTheme(0)
	pn.WithSize(200, 200)

	// 第一次生成头像 - 会存入缓存
	startTime1 := time.Now()
	_, err := pn.Generate("default-cache-example", false).ToSVG()
	if err != nil {
		fmt.Printf("生成SVG失败: %v\n", err)
		return
	}
	duration1 := time.Since(startTime1)

	// 第二次生成相同头像 - 应该从缓存中获取
	startTime2 := time.Now()
	svg2, err := pn.Generate("default-cache-example", false).ToSVG()
	if err != nil {
		fmt.Printf("从缓存生成SVG失败: %v\n", err)
		return
	}
	duration2 := time.Since(startTime2)

	// 保存第二次生成的头像
	err = os.WriteFile("default_cache.svg", []byte(svg2), 0644)
	if err != nil {
		fmt.Printf("保存缓存生成的SVG文件失败: %v\n", err)
		return
	}

	fmt.Printf("第一次生成耗时: %v\n", duration1)
	fmt.Printf("第二次生成耗时: %v (使用缓存)\n", duration2)
	fmt.Printf("性能提升: %.2f倍\n", float64(duration1)/float64(duration2))
	fmt.Println("成功生成带默认缓存的头像: default_cache.svg")
}

// 使用自定义缓存示例
func customCacheExample() {
	// 创建自定义缓存选项
	customCacheOptions := cache.CacheOptions{
		Enabled:    true,
		Size:       100,           // 最大缓存条目数
		Expiration: 1 * time.Hour, // 缓存有效期
	}

	// 创建一个带自定义缓存的PixelNebula实例
	pn := pixelnebula.NewPixelNebula().WithCache(customCacheOptions)

	// 设置基本属性
	pn.WithStyle(style.GirlStyle)
	pn.WithTheme(1)
	pn.WithSize(200, 200)

	// 第一次生成头像 - 会存入缓存
	startTime1 := time.Now()
	_, err := pn.Generate("custom-cache-example", false).ToSVG()
	if err != nil {
		fmt.Printf("生成SVG失败: %v\n", err)
		return
	}
	duration1 := time.Since(startTime1)

	// 第二次生成相同头像 - 应该从缓存中获取
	startTime2 := time.Now()
	svg2, err := pn.Generate("custom-cache-example", false).ToSVG()
	if err != nil {
		fmt.Printf("从缓存生成SVG失败: %v\n", err)
		return
	}
	duration2 := time.Since(startTime2)

	// 保存第二次生成的头像
	err = os.WriteFile("custom_cache.svg", []byte(svg2), 0644)
	if err != nil {
		fmt.Printf("保存缓存生成的SVG文件失败: %v\n", err)
		return
	}

	fmt.Printf("第一次生成耗时: %v\n", duration1)
	fmt.Printf("第二次生成耗时: %v (使用缓存)\n", duration2)
	fmt.Printf("性能提升: %.2f倍\n", float64(duration1)/float64(duration2))
	fmt.Println("成功生成带自定义缓存的头像: custom_cache.svg")
}

// 使用带监控的缓存示例
func monitoredCacheExample() {
	// 创建带监控的缓存选项
	monitorOptions := cache.MonitorOptions{
		Enabled:        true,
		SampleInterval: 5 * time.Second,
	}

	// 创建一个带默认缓存和监控的PixelNebula实例
	pn := pixelnebula.NewPixelNebula().WithDefaultCache().WithMonitoring(monitorOptions)

	// 设置基本属性
	pn.WithStyle(style.AsianStyle)
	pn.WithTheme(0)
	pn.WithSize(200, 200)

	// 生成多个头像以展示监控效果
	for i := 0; i < 5; i++ {
		uniqueID := fmt.Sprintf("monitor-example-%d", i)
		svg, err := pn.Generate(uniqueID, false).ToSVG()
		if err != nil {
			fmt.Printf("生成SVG失败: %v\n", err)
			continue
		}

		// 重复生成相同头像以测试缓存命中
		for j := 0; j < 3; j++ {
			_, err = pn.Generate(uniqueID, false).ToSVG()
			if err != nil {
				fmt.Printf("从缓存生成SVG失败: %v\n", err)
			}
		}

		// 保存最后一个头像
		if i == 4 {
			err = os.WriteFile("monitored_cache.svg", []byte(svg), 0644)
			if err != nil {
				fmt.Printf("保存带监控缓存的SVG文件失败: %v\n", err)
			} else {
				fmt.Println("成功生成带监控缓存的头像: monitored_cache.svg")
			}
		}
	}

	// 等待监控报告生成
	fmt.Println("等待监控报告生成...")
	time.Sleep(6 * time.Second)
}

// 使用压缩缓存示例
func compressedCacheExample() {
	// 创建压缩选项
	compressOptions := cache.CompressOptions{
		Enabled: true,
		Level:   6,
		MinSize: 100, // 最小压缩大小 (字节)
	}

	// 创建一个带默认缓存和压缩的PixelNebula实例
	pn := pixelnebula.NewPixelNebula().WithDefaultCache().WithCompression(compressOptions)

	// 设置基本属性
	pn.WithStyle(style.AfrohairStyle)
	pn.WithTheme(0)
	pn.WithSize(300, 300)

	// 添加一些动画以增加SVG大小
	pn.WithRotateAnimation("env", 0, 360, 10, -1)
	pn.WithGradientAnimation("env", []string{"#3498db", "#2ecc71", "#f1c40f", "#e74c3c", "#9b59b6"}, 8, -1, true)
	pn.WithFadeAnimation("eyes", "1", "0.3", 2, -1)
	pn.WithTransformAnimation("mouth", "scale", "1 1", "1.2 1.2", 1, -1)

	// 第一次生成头像 - 会存入压缩缓存
	startTime1 := time.Now()
	_, err := pn.Generate("compress-cache-example", false).ToSVG()
	if err != nil {
		fmt.Printf("生成SVG失败: %v\n", err)
		return
	}
	duration1 := time.Since(startTime1)

	// 第二次生成相同头像 - 应该从压缩缓存中获取
	startTime2 := time.Now()
	svg2, err := pn.Generate("compress-cache-example", false).ToSVG()
	if err != nil {
		fmt.Printf("从压缩缓存生成SVG失败: %v\n", err)
		return
	}
	duration2 := time.Since(startTime2)

	// 保存第二次生成的头像
	err = os.WriteFile("compressed_cache.svg", []byte(svg2), 0644)
	if err != nil {
		fmt.Printf("保存压缩缓存的SVG文件失败: %v\n", err)
		return
	}

	fmt.Printf("第一次生成耗时: %v\n", duration1)
	fmt.Printf("第二次生成耗时: %v (使用压缩缓存)\n", duration2)
	fmt.Printf("性能提升: %.2f倍\n", float64(duration1)/float64(duration2))
	fmt.Println("成功生成带压缩缓存的头像: compressed_cache.svg")
}
