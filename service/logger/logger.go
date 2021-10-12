package loggerService

import "log"

type Logger struct{}

func CreateLogger(logger Logger) {
    log.Println("LoggerService: CreateLogger")
}
