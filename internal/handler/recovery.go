package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "net/http"
    "runtime/debug"
    "skeleton/pkg/exception"
    "skeleton/pkg/validatorx"
    "strings"
)

func Recovery(c *gin.Context) {
    defer func() {
        if r := recover(); r != nil {
            res := gin.H{}

            stack := string(debug.Stack())
            if gin.IsDebugging() {
                res["err_stack"] = stack
            }

            switch v := r.(type) {
            case exception.Exception:
                res["err_code"] = v.Code()
                res["err_message"] = v.Message()
                c.AbortWithStatusJSON(v.Status(), res)
            case validator.ValidationErrors:
                res["err_message"] = http.StatusBadRequest
                res["err_message"] = removeTopStruct(v.Translate(validatorx.Translator))
                c.AbortWithStatusJSON(http.StatusBadRequest, res)
            default:
                res["err_code"] = http.StatusInternalServerError
                res["err_message"] = http.StatusText(http.StatusInternalServerError)
                c.AbortWithStatusJSON(http.StatusInternalServerError, res)
            }
        }
    }()

    c.Next()
}

func removeTopStruct(fields map[string]string) map[string]string {
    rsp := map[string]string{}
    for field, err := range fields {
        rsp[field[strings.Index(field, ".")+1:]] = err
    }
    return rsp
}
