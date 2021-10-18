package basicMiddleware

import (
	"skeleton/context"
	operateHelper "skeleton/helper/operate"
	"skeleton/logger"
	"time"
)

func Logger(c *context.Context) {
	start := time.Now()
	c.Next()
	logger.HTTP.WithField("status", c.Writer.Status()).
		WithField("method", c.Request.Method).
		WithField("path", c.Request.URL.Path).
		WithField("body", operateHelper.JSONEncodeToString(c.All())).
		WithField("headers", operateHelper.JSONEncodeToString(c.Headers())).
		WithField("ip", c.ClientIP()).
		WithField("latency", time.Since(start)).
		Info()
}
