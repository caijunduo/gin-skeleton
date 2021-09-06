package validator

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Validator  *validator.Validate
	Translator ut.Translator
	ok         bool
)

func New() error {
	locale := "zh"
	Validator, ok = binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return errors.New("validator engine error")
	}

	uni := ut.New(zh.New(), zh.New())
	Translator, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s)", locale)
	}

	_ = zhTranslations.RegisterDefaultTranslations(Validator, Translator)
	RegisterValidation("mobile", "非法手机号码", Mobile)

	return nil
}

func RegisterValidation(tag string, message string, validate func(fl validator.FieldLevel) bool) {
	_ = Validator.RegisterValidation(tag, validate)

	_ = Validator.RegisterTranslation(tag, Translator, func(ut ut.Translator) error {
		return ut.Add(tag, "{0}"+message, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
}
