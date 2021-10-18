package accountInternal

import (
	"github.com/gookit/event"
	"skeleton/context"
	"skeleton/listener"
	"skeleton/response"
)

func GetInfo(c *context.Context) {
	_, _ = event.Trigger(listener.AccountExampleStart, event.M{})
	c.IndentedJSON(response.OK.Slice())
	_, _ = event.Trigger(listener.AccountExampleEnd, event.M{})
}

func GetMessageInfo(c *context.Context) {
	c.IndentedJSON(response.OK.SetData("Hello, Skeleton Message.").Slice())
}
