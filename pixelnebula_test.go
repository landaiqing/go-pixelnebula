package pixelnebula

import (
	"encoding/hex"
	"fmt"
	"github.com/landaiqing/go-pixelnebula/style"
	"os"
	"regexp"
	"testing"
)

func TestPixelNebula(t *testing.T) {
	pn := NewPixelNebula()
	numRegex := regexp.MustCompile(`[0-9]`)

	// 测试多个不同的ID
	testIDs := []string{
		"example_avatar0",
		"example_avatar1",
		"example_avatar2",
		"example_avatar3",
		"example_avatar4",
	}

	// 打印可用的风格和主题数量
	fmt.Printf("总风格数量: %d\n", pn.themeManager.StyleCount())
	for i := 0; i < pn.themeManager.StyleCount(); i++ {
		fmt.Printf("风格 %d 的主题数量: %d\n", i, pn.themeManager.ThemeCount(i))
	}

	for i, id := range testIDs {
		// 生成并保存头像
		builder := pn.Generate(id, false)
		svg, err := builder.ToSVG()
		if err != nil {
			t.Errorf("生成头像失败 (ID: %s): %v", id, err)
			continue
		}

		// 保存每个头像到不同的文件
		filename := fmt.Sprintf("avatar_%d.svg", i)
		err = os.WriteFile(filename, []byte(svg), 0644)
		if err != nil {
			t.Errorf("保存头像失败 (ID: %s): %v", id, err)
			continue
		}

		// 打印调试信息
		pn.hasher.Reset()
		pn.hasher.Write([]byte(id))
		sum := pn.hasher.Sum(nil)
		hashStr := hex.EncodeToString(sum)

		// 提取数字
		numbers := numRegex.FindAllString(hashStr, -1)
		hashNum := pn.hashToNum(numbers)

		fmt.Printf("\nID: %s\n", id)
		fmt.Printf("Hash: %s\n", hashStr)
		fmt.Printf("Numbers: %v\n", numbers)
		fmt.Printf("HashNum: %d\n", hashNum)

		// 计算并打印每个部分的索引
		parts := []string{"env", "clo", "head", "mouth", "eyes", "top"}
		for j, part := range parts {
			start := j * 2
			end := start + 2
			if end > len(numbers) {
				end = len(numbers)
			}
			partHash := numbers[start:end]
			key := pn.calcKey(partHash, nil)
			fmt.Printf("%s - StyleIndex: %d, ThemeIndex: %d\n", part, key[0], key[1])
		}
		fmt.Printf("------------------\n")
	}
}

func TestAnimation(t *testing.T) {
	pn := NewPixelNebula()

	// 设置风格和尺寸
	pn.WithStyle(style.AfrohairStyle)
	pn.WithTheme(0)

	// 1. 旋转动画 - 让环境和头部旋转
	pn.WithRotateAnimation("env", 0, 360, 10, -1) // 无限循环旋转环境

	// 2. 渐变动画  - 让环境渐变
	pn.WithGradientAnimation("env", []string{"#3498db", "#2ecc71", "#f1c40f", "#e74c3c", "#9b59b6"}, 8, -1, true)
	// 2. 渐变动画  - 让眼睛渐变
	pn.WithGradientAnimation("eyes", []string{"#3498db", "#2ecc71", "#f1c40f", "#e74c3c", "#9b59b6"}, 8, -1, true)

	// 3. 淡入淡出动画 - 让眼睛闪烁
	pn.WithFadeAnimation("eyes", "1", "0.3", 2, -1)

	// 4. 变换动画 - 让嘴巴缩放
	//pn.WithTransformAnimation("mouth", "scale", "1 1", "1.2 1.2", 1, -1)

	// 5. 颜色变换动画 - 让头发颜色变换
	pn.WithColorAnimation("top", "fill", "#9b59b6", "#e74c3c", 3, -1)
	// 5. 颜色变换动画 - 让衣服颜色变换
	pn.WithColorAnimation("clo", "fill", "#9b59b6", "#e74c3c", 3, -1)

	// 6. 弹跳动画 - 让嘴巴弹跳
	pn.WithBounceAnimation("mouth", "transform", "0,0", "0,-10", 5, 2.5, -1)
	// 6. 旋转动画 - 让嘴巴旋转
	pn.WithRotateAnimation("mouth", 0, 360, 10, -1) // 无限循环旋转环境

	//// 7. 波浪动画 - 让衣服产生波浪效果
	//pn.WithWaveAnimation("clo", 5, 0.2, "horizontal", 4, -1)

	// 8. 闪烁动画 - 让头顶装饰闪烁
	//pn.WithBlinkAnimation("head", 0.3, 1.0, 4, 6, -1)
	// 8. 波浪动画 - 让环境产生波浪效果
	//pn.WithWaveAnimation("clo", 5, 2, "horizontal", 4, -1)

	// 9. 路径动画 - 让眼睛沿着路径移动
	//pn.WithPathAnimation("eyes", "M 0,0 C 10,-10 -10,-10 0,0", 3, -1)

	pn.WithBounceAnimation("eyes", "transform", "0,0", "0,-5", 5, 2, -1)

	// 10. 带旋转的路径动画 - 让眼睛在移动的同时旋转
	//pn.WithPathAnimationRotate("mouth", "M 0,0 C 5,5 -5,5 0,0", "auto", 4, -1)

	// 生成SVG
	svg, err := pn.Generate("example_avatar", false).ToSVG()
	if err != nil {
		fmt.Printf("生成SVG失败: %v\n", err)
		os.Exit(1)
	}

	// 保存到文件
	err = os.WriteFile("./assets/example_avatar.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Printf("保存文件失败: %v\n", err)
		os.Exit(1)
	}
}

func TestDemo(t *testing.T) {

	// 创建一个新的 PixelNebula 实例
	pn := NewPixelNebula()

	// 设置风格和尺寸
	pn.WithStyle(style.GirlStyle)
	pn.WithSize(231, 231)

	// 生成 SVG 并保存到文件
	svg, err := pn.Generate("unique-id-123", false).ToSVG()
	if err != nil {
		fmt.Printf("生成 SVG 失败: %v\n", err)
		return
	}

	// 保存到文件
	err = os.WriteFile("my_avatar.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Printf("保存文件失败: %v\n", err)
		return
	}

	fmt.Println("头像成功生成: my_avatar.svg")

}

func TestRotateAnimation(t *testing.T) {
	pn := NewPixelNebula()

	// 设置风格和尺寸
	pn.WithStyle(style.FirehairStyle)
	pn.WithTheme(0)

	// 1. 旋转动画 - 让环境和头部旋转
	pn.WithRotateAnimation("eyes", 0, 360, 10, -1) // 无限循环旋转环境

	err := pn.Generate("example_avatar", false).ToFile("example_avatar.svg")
	if err != nil {
		fmt.Printf("生成 SVG 失败: %v\n", err)
		os.Exit(1)
	}
}

func TestGradientAnimation(t *testing.T) {
	pn := NewPixelNebula()

	// 设置风格和尺寸
	pn.WithStyle(style.FirehairStyle)
	pn.WithTheme(0)

	// 2. 渐变动画  - 让环境渐变
	pn.WithGradientAnimation("env", []string{"#3498db", "#2ecc71", "#f1c40f", "#e74c3c", "#9b59b6"}, 8, -1, true)

	err := pn.Generate("example_avatar", false).ToFile("example_avatar.svg")
	if err != nil {
		fmt.Printf("生成 SVG 失败: %v\n", err)
		os.Exit(1)
	}
}

// 测试淡入淡出动画
func TestFadeAnimation(t *testing.T) {
	pn := NewPixelNebula()

	// 设置风格和尺寸
	pn.WithStyle(style.FirehairStyle)
	pn.WithTheme(0)

	// 3. 淡入淡出动画 - 让眼睛闪烁
	pn.WithFadeAnimation("head", "1", "0.3", 2, -1)

	err := pn.Generate("example_avatar", false).ToFile("example_avatar.svg")
	if err != nil {
		fmt.Printf("生成 SVG 失败: %v\n", err)
		os.Exit(1)
	}
}

// 测试路径动画
func TestPathAnimation(t *testing.T) {
	pn := NewPixelNebula()

	// 设置风格和尺寸
	pn.WithStyle(style.FirehairStyle)
	pn.WithTheme(0)

	// 9. 路径动画 - 让clo沿着路径移动
	pn.WithPathAnimation("clo", "M 0,0 C 10,-10 -10,-10 0,0", 3, -1)

	err := pn.Generate("example_avatar", false).ToFile("example_avatar.svg")
	if err != nil {
		fmt.Printf("生成 SVG 失败: %v\n", err)
		os.Exit(1)
	}
}
