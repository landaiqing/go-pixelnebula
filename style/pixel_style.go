package style

// PixelStyle 像素风格类型
const PixelStyle StyleType = "pixel"

// PixelStyleShapes 像素风格形状集合
var PixelStyleShapes = StyleSet{
	TypeClo:   "<path id='clo'   d=\"m60 180h120v20h-120zm0 20h20v10h-20zm100 0h20v10h-20zm-80 0h20v10h-20zm20 0h20v10h-20zm20 0h20v10h-20z\" style=\"fill:#333;\"/>",
	TypeMouth: "<path id='mouth' d=\"m100 150h30v5h-30zm5-5h20v5h-20z\" style=\"fill:#1a1a1a;\"/>",
	TypeEyes:  "<path id='eyes'  d=\"m80 100h10v10h-10zm60 0h10v10h-10z\" style=\"fill:#1a1a1a;\"/>",
	TypeTop:   "<path id='top'   d=\"m90 40h50v10h-50zm-10 10h70v10h-70zm-10 10h90v10h-90zm0 10h90v10h-90z\" style=\"fill:#333;\"/>",
	TypeHead:  "<path id='head'  d=\"m115.5 51.75a63.75 63.75 0 0 0-10.5 126.63v14.09a115.5 115.5 0 0 0-53.729 19.027 115.5 115.5 0 0 0 128.46 0 115.5 115.5 0 0 0-53.729-19.029v-14.084a63.75 63.75 0 0 0 53.25-62.881 63.75 63.75 0 0 0-63.65-63.75 63.75 63.75 0 0 0-0.09961 0z\" style=\"fill:#000;\"/>",
	TypeEnv:   "<path id='env'   d=\"M33.83,33.83a115.5,115.5,0,1,1,0,163.34,115.49,115.49,0,0,1,0-163.34Z\" style=\"fill:#01;\"/>",
}
