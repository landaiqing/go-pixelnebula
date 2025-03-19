package theme

import (
	"github.com/landaiqing/go-pixelnebula/errors"
)

// ColorScheme 表示一个颜色方案，包含多个颜色
type ColorScheme []string

// ThemePart 表示主题的一个部分，包含多个组件的颜色方案
type ThemePart map[string]ColorScheme

// Theme 表示一个完整的主题，包含多个部分
type Theme []ThemePart

// Manager 主题管理器，负责管理所有主题
type Manager struct {
	themes []Theme
}

// NewThemeManager 创建一个新的主题管理器
func NewThemeManager() *Manager {
	m := &Manager{}
	m.initThemes()
	return m
}

// GetTheme 获取指定索引的主题
func (m *Manager) GetTheme(themeIndex, partIndex int) (ThemePart, error) {
	if themeIndex < 0 || themeIndex >= len(m.themes) {
		return nil, errors.ErrInvalidTheme
	}

	theme := m.themes[themeIndex]
	if partIndex < 0 || partIndex >= len(theme) {
		return nil, errors.ErrInvalidPart
	}

	return theme[partIndex], nil
}

// StyleCount 返回主题数量
func (m *Manager) StyleCount() int {
	return len(m.themes)
}

// ThemeCount 返回指定主题的部分数量
func (m *Manager) ThemeCount(themeIndex int) int {
	if themeIndex < 0 || themeIndex >= len(m.themes) {
		return 0
	}
	return len(m.themes[themeIndex])
}

// AddTheme 添加一个新主题
func (m *Manager) AddTheme(theme Theme) int {
	m.themes = append(m.themes, theme)
	return len(m.themes) - 1
}

// CustomizeTheme 自定义主题
func (m *Manager) CustomizeTheme(theme []Theme) {
	m.themes = theme
}

// GetThemeCountByStyle 获取指定风格索引下的主题数量
func (m *Manager) GetThemeCountByStyle(styleIndex int) int {
	if styleIndex < 0 || styleIndex >= len(m.themes) {
		return 0
	}
	return len(m.themes[styleIndex])
}
