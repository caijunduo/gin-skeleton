package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "go.uber.org/zap"
    "skeleton/internal/errno"
    "skeleton/pkg/ginx"
    "skeleton/pkg/signaturex"
)

func Signature(conn string) gin.HandlerFunc {
    return func(c *gin.Context) {
        all := ginx.Ctx.All()

        switch viper.GetString("signature." + conn + ".method") {
        case "md5":
            m := signaturex.NewMd5(conn).SetData(all)
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
        case "rsa":
            r := signaturex.NewRSA(conn).SetData(all)
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
