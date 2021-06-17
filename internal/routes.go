package internal

import (
    "github.com/gin-gonic/gin"
    "skeleton/internal/handler"
)

func Routes() *gin.Engine {
    r := gin.New()

    r.Use(handler.Recovery)
    r.Use(handler.RequestId)
    r.Use(handler.Cors)
    r.Use(handler.Logger)

    r.GET("/", handler.Example)

    return r
}
