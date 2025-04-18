package cache

import (
	"bytes"
	"compress/gzip"
	"io"
	"strings"
	"sync"
)

var (
	// 字符串构建器对象池
	optimizeBuilderPool = sync.Pool{
		New: func() interface{} {
			return new(strings.Builder)
		},
	}

	// Gzip Writer对象池
	gzipWriterPool = sync.Pool{
		New: func() interface{} {
			writer, _ := gzip.NewWriterLevel(nil, gzip.BestCompression)
			return writer
		},
	}

	// Gzip Reader对象池
	gzipReaderPool = sync.Pool{
		New: func() interface{} {
			return new(gzip.Reader)
		},
	}

	// 字节缓冲区对象池
	bytesBufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

// CompressOptions 压缩配置选项
type CompressOptions struct {
	Enabled          bool    // 是否启用压缩
	Level            int     // 压缩级别，范围从-2(不压缩)到9(最高压缩)
	MinSize          int     // 最小压缩大小，小于此大小的SVG不压缩
	CompressionRatio float64 // 最小压缩比率，压缩后大小/原始大小，小于此比率才使用压缩结果
}

// DefaultCompressOptions 默认压缩配置
var DefaultCompressOptions = CompressOptions{
	Enabled:          true,
	Level:            9,    // 默认使用最高压缩级别
	MinSize:          1024, // 默认最小压缩大小为1KB
	CompressionRatio: 0.8,  // 默认最小压缩比率为0.8
}

// CompressSVG 压缩SVG字符串
func CompressSVG(svg string, options CompressOptions) ([]byte, bool) {
	if !options.Enabled {
		return []byte(svg), false
	}

	// 如果SVG太小，不进行压缩
	if len(svg) < options.MinSize {
		return []byte(svg), false
	}

	// 从对象池获取缓冲区
	buf := bytesBufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bytesBufferPool.Put(buf)

	// 从对象池获取gzip写入器
	writer := gzipWriterPool.Get().(*gzip.Writer)
	writer.Reset(buf)
	defer gzipWriterPool.Put(writer)

	// 写入数据
	if _, err := writer.Write([]byte(svg)); err != nil {
		return []byte(svg), false
	}

	// 关闭写入器
	if err := writer.Close(); err != nil {
		return []byte(svg), false
	}

	// 获取压缩后的数据
	compressed := make([]byte, buf.Len())
	copy(compressed, buf.Bytes())

	// 计算压缩比率
	ratio := float64(len(compressed)) / float64(len(svg))
	if ratio > options.CompressionRatio {
		// 压缩效果不好，返回原始数据
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

	// 获取缓冲区
	buf := bytesBufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bytesBufferPool.Put(buf)

	// 创建Reader
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return string(data), err
	}
	defer reader.Close()

	// 读取解压缩后的数据
	if _, err := io.Copy(buf, reader); err != nil {
		return string(data), err
	}

	return buf.String(), nil
}

// isGzipped 检查数据是否为gzip格式
func isGzipped(data []byte) bool {
	// gzip文件的魔数是0x1f 0x8b
	return len(data) > 2 && data[0] == 0x1f && data[1] == 0x8b
}

// OptimizeSVG 优化SVG字符串，移除不必要的空白和注释
func OptimizeSVG(svg string) string {
	// 如果SVG太小，不进行优化
	if len(svg) < 100 {
		return svg
	}

	// 移除XML注释
	svg = removeXMLComments(svg)

	// 移除多余的空白
	svg = removeExtraWhitespace(svg)

	return svg
}

// removeXMLComments 移除XML注释 - 优化版本
func removeXMLComments(svg string) string {
	// 获取构建器
	sb := optimizeBuilderPool.Get().(*strings.Builder)
	sb.Reset()
	sb.Grow(len(svg))
	defer optimizeBuilderPool.Put(sb)

	for i := 0; i < len(svg); {
		// 检查是否遇到注释开始
		if i+3 < len(svg) && svg[i] == '<' && svg[i+1] == '!' && svg[i+2] == '-' && svg[i+3] == '-' {
			// 查找注释结束
			end := i + 4
			for end+2 <= len(svg) {
				if svg[end] == '-' && svg[end+1] == '-' && svg[end+2] == '>' {
					end += 3
					break
				}
				end++
			}
			// 跳过注释
			i = end
		} else {
			// 添加非注释字符
			sb.WriteByte(svg[i])
			i++
		}
	}

	return sb.String()
}

// removeExtraWhitespace 移除多余的空白 - 优化版本
func removeExtraWhitespace(svg string) string {
	// 获取构建器
	sb := optimizeBuilderPool.Get().(*strings.Builder)
	sb.Reset()
	sb.Grow(len(svg))
	defer optimizeBuilderPool.Put(sb)

	inTag := false
	inQuote := false
	quoteChar := byte(0)
	lastWasSpace := false

	for i := 0; i < len(svg); i++ {
		c := svg[i]

		// 处理引号内的内容
		if inQuote {
			sb.WriteByte(c)
			if c == quoteChar {
				inQuote = false
			}
			continue
		}

		// 处理标签
		if c == '<' {
			inTag = true
			lastWasSpace = false
			sb.WriteByte(c)
			continue
		}

		if c == '>' {
			inTag = false
			lastWasSpace = false
			sb.WriteByte(c)
			continue
		}

		// 处理引号开始
		if (c == '"' || c == '\'') && inTag {
			inQuote = true
			quoteChar = c
			sb.WriteByte(c)
			continue
		}

		// 处理空白字符
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			// 压缩连续空白
			if !lastWasSpace && (inTag || i > 0 && svg[i-1] != '>') {
				sb.WriteByte(' ')
				lastWasSpace = true
			}
			continue
		}

		// 处理普通字符
		lastWasSpace = false
		sb.WriteByte(c)
	}

	return sb.String()
}
