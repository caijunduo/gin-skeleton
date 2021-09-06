package middleware

import (
    "github.com/gin-gonic/gin"
	ginPkg "skeleton/pkg/gin"
)

func Response(c *gin.Context) {
	ginPkg.NewResponseWriter(c.Writer)
    c.Writer = ginPkg.Writer
    c.Next()
}
