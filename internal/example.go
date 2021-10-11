package internal

import (
    "github.com/gin-gonic/gin"
    accountService "skeleton/account_service"
    "skeleton/response"
)

type Example struct{}

func (e Example) RouteGroup(r *gin.RouterGroup) {
    r.GET("", e.example)
}

func (e Example) example(c *gin.Context) {
    accountService.CreateAccount(accountService.Account{})
    c.JSON(response.OK.Slice())
}
