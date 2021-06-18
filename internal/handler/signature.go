package handler

import (
    "bytes"
    "encoding/json"
    "github.com/gin-gonic/gin"
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

        m := signaturex.NewMd5(conn).SetData(data)
        if err := m.Verify(); err != nil {
            if gin.IsDebugging() {
                zap.L().Debug("[Signature]", zap.String("md5", m.GetMd5()), zap.String("no_md5", m.GetNoMd5()), zap.String("no_encrypt", m.GetNoEncrypt()))
            }
            panic(errno.InvalidSignature.SetMessage(err.Error()))
        }

        c.Next()
    }
}
