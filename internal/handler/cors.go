package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "net/http"
)

func Cors(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", viper.GetString("cors.allowOrigin"))
    c.Header("Access-Control-Allow-Methods", viper.GetString("cors.allowMethods"))
    c.Header("Access-Control-Allow-Headers", viper.GetString("cors.allowHeaders"))
    c.Header("Access-Control-Expose-Headers", viper.GetString("cors.exposeHeaders"))
    c.Header("Access-Control-Allow-Credentials", viper.GetString("cors.allowCredentials"))

    if c.Request.Method == "OPTIONS" {
        c.AbortWithStatus(http.StatusNoContent)
    }

    c.Next()
}
