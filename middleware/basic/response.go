package basicMiddleware

import (
	"github.com/gin-gonic/gin"
	"skeleton/response"
)

func Response(c *gin.Context) {
	response.Writer.ResponseWriter = c.Writer
	c.Writer = response.Writer
	c.Next()
}
