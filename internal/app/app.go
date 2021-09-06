package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func App(c *gin.Context) {
	c.String(http.StatusOK, "Hello Skeleton")
}
