package converter

import (
	"encoding/base64"
	"regexp"
	"strings"
)

var (
	colorAttrRegex = regexp.MustCompile(`(?i)(fill|stroke):(?:#none|transparent)(?:;|\s|"|'|$)`)
)

type Converter interface {
	ToBase64() (string, error)
	ToPNG() ([]byte, error)
	ToJPEG() ([]byte, error)
}

type SVGConverter struct {
	svgData []byte
	width   int
	height  int
}

func NewSVGConverter(svgData []byte, width, height int) *SVGConverter {
	processed := preprocessSVG(svgData)
	return &SVGConverter{
		svgData: processed,
		width:   width,
		height:  height,
	}
}

func preprocessSVG(data []byte) []byte {
	// 1. 移除动画元素
	data = regexp.MustCompile(`<animateTransform[^>]*>`).ReplaceAll(data, []byte{})

	// 2. 替换 fill:#none, fill:transparent 为 fill:#000000
	processed := colorAttrRegex.ReplaceAllStringFunc(string(data), func(match string) string {
		if strings.HasPrefix(strings.ToLower(match), "fill:") {
			return "fill:#000;"
		}
		return match
	})
	return []byte(processed)
}

// ToBase64 returns the SVG data as a base64-encoded string.
func (c *SVGConverter) ToBase64() (string, error) {
	return "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString(c.svgData), nil
}

// ToPNG returns the SVG data as a PNG image.
// Note: This is not implemented yet.
// Deprecated: It can't be perfectly implemented for the time being, so it's better to Use ToBase64 instead.
func (c *SVGConverter) ToPNG() ([]byte, error) {
	// TODO: implement
	return nil, nil
}

// ToJPEG returns the SVG data as a JPEG image.
// Note: This is not implemented yet.
// Deprecated: It can't be perfectly implemented for the time being, so it's better to Use ToBase64 instead.
func (c *SVGConverter) ToJPEG() ([]byte, error) {
	// TODO: implement
	return nil, nil
}
