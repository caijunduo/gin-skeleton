package listener

import (
	"github.com/gookit/event"
	"skeleton/logger"
)

var (
	AccountExampleStart = "account.example.start"
	AccountExampleEnd   = "account.example.end"
)

func init() {
	event.On(AccountExampleStart, event.ListenerFunc(func(e event.Event) error {
		logger.HTTP.Info("Account Example Start")
		return nil
	}))
	event.On(AccountExampleEnd, event.ListenerFunc(func(e event.Event) error {
		logger.HTTP.Info("Account Example End")
		return nil
	}))
}
