package handler

import (
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "time"
)

func Logger(c *gin.Context) {
    start := time.Now()
    path := c.Request.URL.Path
    query := c.Request.URL.RawQuery
    c.Next()

    cost := time.Since(start)
    zap.L().Info(path,
        zap.Int("status", c.Writer.Status()),
        zap.String("method", c.Request.Method),
        zap.String("path", path),
        zap.String("query", query),
        zap.String("ip", c.ClientIP()),
        zap.String("user-agent", c.Request.UserAgent()),
        zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
        zap.Duration("cost", cost),
    )
}
