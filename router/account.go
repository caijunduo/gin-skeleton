package router

import (
	"github.com/gin-gonic/gin"
	"skeleton/context"
	accountInternal "skeleton/internal/accounts"
)

func init() {
	include("example", func(r *gin.RouterGroup) {
		r.GET("info", context.Handler(accountInternal.GetInfo))
		r.GET("message/info", context.Handler(accountInternal.GetMessageInfo))
	})
}
