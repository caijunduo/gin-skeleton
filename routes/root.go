package routes

import (
	"github.com/gin-gonic/gin"
	"skeleton/internal"
)

func init() {
	Include(Root)
}

func Root(r *gin.Engine) {
	internal.Example.RouteGroup(r.Group("example"))
}
