package accountInternal

import (
    "github.com/gin-gonic/gin"
    "github.com/gookit/event"
    "skeleton/listener"
    "skeleton/response"
)

func GetInfo(c *gin.Context) {
    _, _ = event.Trigger(listener.AccountExampleStart, event.M{})
    c.IndentedJSON(response.OK.SetData("Hello，Skeleton.").Slice())
    _, _ = event.Trigger(listener.AccountExampleEnd, event.M{})
}

func GetMessageInfo(c *gin.Context) {
    c.IndentedJSON(response.OK.SetData("Hello, Skeleton Message.").Slice())
}
