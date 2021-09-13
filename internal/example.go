package internal

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var Example = new(example)

type example struct{}

func (e example) RouteGroup(r *gin.RouterGroup) {
	r.GET("", e.gateway)
}

func (e example) beforeGateway() {
	log.Println("before")
}

func (e example) afterGateway() {
	log.Println("after")
}

func (e example) gateway(c *gin.Context) {
	e.beforeGateway()
	defer e.afterGateway()
	log.Println("main")
	c.String(http.StatusOK, "Example")
}
