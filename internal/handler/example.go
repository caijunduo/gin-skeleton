package handler

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

func Example(c *gin.Context) {
    log.Println("main")
    c.String(http.StatusOK, "Hello Skeleton")
}
