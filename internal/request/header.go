package request

type Header struct {
    VersionCode int64  `header:"version_code"` // 应用版本
    VersionName string `header:"version_name"` // 应用版本名称
    MacCode     string `header:"mac_code"`     // Mac地址
    OsVersion   string `header:"os_version"`   // 系统版本
    OsName      string `header:"os_name"`      // 系统名称
    Resolution  string `header:"resolution"`   // 分辨率
    PlatformID  int64  `header:"platform_id"`  // 平台ID
    ChannelID   int64  `header:"channel_id"`   // 渠道ID
    ChannelName string `header:"channel_name"` // 渠道名称
}
