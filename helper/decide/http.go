package decide

import "strings"

func IsHttpOrHttps(url string) bool {
	return strings.HasPrefix(url, "http") || strings.HasPrefix(url, "https")
}
