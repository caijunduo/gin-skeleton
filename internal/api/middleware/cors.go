package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "skeleton/configs"
)

func Cors(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", configs.Config.Cors.AllowOrigin)
    c.Header("Access-Control-Allow-Methods", configs.Config.Cors.AllowMethods)
    c.Header("Access-Control-Allow-Headers", configs.Config.Cors.AllowHeaders)
    c.Header("Access-Control-Expose-Headers", configs.Config.Cors.AllowExposeHeaders)
    c.Header("Access-Control-Allow-Credentials", configs.Config.Cors.AllowCredentials)

    if c.Request.Method == "OPTIONS" {
        c.AbortWithStatus(http.StatusNoContent)
    }

    c.Next()
}
