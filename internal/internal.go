package internal

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"skeleton/pkg/authorization"
	"skeleton/request"
	"skeleton/response"
	"time"
)

type Root struct {
	exampleInternal Example
}

func (r Root) RouteGroup(rg *gin.RouterGroup) {
	rg.GET("", func(c *gin.Context) {
		c.JSON(response.OK.Slice())
	})
	r.exampleInternal.RouteGroup(rg.Group("example"))
}

func (r Root) RouteEngine(e *gin.Engine) {
	e.NoRoute(r.noRoute)
	e.Use(r.request)
	e.Use(r.response)
	e.Use(r.logger)
	if gin.IsDebugging() {
		e.Use(cors.Default())
	}
}

func (Root) noRoute(c *gin.Context) {
	c.AbortWithStatusJSON(response.NotFound.Slice())
	return
}

func (Root) request(c *gin.Context) {
	_ = c.ShouldBindHeader(&request.Header)
	request.All(c)
	c.Next()
}

func (Root) response(c *gin.Context) {
	response.Writer.ResponseWriter = c.Writer
	c.Writer = response.Writer
	c.Next()
}

func (Root) logger(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	body := request.All(c)
	headers := c.Request.Header
	c.Next()
	latency := time.Since(start)
	zap.L().Info("Request",
		zap.Int("status", c.Writer.Status()),
		zap.String("method", c.Request.Method),
		zap.String("path", path),
		zap.Any("body", body),
		zap.Any("headers", headers),
		zap.String("ip", c.ClientIP()),
		zap.String("user-agent", c.Request.UserAgent()),
		zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		zap.Duration("latency", latency),
	)
}

func (Root) authorization(opt authorization.JwtOption, handler func(c *gin.Context, uuid string)) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		j := authorization.JwtClaims{Opt: opt}

		if valid, data, err = j.Parse(token); err != nil || valid {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				if token, err = j.SetData(map[string]interface{}{
					"uuid": cast.ToString(data["uuid"]),
				}).Generate(); err != nil {
					c.AbortWithStatusJSON(response.UnAuthorization.Slice())
					return
				}
				c.Header("Authorization", token)
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

		handler(c, uuid)

		c.Next()
	}
}
