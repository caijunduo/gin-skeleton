package helper

import (
	"strings"
)

type String struct{}

func (String) HidePartMobile(mobile string) string {
	return mobile[0:3] + "****" + mobile[7:]
}

func (String) IsHttpOrHttps(str string) bool {
	return strings.HasPrefix(str, "http") || strings.HasPrefix(str, "https")
}
