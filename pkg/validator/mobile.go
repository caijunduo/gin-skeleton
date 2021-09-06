package validator

import (
    "github.com/go-playground/validator/v10"
    "regexp"
)

// Mobile 校验手机号
func Mobile(fl validator.FieldLevel) (ok bool) {
    mobile := fl.Field().String()
    ok, _ = regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
    return
}
