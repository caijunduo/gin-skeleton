package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"skeleton/errno"
	jwtPkg "skeleton/pkg/jwt"
)

func Authorization(opt jwtPkg.Option, handler func(c *gin.Context, uuid string)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uuid          string
			authorization string
			err           error
			data          map[string]interface{}
			valid         bool
		)

		if authorization = c.GetHeader("Authorization"); authorization == "" {
			panic(errno.InvalidAuthorization)
		}

		j := jwtPkg.New(opt)

		if valid, data, err = j.Parse(authorization); err != nil || valid {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				if authorization, err = j.SetData(map[string]interface{}{
					"uuid": cast.ToString(data["uuid"]),
				}).Generate(); err != nil {
					panic(errno.UnAuthorization)
				}
				c.Header("Authorization", authorization)
			default:
				panic(errno.InvalidAuthorization)
			}
		} else {
			uuid = cast.ToString(data["uuid"])
		}

		if uuid == "" {
			panic(errno.InvalidAuthorization)
		}

		c.Set("uuid", uuid)

		handler(c, uuid)

		c.Next()
	}
}
