package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Example(c *gin.Context) {
    c.String(http.StatusOK, "Hello Skeleton")
}
