package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"skeleton/errno"
	"skeleton/pkg"
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
			c.AbortWithStatusJSON(errno.InvalidAuthorization.ToSlice())
			return
		}

		j := pkg.Jwt(opt)

		if valid, data, err = j.Parse(authorization); err != nil || valid {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				if authorization, err = j.SetData(map[string]interface{}{
					"uuid": cast.ToString(data["uuid"]),
				}).Generate(); err != nil {
					c.AbortWithStatusJSON(errno.UnAuthorization.ToSlice())
					return
				}
				c.Header("Authorization", authorization)
			default:
				c.AbortWithStatusJSON(errno.InvalidAuthorization.ToSlice())
				return
			}
		} else {
			uuid = cast.ToString(data["uuid"])
		}

		if uuid == "" {
			c.AbortWithStatusJSON(errno.InvalidAuthorization.ToSlice())
			return
		}

		c.Set("uuid", uuid)

		handler(c, uuid)

		c.Next()
	}
}
