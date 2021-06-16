package console

import (
    "log"
    "time"
)

type Example struct {
}

func (e Example) Run() {
    log.Println(time.Now().Format("2006-01-02 15:04:05"))
}
