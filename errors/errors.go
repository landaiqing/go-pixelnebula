package errors

import "errors"

// 定义错误常量
var (
	ErrAvatarIDRequired     = errors.New("pixelnebula: avatar id is required")
	ErrInvalidTheme         = errors.New("pixelnebula: invalid theme index")
	ErrInvalidPart          = errors.New("pixelnebula: invalid part index")
	ErrInvalidShapeSetIndex = errors.New("pixelnebula: invalid shape set index")
	ErrInvalidShapeType     = errors.New("pixelnebula: invalid shape type")
	ErrInvalidColor         = errors.New("pixelnebula: invalid color scheme")
	ErrInsufficientHash     = errors.New("pixelnebula: insufficient hash digits generated")
	ErrInvalidStyleName     = errors.New("pixelnebula: invalid style name")
)
