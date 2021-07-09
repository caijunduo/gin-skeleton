package middleware

import (
    "bytes"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "go.uber.org/zap"
    "io/ioutil"
    "net/url"
    "skeleton/internal/errno"
    "skeleton/pkg/signaturex"
)

func Signature(conn string) gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method
        data := make(map[string]interface{})

        switch method {
        case "GET":
            if c.Request.URL.Query() != nil {
                for k, v := range c.Request.URL.Query() {
                    data[k] = v[0]
                }
            }
        case "POST":
            fallthrough
        case "PUT":
            fallthrough
        case "PATCH":
            fallthrough
        case "DELETE":
            if c.ContentType() == "application/json" {
                b, _ := c.GetRawData()
                var body map[string]interface{}
                _ = json.Unmarshal(b, &body)
                c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
                for k, v := range body {
                    data[k] = v
                }
                break
            }

            if c.Request.ParseForm() == nil {
                _ = c.Request.ParseMultipartForm(32 << 20)
            }

            formData := make(url.Values)
            switch c.ContentType() {
            case "multipart/form-data":
                formData = c.Request.PostForm
            case "application/x-www-form-urlencoded":
                formData = c.Request.Form
            }
            if formData != nil {
                for k, v := range formData {
                    data[k] = v[0]
                }
            }
        }

        switch viper.GetString("signature." + conn + ".method") {
        case "md5":
            m := signaturex.NewMd5(conn).SetData(data)
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
            r := signaturex.NewRSA(conn).SetData(data)
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
