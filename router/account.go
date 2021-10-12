package router

import (
	"github.com/gin-gonic/gin"
	accountInternal "skeleton/internal/accounts"
)

func init() {
	include("example", func(r *gin.RouterGroup) {
		r.GET("info", accountInternal.GetInfo)
		r.GET("message/info", accountInternal.GetMessageInfo)
	})
}
