package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "go.uber.org/zap"
)

func RequestId(c *gin.Context) {
    requestId := c.Request.Header.Get("X-Request-Id")

    if requestId == "" {
        uuidStr, err := uuid.NewRandom()
        if err != nil {
            zap.L().Error("[RequestId] uuid generate failed")
            c.Abort()
            return
        }
        requestId = uuidStr.String()
    }

    c.Set("X-Request-Id", requestId)
    c.Header("X-Request-Id", requestId)

    c.Next()
}
