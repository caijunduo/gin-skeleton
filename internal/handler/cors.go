package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
)

func Cors(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", os.Getenv("CORS_ALLOW_ORIGIN"))
    c.Header("Access-Control-Allow-Methods", os.Getenv("CORS_ALLOW_METHODS"))
    c.Header("Access-Control-Allow-Headers", os.Getenv("CORS_ALLOW_HEADERS"))
    c.Header("Access-Control-Expose-Headers", os.Getenv("CORS_EXPOSE_HEADERS"))
    c.Header("Access-Control-Allow-Credentials", os.Getenv("CORS_ALLOW_CREDENTIALS"))

    if c.Request.Method == "OPTIONS" {
        c.AbortWithStatus(http.StatusNoContent)
    }

    c.Next()
}
