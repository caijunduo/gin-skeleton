package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"skeleton/errno"
	validatorPkg "skeleton/pkg/validator"
)

func Validator(c *gin.Context) {
	c.Next()
	for _, v := range c.Errors {
		switch e := v.Err.(type) {
		case validator.ValidationErrors:
			err := errno.InvalidParameters
			c.AbortWithStatusJSON(err.Status(), gin.H{
				"err_code":    err.Code(),
				"err_message": err.Message(),
				"err_details": e.Translate(validatorPkg.Translator),
			})
		}
	}
}
