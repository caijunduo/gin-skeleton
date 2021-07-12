package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "net/http"
    "runtime/debug"
    "skeleton/internal/errno"
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
            case validator.ValidationErrors:
                e := errno.InvalidParameters
                res["err_code"] = e.Code()
                res["err_message"] = e.Message()
                if gin.IsDebugging() {
                    res["err_details"] = removeTopStruct(v.Translate(validatorx.Translator))
                }
                c.AbortWithStatusJSON(e.Status(), res)
            case exception.Exception:
                res["err_code"] = v.Code()
                res["err_message"] = v.Message()
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

func removeTopStruct(fields map[string]string) map[string]interface{} {
    lowerMap := map[string]string{}
    for field, err := range fields {
        fieldArr := strings.SplitN(field, ".", 2)
        lowerMap[fieldArr[1]] = err
    }
    res := addValueToMap(lowerMap)
    return res
}

func addValueToMap(fields map[string]string) map[string]interface{} {
    res := make(map[string]interface{})
    for field, err := range fields {
        fieldArr := strings.SplitN(field, ".", 2)
        if len(fieldArr) > 1 {
            NewFields := map[string]string{fieldArr[1]: err}
            returnMap := addValueToMap(NewFields)
            if res[fieldArr[0]] != nil {
                for k, v := range returnMap {
                    res[fieldArr[0]].(map[string]interface{})[k] = v
                }
            } else {
                res[fieldArr[0]] = returnMap
            }
            continue
        } else {
            res[field] = err
            continue
        }
    }
    return res
}
