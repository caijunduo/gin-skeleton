package accountMiddleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"skeleton/config"
	authorizationPkg "skeleton/pkg/authorization"
	"skeleton/request"
	"skeleton/response"
	"skeleton/text"
)

func Authorization(c *gin.Context) {
	var (
		uuid  string
		token string
		err   error
		data  map[string]interface{}
		valid bool
	)

	if token = request.Header.Authorization; token == "" {
		c.AbortWithStatusJSON(response.InvalidAuthorization.Slice())
		return
	}

	j := authorizationPkg.JwtClaims{Opt: authorizationPkg.JwtOption{
		Key:           config.Authorization.Key,
		Issuer:        config.Authorization.Issuer,
		Subject:       config.Authorization.Subject,
		ExpireMinutes: config.Authorization.ExpireMinutes,
	}}

	if valid, data, err = j.Parse(token); err != nil || valid {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			if token, err = j.SetData(map[string]interface{}{
				"uuid": cast.ToString(data["uuid"]),
			}).Generate(); err != nil {
				c.AbortWithStatusJSON(response.UnAuthorization.Slice())
				return
			}
			c.Header(text.HTTPAuthorization.ToValue(), token)
		default:
			c.AbortWithStatusJSON(response.InvalidAuthorization.Slice())
			return
		}
	} else {
		uuid = cast.ToString(data["uuid"])
	}

	if uuid == "" {
		c.AbortWithStatusJSON(response.InvalidAuthorization.Slice())
		return
	}

	c.Set("uuid", uuid)

	c.Next()
}
