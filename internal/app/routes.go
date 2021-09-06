package app

import (
	"github.com/gin-gonic/gin"
	"skeleton/internal/middleware"
)

func Routes() *gin.Engine {
	r := gin.New()

	r.Use(middleware.Context)
	r.Use(middleware.Response)
	r.Use(middleware.Recovery)
	r.Use(middleware.RequestId)
	if !gin.IsDebugging() {
		r.Use(middleware.Logger)
	} else {
		r.Use(middleware.Cors)
	}
	r.Use(middleware.Validator)

	r.GET("/", App)

	return r
}
