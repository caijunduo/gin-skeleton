package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
)

func Before(c *gin.Context) {
    log.Println("before")
    c.Next()
}
