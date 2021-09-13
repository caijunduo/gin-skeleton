package interfaces

import "github.com/gin-gonic/gin"

type InternalRoute interface {
	RouteGroup(group *gin.RouterGroup)
}
