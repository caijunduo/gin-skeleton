package structure

import "strings"

type String string

func (s String) HidePartMobile() string {
    return string(s)[0:3] + "****" + string(s)[7:]
}

func (s String) IsHttpOrHttps() bool {
    return strings.HasPrefix(string(s), "http") || strings.HasPrefix(string(s), "https")
}
