package handler

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "skeleton/internal/errno"
    "skeleton/internal/pkg/jwtx"
)

func Authorization(c *gin.Context) {
    var (
        uuid          string
        authorization string
        err           error
        t             *jwt.Token
    )

    if authorization = c.GetHeader("Authorization"); authorization == "" {
        panic(errno.InvalidAuthorization)
    }

    if t, err = jwtx.Parse(authorization); err != nil {
        panic(errno.InvalidAuthorization)
    }

    if !t.Valid {
        switch err.(*jwt.ValidationError).Errors {
        case jwt.ValidationErrorExpired:
            uuid = t.Claims.(*jwtx.Claims).UUID
            if authorization, err = jwtx.New(uuid); err != nil {
                panic(errno.UnAuthorization)
            }
            c.Header("Authorization", authorization)
        default:
            panic(errno.InvalidAuthorization)
        }
    } else {
        uuid = t.Claims.(*jwtx.Claims).UUID
    }

    if uuid == "" {
        panic(errno.InvalidAuthorization)
    }

    c.Set("uuid", uuid)

    c.Next()
}
