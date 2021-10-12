package basicMiddleware

import (
	"github.com/gin-gonic/gin"
	"skeleton/request"
)

func Request(c *gin.Context) {
	_ = c.ShouldBindHeader(&request.Header)
	request.All(c)
	request.Headers(c)
	c.Next()
}
