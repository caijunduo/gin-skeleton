package logger

import (
    "github.com/sirupsen/logrus"
)

var HTTP *logrus.Logger

func init() {
    HTTP = logrus.New()
    HTTP.SetFormatter(&logrus.TextFormatter{
        DisableColors:   true,
        FullTimestamp:   true,
        TimestampFormat: "2006-01-02 15:04:05",
    })
    HTTP.AddHook(&loggerHook{typeName: "http"})
}
