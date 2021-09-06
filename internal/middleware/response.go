package middleware

import (
    "github.com/gin-gonic/gin"
    "skeleton/pkg/ginx"
)

func Response(c *gin.Context) {
    ginx.NewResponseWriter(c.Writer)
    c.Writer = ginx.Writer
    c.Next()
}
