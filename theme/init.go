package theme

import "github.com/landaiqing/go-pixelnebula/style"

var defaultThemeSet = map[style.StyleType]Theme{
	style.RoboStyle: {
		ThemePart{
			"env":   {"ff2f2b"},
			"clo":   {"fff", "000"},
			"head":  {"fff"},
			"mouth": {"fff", "000", "000"},
			"eyes":  {"000", "none", "0ff"},
			"top":   {"fff", "fff"},
		},
		// 第二部分
		ThemePart{
			"env":   {"ff1ec1"},
			"clo":   {"000", "fff"},
			"head":  {"ffc1c1"},
			"mouth": {"fff", "000", "000"},
			"eyes":  {"FF2D00", "fff", "none"},
			"top":   {"a21d00", "fff"},
		},
		// 第三部分
		ThemePart{
			"env":   {"0079b1"},
			"clo":   {"0e00b1", "d1fffe"},
			"head":  {"f5aa77"},
			"mouth": {"fff", "000", "000"},
			"eyes":  {"0c00de", "fff", "none"},
			"top":   {"acfffd", "acfffd"},
		},
	},

	// 创建Girl主题
	style.GirlStyle: {
		// 第一部分
		ThemePart{
			"env":   {"a50000"},
			"clo":   {"f06", "8e0039"},
			"head":  {"85492C"},
			"mouth": {"000"},
			"eyes":  {"000", "ff9809"},
			"top":   {"ff9809", "ff9809", "none", "none"},
		},
		// 第二部分
		ThemePart{
			"env":   {"40E83B"},
			"clo":   {"00650b", "62ce5a"},
			"head":  {"f7c1a6"},
			"mouth": {"6e1c1c"},
			"eyes":  {"000", "ff833b"},
			"top":   {"67FFCC", "none", "none", "ecff3b"},
		},
		// 第三部分
		ThemePart{
			"env":   {"ff2c2c"},
			"clo":   {"fff", "000"},
			"head":  {"ffce8b"},
			"mouth": {"000"},
			"eyes":  {"000", "ff9809"},
			"top":   {"ff9809", "ff9809", "none", "none"},
		},
	},

	// 创建Blonde主题
	style.BlondeStyle: {
		// 第一部分
		ThemePart{
			"env":   {"00aad4"},
			"clo":   {"fff", "000"},
			"head":  {"ffe0bd"},
			"mouth": {"ff9a84"},
			"eyes":  {"000", "fff"},
			"top":   {"fff200", "fff200"},
		},
		// 第二部分
		ThemePart{
			"env":   {"00aad4"},
			"clo":   {"fff", "000"},
			"head":  {"ffe0bd"},
			"mouth": {"ff9a84"},
			"eyes":  {"000", "fff"},
			"top":   {"fff200", "fff200"},
		},
		// 第三部分
		ThemePart{
			"env":   {"00aad4"},
			"clo":   {"fff", "000"},
			"head":  {"ffe0bd"},
			"mouth": {"ff9a84"},
			"eyes":  {"000", "fff"},
			"top":   {"fff200", "fff200"},
		},
	},

	// Guy 主题
	style.GuyStyle: {
		ThemePart{
			"env":   {"#6FC30E"},
			"clo":   {"#b4e1fa", "#5b5d6e", "#515262", "#a0d2f0", "#a0d2f0"},
			"head":  {"#fae3b9"},
			"mouth": {"#fff", "#000"},
			"eyes":  {"#000"},
			"top":   {"#8eff45", "#8eff45", "none", "none"},
		},
		ThemePart{
			"env":   {"#00a58c"},
			"clo":   {"#000", "#5b00", "#5100", "#a000", "#a000"},
			"head":  {"#FAD2B9"},
			"mouth": {"#fff", "#000"},
			"eyes":  {"#000"},
			"top":   {"#FFC600", "none", "#FFC600", "none"},
		},
		ThemePart{
			"env":   {"#ff501f"},
			"clo":   {"#000", "#ff0000", "#ff0000", "#7d7d7d", "#7d7d7d"},
			"head":  {"#fff3dc"},
			"mouth": {"#d2001b", "none"},
			"eyes":  {"#000"},
			"top":   {"#D2001B", "none", "none", "#D2001B"},
		},
	},

	// Country主题
	style.CountryStyle: {
		ThemePart{
			"env":   {"#fc0"},
			"clo":   {"#901e0e", "#ffbe1e", "#ffbe1e", "#c55f54"},
			"head":  {"#f8d9ad"},
			"mouth": {"#000", "none", "#000", "none"},
			"eyes":  {"#000"},
			"top":   {"#583D00", "#AF892E", "#462D00", "#a0a0a0"},
		},
		ThemePart{
			"env":   {"#386465"},
			"clo":   {"#fff", "#333", "#333", "#333"},
			"head":  {"#FFD79D"},
			"mouth": {"#000", "#000", "#000", "#000"},
			"eyes":  {"#000"},
			"top":   {"#27363C", "#5DCAD4", "#314652", "#333"},
		},
		ThemePart{
			"env":   {"#DFFF00"},
			"clo":   {"#304267", "#aab0b1", "#aab0b1", "#aab0b1"},
			"head":  {"#e6b876"},
			"mouth": {"#50230a", "#50230a", "#50230a", "#50230a"},
			"eyes":  {"#000"},
			"top":   {"#333", "#afafaf", "#222", "#6d3a1d"},
		},
	},

	// Geeknot主题
	style.GeeknotStyle: {
		ThemePart{
			"env":   {"#a09300"},
			"clo":   {"#c7d4e2", "#435363", "#435363", "#141720", "#141720", "#e7ecf2", "#e7ecf2"},
			"head":  {"#f5d4a6"},
			"mouth": {"#000", "#cf9f76"},
			"eyes":  {"#000", "#000", "#000", "#000", "#000", "#000", "#fff", "#fff", "#fff", "#fff", "#000", "#000"},
			"top":   {"none", "#fdff00"},
		},
		ThemePart{
			"env":   {"#b3003e"},
			"clo":   {"#000", "#435363", "#435363", "#000", "none", "#e7ecf2", "#e7ecf2"},
			"head":  {"#f5d4a6"},
			"mouth": {"#000", "#af9f94"},
			"eyes":  {"#9ff3ffdb", "#000", "#9ff3ffdb", "#000", "#2f508a", "#000", "#000", "#000", "none", "none", "none", "none"},
			"top":   {"#ff9a00", "#ff9a00"},
		},
		ThemePart{
			"env":   {"#884f00"},
			"clo":   {"#ff0000", "#fff", "#fff", "#141720", "#141720", "#e7ecf2", "#e7ecf2"},
			"head":  {"#c57b14"},
			"mouth": {"#000", "#cf9f76"},
			"eyes":  {"none", "#000", "none", "#000", "#5a0000", "#000", "#000", "#000", "none", "none", "none", "none"},
			"top":   {"#efefef", "none"},
		},
	},

	// Asian主题
	style.AsianStyle: {
		ThemePart{
			"env":   {"#8acf00"},
			"clo":   {"#ee2829", "#ff0"},
			"head":  {"#ffce73"},
			"mouth": {"#fff", "#000"},
			"eyes":  {"#000"},
			"top":   {"#000", "#000", "none", "#000", "#ff4e4e", "#000"},
		},
		ThemePart{
			"env":   {"#00d2a3"},
			"clo":   {"#0D0046", "#ffce73"},
			"head":  {"#ffce73"},
			"mouth": {"#000", "none"},
			"eyes":  {"#000"},
			"top":   {"#000", "#000", "#000", "none", "#ffb358", "#000", "none", "none"},
		},
		ThemePart{
			"env":   {"#ff184e"},
			"clo":   {"#000", "none"},
			"head":  {"#ffce73"},
			"mouth": {"#ff0000", "none"},
			"eyes":  {"#000"},
			"top":   {"none", "none", "none", "none", "none", "#ffc107", "none", "none"},
		},
	},

	// Punk主题
	style.PunkStyle: {
		ThemePart{
			"env":   {"#00deae"},
			"clo":   {"#ff0000"},
			"head":  {"#ffce94"},
			"mouth": {"#f73b6c", "#000"},
			"eyes":  {"#e91e63", "#000", "#e91e63", "#000", "#000", "#000"},
			"top":   {"#dd104f", "#dd104f", "#f73b6c", "#dd104f"},
		},
		ThemePart{
			"env":   {"#181284"},
			"clo":   {"#491f49", "#ff9809", "#491f49"},
			"head":  {"#f6ba97"},
			"mouth": {"#ff9809", "#000"},
			"eyes":  {"#c4ffe4", "#000", "#c4ffe4", "#000", "#000", "#000"},
			"top":   {"none", "none", "#d6f740", "#516303"},
		},
		ThemePart{
			"env":   {"#bcf700"},
			"clo":   {"#ff14e4", "#000", "#14fffd"},
			"head":  {"#7b401e"},
			"mouth": {"#666", "#000"},
			"eyes":  {"#00b5b4", "#000", "#00b5b4", "#000", "#000", "#000"},
			"top":   {"#14fffd", "#14fffd", "#14fffd", "#0d3a62"},
		},
	},

	// Afrohair主题
	style.AfrohairStyle: {
		ThemePart{
			"env":   {"#0df"},
			"clo":   {"#571e57", "#ff0"},
			"head":  {"#f2c280"},
			"mouth": {"#ff0000"},
			"eyes":  {"#795548", "#000"},
			"top":   {"#de3b00", "none"},
		},
		ThemePart{
			"env":   {"#B400C2"},
			"clo":   {"#0D204A", "#00ffdf"},
			"head":  {"#ca8628"},
			"mouth": {"#1a1a1a"},
			"eyes":  {"#cbbdaf", "#000"},
			"top":   {"#000", "#000"},
		},
		ThemePart{
			"env":   {"#ffe926"},
			"clo":   {"#00d6af", "#000"},
			"head":  {"#8c5100"},
			"mouth": {"#7d0000"},
			"eyes":  {"none", "#000"},
			"top":   {"#f7f7f7", "none"},
		},
	},

	// Normie female主题
	style.NormieFemaleStyle: {
		ThemePart{
			"env":   {"#4aff0c"},
			"clo":   {"#101010", "#fff", "#fff"},
			"head":  {"#dbbc7f"},
			"mouth": {"#000"},
			"eyes":  {"#000", "none", "none"},
			"top":   {"#531148", "#531148", "#531148", "none"},
		},
		ThemePart{
			"env":   {"#FFC107"},
			"clo":   {"#033c58", "#fff", "#fff"},
			"head":  {"#dbc97f"},
			"mouth": {"#000"},
			"eyes":  {"none", "#fff", "#000"},
			"top":   {"#FFEB3B", "#FFEB3B", "none", "#FFEB3B"},
		},
		ThemePart{
			"env":   {"#FF9800"},
			"clo":   {"#b40000", "#fff", "#fff"},
			"head":  {"#E2AF6B"},
			"mouth": {"#000"},
			"eyes":  {"none", "#fff", "#000"},
			"top":   {"#ec0000", "#ec0000", "none", "none"},
		},
	},

	// Older主题
	style.OlderStyle: {
		ThemePart{
			"env":   {"#104c8c"},
			"clo":   {"#354B65", "#3D8EBB", "#89D0DA", "#00FFFD"},
			"head":  {"#cc9a5c"},
			"mouth": {"#222", "#fff"},
			"eyes":  {"#000", "#000"},
			"top":   {"#fff", "#fff", "none"},
		},
		ThemePart{
			"env":   {"#0DC15C"},
			"clo":   {"#212121", "#fff", "#212121", "#fff"},
			"head":  {"#dca45f"},
			"mouth": {"#111", "#633b1d"},
			"eyes":  {"#000", "#000"},
			"top":   {"none", "#792B74", "#792B74"},
		},
		ThemePart{
			"env":   {"#ffe500"},
			"clo":   {"#1e5e80", "#fff", "#1e5e80", "#fff"},
			"head":  {"#e8bc86"},
			"mouth": {"#111", "none"},
			"eyes":  {"#000", "#000"},
			"top":   {"none", "none", "#633b1d"},
		},
	},

	// Firehair主题
	style.FirehairStyle: {
		ThemePart{
			"env":   {"#4a3f73"},
			"clo":   {"#e6e9ee", "#f1543f", "#ff7058", "#fff", "#fff"},
			"head":  {"#b27e5b"},
			"mouth": {"#191919", "#191919"},
			"eyes":  {"#000", "#000", "#57FFFD"},
			"top":   {"#ffc", "#ffc", "#ffc"},
		},
		ThemePart{
			"env":   {"#00a08d"},
			"clo":   {"#FFBA32", "#484848", "#4e4e4e", "#fff", "#fff"},
			"head":  {"#ab5f2c"},
			"mouth": {"#191919", "#191919"},
			"eyes":  {"#000", "#ff23fa63", "#000"},
			"top":   {"#ff90f4", "#ff90f4", "#ff90f4"},
		},
		ThemePart{
			"env":   {"#22535d"},
			"clo":   {"#000", "#ff2500", "#ff2500", "#fff", "#fff"},
			"head":  {"#a76c44"},
			"mouth": {"#191919", "#191919"},
			"eyes":  {"#000", "none", "#000"},
			"top":   {"none", "#00efff", "none"},
		},
	},

	// Blond主题
	style.BlondStyle: {
		ThemePart{
			"env":   {"#2668DC"},
			"clo":   {"#2385c6", "#b8d0e0", "#b8d0e0"},
			"head":  {"#ad8a60"},
			"mouth": {"#000", "#4d4d4d"},
			"eyes":  {"#7fb5a2", "#d1eddf", "#301e19"},
			"top":   {"#fff510", "#fff510"},
		},
		ThemePart{
			"env":   {"#643869"},
			"clo":   {"#D67D1B", "#b8d0e0", "#b8d0e0"},
			"head":  {"#CC985A", "none0000"},
			"mouth": {"#000", "#ececec"},
			"eyes":  {"#1f2644", "#9b97ce", "#301e19"},
			"top":   {"#00eaff", "none"},
		},
		ThemePart{
			"env":   {"#F599FF"},
			"clo":   {"#2823C6", "#b8d0e0", "#b8d0e0"},
			"head":  {"#C7873A"},
			"mouth": {"#000", "#4d4d4d"},
			"eyes":  {"#581b1b", "#FF8B8B", "#000"},
			"top":   {"none", "#9c0092"},
		},
	},

	// Ateam主题
	style.AteamStyle: {
		ThemePart{
			"env":   {"#d10084"},
			"clo":   {"#efedee", "#00a1e0", "#00a1e0", "#efedee", "#ffce1c"},
			"head":  {"#b35f49"},
			"mouth": {"#3a484a", "#000"},
			"eyes":  {"#000"},
			"top":   {"#000", "none", "#000", "none"},
		},
		ThemePart{
			"env":   {"#E6C117"},
			"clo":   {"#efedee", "#ec0033", "#ec0033", "#efedee", "#f2ff05"},
			"head":  {"#ffc016"},
			"mouth": {"#4a3737", "#000"},
			"eyes":  {"#000"},
			"top":   {"#ffe900", "#ffe900", "none", "#ffe900"},
		},
		ThemePart{
			"env":   {"#1d8c00"},
			"clo":   {"#e000cb", "#fff", "#fff", "#e000cb", "#ffce1c"},
			"head":  {"#b96438"},
			"mouth": {"#000", "#000"},
			"eyes":  {"#000"},
			"top":   {"#53ffff", "#53ffff", "none", "none"},
		},
	},

	// Rasta主题
	style.RastaStyle: {
		ThemePart{
			"env":   {"#fc0065"},
			"clo":   {"#708913", "#fdea14", "#708913", "#fdea14", "#708913"},
			"head":  {"#DEA561"},
			"mouth": {"#444", "#000"},
			"eyes":  {"#000"},
			"top":   {"#32393f", "#32393f", "#32393f", "#32393f", "#32393f", "#32393f", "#32393f", "#32393f", "#32393f", "#32393f", "#32393f", "#32393f", "#32393f", "#32393f", "#32393f", "#32393f", "#32393f"},
		},
		ThemePart{
			"env":   {"#81f72e"},
			"clo":   {"#ff0000", "#ffc107", "#ff0000", "#ffc107", "#ff0000"},
			"head":  {"#ef9831"},
			"mouth": {"#6b0000", "#000"},
			"eyes":  {"#000"},
			"top":   {"#FFFAAD", "#FFFAAD", "#FFFAAD", "#FFFAAD", "#FFFAAD", "#FFFAAD", "#FFFAAD", "#FFFAAD", "#FFFAAD", "#FFFAAD", "#FFFAAD", "#FFFAAD", "#FFFAAD", "none", "none", "none", "none"},
		},
		ThemePart{
			"env":   {"#00D872"},
			"clo":   {"#590D00", "#FD1336", "#590D00", "#FD1336", "#590D00"},
			"head":  {"#c36c00"},
			"mouth": {"#56442b", "#000"},
			"eyes":  {"#000"},
			"top":   {"#004E4C", "#004E4C", "#004E4C", "#004E4C", "#004E4C", "#004E4C", "#004E4C", "#004E4C", "#004E4C", "none", "none", "none", "none", "none", "none", "none", "none"},
		},
	},

	//Meta主题
	style.MetaStyle: {
		ThemePart{
			"env":   {"#111"},
			"clo":   {"#000", "#00FFFF"},
			"head":  {"#755227"},
			"mouth": {"#fff", "#000"},
			"eyes":  {"black", "#008a", "aqua"},
			"top":   {"#fff", "#fff", "#fff", "#fff", "#fff"},
		},
		ThemePart{
			"env":   {"#00D0D4"},
			"clo":   {"#000", "#fff"},
			"head":  {"#755227"},
			"mouth": {"#fff", "#000"},
			"eyes":  {"black", "#1df7ffa3", "#fcff2c"},
			"top":   {"#fff539", "none", "#fff539", "none", "#fff539"},
		},
		ThemePart{
			"env":   {"#DC75FF"},
			"clo":   {"#000", "#FFBDEC"},
			"head":  {"#997549"},
			"mouth": {"#fff", "#000"},
			"eyes":  {"black", "black", "aqua"},
			"top":   {"#00fffd", "none", "none", "none", "none"},
		},
	},
	// Square主题
	style.SquareStyle: {
		ThemePart{
			"env":   {"#111"},
			"clo":   {"#000", "#00FFFF"},
			"head":  {"#755227"},
			"mouth": {"#fff", "#000"},
			"eyes":  {"black", "#008a", "aqua"},
			"top":   {"#fff", "#fff", "#fff", "#fff", "#fff"},
		},
		ThemePart{
			"env":   {"#00D0D4"},
			"clo":   {"#000", "#fff"},
			"head":  {"#755227"},
			"mouth": {"#fff", "#000"},
			"eyes":  {"black", "#1df7ffa3", "#fcff2c"},
			"top":   {"#fff539", "none", "#fff539", "none", "#fff539"},
		},
		ThemePart{
			"env":   {"#DC75FF"},
			"clo":   {"#000", "#FFBDEC"},
			"head":  {"#997549"},
			"mouth": {"#fff", "#000"},
			"eyes":  {"black", "black", "aqua"},
			"top":   {"#00fffd", "none", "none", "none", "none"},
		},
	},
}

// initThemes 初始化主题数据
func (m *Manager) initThemes() {

	for _, theme := range []style.StyleType{
		style.RoboStyle,
		style.GirlStyle,
		style.BlondeStyle,
		style.GuyStyle,
		style.CountryStyle,
		style.GeeknotStyle,
		style.AsianStyle,
		style.PunkStyle,
		style.AfrohairStyle,
		style.NormieFemaleStyle,
		style.OlderStyle,
		style.FirehairStyle,
		style.BlondStyle,
		style.AteamStyle,
		style.RastaStyle,
		style.MetaStyle,
		style.SquareStyle,
	} {
		if themeSet, exists := defaultThemeSet[theme]; exists {
			m.AddTheme(themeSet)
		}
	}
}
