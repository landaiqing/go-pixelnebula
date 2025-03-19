package cache

import (
	"bytes"
	"compress/gzip"
	"io"
	"strings"
)

// CompressOptions 压缩选项
type CompressOptions struct {
	Enabled      bool    // 是否启用压缩
	Level        int     // 压缩级别 (1-9)，1为最快压缩，9为最佳压缩
	MinSizeBytes int     // 最小压缩大小，小于此大小的数据不进行压缩
	Ratio        float64 // 压缩比阈值，压缩后大小/原始大小，小于此值才保存压缩结果
}

// DefaultCompressOptions 默认压缩选项
var DefaultCompressOptions = CompressOptions{
	Enabled:      true,
	Level:        6,   // 默认压缩级别为6，平衡压缩率和性能
	MinSizeBytes: 100, // 默认最小压缩大小为100字节
	Ratio:        0.9, // 默认压缩比阈值为0.9，即至少要压缩到原始大小的90%以下才保存压缩结果
}

// CompressSVG 压缩SVG数据
// 返回压缩后的数据和是否进行了压缩
func CompressSVG(svg string, options CompressOptions) ([]byte, bool) {
	if !options.Enabled || len(svg) < options.MinSizeBytes {
		return []byte(svg), false
	}

	// 创建一个bytes.Buffer来存储压缩数据
	var buf bytes.Buffer

	// 创建一个gzip.Writer，设置压缩级别
	writer, err := gzip.NewWriterLevel(&buf, options.Level)
	if err != nil {
		return []byte(svg), false
	}

	// 写入SVG数据
	_, err = writer.Write([]byte(svg))
	if err != nil {
		return []byte(svg), false
	}

	// 关闭writer，确保所有数据都被写入
	err = writer.Close()
	if err != nil {
		return []byte(svg), false
	}

	// 获取压缩后的数据
	compressed := buf.Bytes()

	// 计算压缩比
	ratio := float64(len(compressed)) / float64(len(svg))

	// 如果压缩比不理想，返回原始数据
	if ratio >= options.Ratio {
		return []byte(svg), false
	}

	return compressed, true
}

// DecompressSVG 解压缩SVG数据
func DecompressSVG(data []byte, isCompressed bool) (string, error) {
	// 如果数据未压缩，直接返回字符串
	if !isCompressed {
		return string(data), nil
	}

	// 检查数据是否为gzip格式
	if !isGzipped(data) {
		return string(data), nil
	}

	// 创建一个gzip.Reader
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return string(data), err
	}
	defer reader.Close()

	// 读取解压缩后的数据
	decompressed, err := io.ReadAll(reader)
	if err != nil {
		return string(data), err
	}

	return string(decompressed), nil
}

// isGzipped 检查数据是否为gzip格式
func isGzipped(data []byte) bool {
	// gzip文件的魔数是0x1f 0x8b
	return len(data) > 2 && data[0] == 0x1f && data[1] == 0x8b
}

// OptimizeSVG 优化SVG字符串，移除不必要的空白和注释
func OptimizeSVG(svg string) string {
	// 移除XML注释
	svg = removeXMLComments(svg)

	// 移除多余的空白
	svg = removeExtraWhitespace(svg)

	return svg
}

// removeXMLComments 移除XML注释
func removeXMLComments(svg string) string {
	for {
		start := strings.Index(svg, "<!--")
		if start == -1 {
			break
		}
		end := strings.Index(svg[start:], "-->") + start
		if end > start {
			svg = svg[:start] + svg[end+3:]
		} else {
			break
		}
	}
	return svg
}

// removeExtraWhitespace 移除多余的空白
func removeExtraWhitespace(svg string) string {
	// 替换多个空白字符为单个空格
	svg = strings.Join(strings.Fields(svg), " ")

	// 优化常见的SVG标签周围的空白
	svg = strings.ReplaceAll(svg, "> <", "><")
	svg = strings.ReplaceAll(svg, " />", "/>")
	svg = strings.ReplaceAll(svg, " =", "=")
	svg = strings.ReplaceAll(svg, "= ", "=")

	return svg
}
