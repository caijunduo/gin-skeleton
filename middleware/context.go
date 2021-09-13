package middleware

import (
	"github.com/gin-gonic/gin"
	"skeleton/pkg"
)

func Context(c *gin.Context) {
	pkg.Context.Gin = c
	c.Next()
}
