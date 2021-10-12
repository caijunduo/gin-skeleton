package validate

import (
	"skeleton/validation"
)

var (
	Mobile = validation.Match("^((13[0-9])|(14[5,6,7,8,9])|(15([0-3]|[5-9]))|(166)|(17[0-8])|(18[0-9])|(19([0-3]|[5-9])))\\d{8}$")
)
