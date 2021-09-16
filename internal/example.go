package internal

import (
	"github.com/gin-gonic/gin"
	"skeleton/response"
)

type Example struct{}

func (e Example) RouteGroup(r *gin.RouterGroup) {
	r.GET("", e.example)
}

func (e Example) example(c *gin.Context) {
	c.JSON(response.Example.Slice())
}
