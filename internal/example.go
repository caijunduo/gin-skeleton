package internal

import (
	"github.com/gin-gonic/gin"
	"skeleton/cache"
	"skeleton/request"
	"skeleton/response"
)

type Example struct {
	exampleCache cache.Example
}

func (e Example) RouteGroup(r *gin.RouterGroup) {
	r.GET("", e.example)
}

func (e Example) example(c *gin.Context) {
	var v request.Pagination
	_ = c.ShouldBindQuery(&v)
	if err := v.Validate(); err != nil {
		c.AbortWithStatusJSON(response.InvalidParameters.SetMessage(err.Error()).Slice())
		return
	}
	c.JSON(response.OK.Slice())
}
