package middleware

import (
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "skeleton/internal/api/errno"
    "skeleton/pkg/ginx"
    "skeleton/pkg/signaturex"
)

func Signature(method string, opt signaturex.Option) gin.HandlerFunc {
    return func(c *gin.Context) {
        all := ginx.Ctx.All()

        switch method {
        case signaturex.Md5:
            m := signaturex.NewMd5(opt).SetData(all)
            if err := m.Verify(); err != nil {
                if gin.IsDebugging() {
                    zap.L().Debug("[Signature]",
                        zap.String("before_encrypt", m.GetBeforeEncrypt()),
                        zap.String("before_signature", m.GetBeforeSignature()),
                        zap.String("signature", m.GetSignature()),
                        zap.Error(err),
                    )
                }
                panic(errno.InvalidSignature)
            }
        case signaturex.RSA:
            r := signaturex.NewRSA(opt).SetData(all)
            if err := r.Verify(); err != nil {
                if gin.IsDebugging() {
                    zap.L().Debug("[Signature]",
                        zap.String("before_encrypt", r.GetBeforeEncrypt()),
                        zap.String("before_signature", r.GetBeforeSignature()),
                        zap.String("signature", r.GetSignature()),
                        zap.Error(err),
                    )
                }
                panic(errno.InvalidSignature)
            }
        }

        c.Next()
    }
}
