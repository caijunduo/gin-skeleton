package basicMiddleware

import (
	"skeleton/context"
	"skeleton/logger"
	"skeleton/response"
)

func Recovery(c *context.Context) {
	defer func() {
		if rc := recover(); rc != nil {
			logger.HTTP.Panic().Stack().Err(rc.(error))
			switch v := rc.(type) {
			default:
				c.AbortWithStatusJSON(response.InternalServerError.SetError(v.(error)).Slice())
			}
		}
	}()

	c.Next()
}
