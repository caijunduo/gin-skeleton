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
    postForm := c.Request.PostForm
    body := c.Request.Body
    headers := c.Request.Header

    c.Next()

    latency := time.Since(start)
    
    zap.L().Info("[Request]",
        zap.Int("status", c.Writer.Status()),
        zap.String("method", c.Request.Method),
        zap.String("path", path),
        zap.String("query", query),
        zap.Any("post-form", postForm),
        zap.Any("body", body),
        zap.Any("headers", headers),
        zap.String("ip", c.ClientIP()),
        zap.String("user-agent", c.Request.UserAgent()),
        zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
        zap.Duration("latency", latency),
    )
}
