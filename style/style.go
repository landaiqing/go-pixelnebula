package style

import "github.com/landaiqing/go-pixelnebula/errors"

// ShapeType 表示形状类型
type ShapeType string

// 预定义形状类型
const (
	TypeClo   ShapeType = "clo"   // 衣服
	TypeMouth ShapeType = "mouth" // 嘴巴
	TypeEyes  ShapeType = "eyes"  // 眼睛
	TypeTop   ShapeType = "top"   // 头顶
	TypeHead  ShapeType = "head"  // 头部
	TypeEnv   ShapeType = "env"   // 环境/背景
)

// 预定义风格类型常量
const (
	RoboStyle         StyleType = "robo"
	GirlStyle         StyleType = "girl"
	BlondeStyle       StyleType = "blonde"
	GuyStyle          StyleType = "guy"
	CountryStyle      StyleType = "country"
	GeeknotStyle      StyleType = "geeknot"
	AsianStyle        StyleType = "asian"
	PunkStyle         StyleType = "punk"
	AfrohairStyle     StyleType = "afrohair"
	NormieFemaleStyle StyleType = "normiefemale"
	OlderStyle        StyleType = "older"
	FirehairStyle     StyleType = "firehair"
	BlondStyle        StyleType = "blond"
	AteamStyle        StyleType = "ateam"
	RastaStyle        StyleType = "rasta"
	MetaStyle         StyleType = "meta"
	SquareStyle       StyleType = "square"
)

// StyleType 表示风格类型
type StyleType string

// StyleSet 表示一组形状
type StyleSet map[ShapeType]string

// Manager 形状管理器，负责管理所有形状
type Manager struct {
	styleSets []StyleSet
}

// NewShapeManager 创建一个新的形状管理器
func NewShapeManager() *Manager {
	m := &Manager{}
	m.initShapes()
	return m
}

// GetShape 获取指定索引和类型的形状
func (m *Manager) GetShape(setIndex int, shapeType ShapeType) (string, error) {
	if setIndex < 0 || setIndex >= len(m.styleSets) {
		return "", errors.ErrInvalidShapeSetIndex
	}

	shapeSet := m.styleSets[setIndex]
	shape, ok := shapeSet[shapeType]
	if !ok {
		return "", errors.ErrInvalidShapeType
	}

	return shape, nil
}

// StyleSetCount 返回形状集合数量
func (m *Manager) StyleSetCount() int {
	return len(m.styleSets)
}

// AddStyleSet 添加一个新形状集合
func (m *Manager) AddStyleSet(shapeSet StyleSet) int {
	m.styleSets = append(m.styleSets, shapeSet)
	return len(m.styleSets) - 1
}

// CustomizeStyle 自定义风格
func (m *Manager) CustomizeStyle(styleSets []StyleSet) {
	m.styleSets = styleSets
}

// GetStyleIndex 根据风格类型获取对应的索引值
func (m *Manager) GetStyleIndex(style StyleType) (int, error) {
	// 遍历已初始化的风格列表获取索引
	for i, s := range []StyleType{
		RoboStyle,
		GirlStyle,
		BlondeStyle,
		GuyStyle,
		CountryStyle,
		GeeknotStyle,
		AsianStyle,
		PunkStyle,
		AfrohairStyle,
		NormieFemaleStyle,
		OlderStyle,
		FirehairStyle,
		BlondStyle,
		AteamStyle,
		RastaStyle,
		MetaStyle,
		SquareStyle,
	} {
		if s == style {
			return i, nil
		}
	}
	return -1, errors.ErrInvalidStyleName
}
