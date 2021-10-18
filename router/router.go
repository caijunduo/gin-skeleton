package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	basicMiddleware "skeleton/middleware/basic"
	"skeleton/response"
)

type initialize func(r *gin.RouterGroup)

var initializes = make(map[string]initialize, 0)

//include Add initialization function
func include(prefix string, init initialize) {
	initializes[prefix] = init
}

func Setup() *gin.Engine {
	r := gin.New()
	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(response.NotFound.Slice())
		return
	})
	r.Use(basicMiddleware.Recovery())
	r.Use(basicMiddleware.Request)
	r.Use(basicMiddleware.Response)
	if gin.IsDebugging() {
		r.Use(cors.Default())
	} else {
		r.Use(basicMiddleware.Logger)
	}
	for prefix, init := range initializes {
		init(r.Group(prefix))
	}
	return r
}
