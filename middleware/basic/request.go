package basicMiddleware

import (
	"github.com/gin-gonic/gin"
	"skeleton/request"
)

func Request() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = c.ShouldBindHeader(&request.Header)
		request.All(c)
		request.Headers(c)
		c.Next()
	}
}
