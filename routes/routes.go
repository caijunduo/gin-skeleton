package routes

import (
	"github.com/gin-gonic/gin"
	"skeleton/middleware"
)

type option func(r *gin.Engine)

var options []option

func Include(opts ...option) {
	options = append(options, opts...)
}

func Init() *gin.Engine {
	r := gin.New()
	r.Use(
		middleware.Context,
		middleware.Recovery,
		middleware.Response,
		middleware.RequestId,
		middleware.Validator,
	)
	if !gin.IsDebugging() {
		r.Use(middleware.Logger)
	} else {
		r.Use(middleware.Cors)
	}
	for _, opt := range options {
		opt(r)
	}
	return r
}
