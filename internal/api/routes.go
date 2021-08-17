package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "skeleton/internal/api/middleware"
)

func Routes() *gin.Engine {
    r := gin.New()

    r.Use(middleware.Context)
    r.Use(middleware.Recovery)
    r.Use(middleware.RequestId)
    r.Use(middleware.Cors)
    r.Use(middleware.Logger)

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, Skeleton")
    })
    // r.Use(middleware.Authorization(jwtx.Option{
    //     Key:           configs.Config.Cmd.Api.Jwt.Key,
    //     Issuer:        configs.Config.Cmd.Api.Jwt.Issuer,
    //     Subject:       configs.Config.Cmd.Api.Jwt.Subject,
    //     ExpireMinutes: configs.Config.Cmd.Api.Jwt.ExpireMinutes,
    // }, func(c *gin.Context, uuid string) {}))
    // r.Use(middleware.Signature(signaturex.Md5, signaturex.Option{
    //     AppKey:    configs.Config.Cmd.Api.Signature.AppKey,
    //     AppSecret: configs.Config.Cmd.Api.Signature.AppSecret,
    //     Expires:   configs.Config.Cmd.Api.Signature.Expires,
    // }))

    return r
}
