package request

//header 全局Header
type header struct {
	Authorization    string `header:"Authorization"`       // 身份凭证
	VersionCode      int64  `header:"X-Version-Code"`      // 应用版本
	VersionName      string `header:"X-Version-Name"`      // 应用版本名称
	ApiVersion       string `header:"X-API-Version"`       // API版本
	MacCode          string `header:"X-Mac-Code"`          // Mac地址
	OsVersion        string `header:"X-OS-Version"`        // 系统版本
	OsName           string `header:"X-OS-Name"`           // 系统名称
	ResolutionWidth  string `header:"X-Resolution-Width"`  // 分辨率-宽
	ResolutionHeight int64  `header:"X-Resolution-Height"` // 分辨率-高
	PlatformID       int64  `header:"X-Platform-ID"`       // 平台ID
	PlatformName     string `header:"X-Platform-Name"`     // 平台名称
	ChannelID        int64  `header:"X-Channel-ID"`        // 渠道ID
	ChannelName      string `header:"X-Channel-Name"`      // 渠道名称
}

var Header = header{}
