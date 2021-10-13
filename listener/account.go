package listener

import (
    "github.com/gookit/event"
    "log"
)

var (
    AccountExample      = "account.example.*"
    AccountExampleStart = "account.example.start"
    AccountExampleEnd   = "account.example.end"
)

func init() {
    event.On(AccountExample, event.ListenerFunc(func(e event.Event) error {
        log.Println("Account Example")
        return nil
    }))
    event.On(AccountExampleStart, event.ListenerFunc(func(e event.Event) error {
        log.Println("Account Example Start")
        return nil
    }))
}
