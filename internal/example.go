package internal

import (
	"github.com/gin-gonic/gin"
	"skeleton/cache"
	"skeleton/response"
)

type Example struct {
	exampleCache cache.Example
}

func (e Example) RouteGroup(r *gin.RouterGroup) {
	r.GET("", e.example)
}

func (e Example) example(c *gin.Context) {
	c.JSON(response.OK.Slice())
}
