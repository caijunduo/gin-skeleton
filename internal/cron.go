package internal

import (
    "github.com/robfig/cron/v3"
    "skeleton/internal/console"
)

func Crontab() error {
    c := cron.New(cron.WithSeconds())

    c.AddJob("* * * * * *", console.Example{})

    c.Start()

    defer c.Stop()

    select {}
}
