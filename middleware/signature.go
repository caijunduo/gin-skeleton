package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"skeleton/errno"
	"skeleton/pkg"
	"skeleton/pkg/signature"
)

func Signature(method string, opt signature.Option) gin.HandlerFunc {
	return func(c *gin.Context) {
		all := pkg.Context.All()

		switch method {
		case "md5":
			r := pkg.SignatureMD5(opt).SetData(all)
			if err := r.Verify(); err != nil {
				if gin.IsDebugging() {
					zap.L().Debug("[Signature]",
						zap.String("before_encrypt", r.GetBeforeEncrypt()),
						zap.String("before_signature", r.GetBeforeSignature()),
						zap.String("signature", r.GetSignature()),
						zap.Error(err),
					)
				}
				c.AbortWithStatusJSON(errno.InvalidSignature.ToSlice())
				return
			}
		case "rsa":
			r := pkg.SignatureRSA(opt).SetData(all)
			if err := r.Verify(); err != nil {
				if gin.IsDebugging() {
					zap.L().Debug("[Signature]",
						zap.String("before_encrypt", r.GetBeforeEncrypt()),
						zap.String("before_signature", r.GetBeforeSignature()),
						zap.String("signature", r.GetSignature()),
						zap.Error(err),
					)
				}
				c.AbortWithStatusJSON(errno.InvalidSignature.ToSlice())
				return
			}
		}

		c.Next()
	}
}
