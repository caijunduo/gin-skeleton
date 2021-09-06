package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
	"skeleton/config"
)

func Cors(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", config.Cors.AllowOrigin)
    c.Header("Access-Control-Allow-Methods", config.Cors.AllowMethods)
    c.Header("Access-Control-Allow-Headers", config.Cors.AllowHeaders)
    c.Header("Access-Control-Expose-Headers", config.Cors.AllowExposeHeaders)
    c.Header("Access-Control-Allow-Credentials", config.Cors.AllowCredentials)

    if c.Request.Method == "OPTIONS" {
        c.AbortWithStatus(http.StatusNoContent)
    }

    c.Next()
}
