package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "runtime/debug"
    "skeleton/pkg/exception"
)

func Recovery(c *gin.Context) {
    defer func() {
        if r := recover(); r != nil {
            res := gin.H{}
            if gin.IsDebugging() {
                res["err_stack"] = string(debug.Stack())
            }
            switch v := r.(type) {
            case exception.Exception:
                res["err_code"] = v.Code()
                res["err_message"] = v.Message()
                if gin.IsDebugging() {
                    res["err_stack_message"] = v.Error()
                }
                c.AbortWithStatusJSON(v.Status(), res)
            default:
                res["err_code"] = http.StatusInternalServerError
                res["err_message"] = http.StatusText(http.StatusInternalServerError)
                if gin.IsDebugging() {
                    res["err_stack_message"] = v
                }
                c.AbortWithStatusJSON(http.StatusInternalServerError, res)
            }
        }
    }()

    c.Next()
}
