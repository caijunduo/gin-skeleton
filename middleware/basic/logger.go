package basicMiddleware

import (
	"skeleton/context"
	"skeleton/logger"
	"time"
)

func Logger(c *context.Context) {
	start := time.Now()
	c.Next()
	logger.HTTP.Info().
		Interface("body", c.All()).
		Interface("headers", c.Headers()).
		Str("ip", c.ClientIP()).
		Dur("latency", time.Since(start)).
		Str("path", c.Request.URL.Path).
		Str("method", c.Request.Method).
		Int("status", c.Writer.Status()).
		Msg("")
}
