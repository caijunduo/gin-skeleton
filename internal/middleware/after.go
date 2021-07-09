package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
)

func After(c *gin.Context)  {
    c.Next()
    log.Println("after")
}