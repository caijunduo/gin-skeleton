package handler

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

            stack := string(debug.Stack())
            if gin.IsDebugging() {
                res["err_stack"] = stack
            }

            switch v := r.(type) {
            case exception.Exception:
                res["err_code"] = v.Code()
                res["err_message"] = v.Message()
                c.AbortWithStatusJSON(v.Status(), res)
            default:
                res["err_code"] = http.StatusInternalServerError
                res["err_message"] = http.StatusText(http.StatusInternalServerError)
                c.AbortWithStatusJSON(http.StatusInternalServerError, res)
            }
        }
    }()

    c.Next()
}
