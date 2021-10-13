package basicMiddleware

import (
    "github.com/gin-gonic/gin"
    operateHelper "skeleton/helper/operate"
    "skeleton/logger"
    "skeleton/request"
    "time"
)

func Logger(c *gin.Context) {
    start := time.Now()
    c.Next()
    logger.HTTP.WithField("status", c.Writer.Status()).
        WithField("method", c.Request.Method).
        WithField("path", c.Request.URL.Path).
        WithField("body", operateHelper.JSONEncodeToString(request.All(c))).
        WithField("headers", operateHelper.JSONEncodeToString(request.Headers(c))).
        WithField("ip", c.ClientIP()).
        WithField("latency", time.Since(start)).
        Info()
}
