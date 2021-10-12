package basicMiddleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"skeleton/request"
	"time"
)

func Logger(c *gin.Context) {
	start := time.Now()
	c.Next()
	zap.L().Info("REQUEST",
		zap.Int("status", c.Writer.Status()),
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.Any("body", request.All(c)),
		zap.Any("headers", request.Headers(c)),
		zap.String("ip", c.ClientIP()),
		zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		zap.Duration("latency", time.Since(start)),
	)
}
