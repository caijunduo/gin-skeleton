package internal

import (
    "github.com/gin-gonic/gin"
    "skeleton/internal/handler"
    "skeleton/internal/middleware"
)

func Routes() *gin.Engine {
    r := gin.New()

    r.Use(middleware.Recovery)
    r.Use(middleware.RequestId)
    r.Use(middleware.Cors)
    r.Use(middleware.Logger)
    // r.Use(handler.Authorization("account"))
    // r.Use(handler.Signature("api"))

    r.GET("/", middleware.Before, middleware.After, handler.Example)

    return r
}
