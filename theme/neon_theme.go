package theme

// NeonTheme 霓虹风格主题
var NeonTheme = Theme{
	// 第一部分 - 赛博朋克风格
	ThemePart{
		"env":   {"000000"},
		"clo":   {"1e1e1e", "5c5c5c", "00f2ff"},
		"head":  {"000000"},
		"mouth": {"ff00d4", "00f2ff"},
		"eyes":  {"ff00d4", "00f2ff", "ffffff", "ffffff"},
		"top":   {"1e1e1e", "00f2ff", "ff00d4"},
	},
	// 第二部分 - 霓虹黄风格
	ThemePart{
		"env":   {"0d0030"},
		"clo":   {"1e1e1e", "5c5c5c", "ffdd00"},
		"head":  {"000000"},
		"mouth": {"ffdd00", "00f2ff"},
		"eyes":  {"ffdd00", "00f2ff", "ffffff", "ffffff"},
		"top":   {"1e1e1e", "00f2ff", "ffdd00"},
	},
	// 第三部分 - 霓虹绿风格
	ThemePart{
		"env":   {"100a00"},
		"clo":   {"1e1e1e", "5c5c5c", "0cff6f"},
		"head":  {"000000"},
		"mouth": {"0cff6f", "ff00d4"},
		"eyes":  {"0cff6f", "ff00d4", "ffffff", "ffffff"},
		"top":   {"1e1e1e", "ff00d4", "0cff6f"},
	},
}
