package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "skeleton/configs"
)

func Cors(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", configs.Cors.AllowOrigin)
    c.Header("Access-Control-Allow-Methods", configs.Cors.AllowMethods)
    c.Header("Access-Control-Allow-Headers", configs.Cors.AllowHeaders)
    c.Header("Access-Control-Expose-Headers", configs.Cors.AllowExposeHeaders)
    c.Header("Access-Control-Allow-Credentials", configs.Cors.AllowCredentials)

    if c.Request.Method == "OPTIONS" {
        c.AbortWithStatus(http.StatusNoContent)
    }

    c.Next()
}
