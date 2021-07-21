package middleware

import (
    "github.com/gin-gonic/gin"
    "skeleton/pkg/ginx"
)

func Context(c *gin.Context) {
    ginx.New(c)
    c.Next()
}
