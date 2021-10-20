package accountInternal

import (
	"github.com/bamzi/jobrunner"
	"github.com/gookit/event"
	"skeleton/context"
	accountJob "skeleton/job/accounts"
	"skeleton/listener"
	"skeleton/logger"
	"skeleton/response"
)

func GetInfo(c *context.Context) {
	_, _ = event.Trigger(listener.AccountExampleStart, event.M{})
	jobrunner.Now(accountJob.SendEmail{})
	logger.HTTP.Debug().Msg("Run after SendEmail Job")
	c.IndentedJSON(response.OK.Slice())
	_, _ = event.Trigger(listener.AccountExampleEnd, event.M{})
}

func GetMessageInfo(c *context.Context) {
	c.IndentedJSON(response.OK.SetData("Hello, Skeleton Message.").Slice())
}
