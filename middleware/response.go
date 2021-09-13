package middleware

import (
	"github.com/gin-gonic/gin"
	"skeleton/pkg"
)

func Response(c *gin.Context) {
	pkg.Writer.ResponseWriter = c.Writer
	c.Writer = pkg.Writer
	c.Next()
}
