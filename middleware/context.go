package middleware

import (
	"github.com/gin-gonic/gin"
	ginPkg "skeleton/pkg/gin"
)

func Context(c *gin.Context) {
	ginPkg.New(c)
	c.Next()
}
