package logger

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

type loggerHook struct {
    environment string
    typeName    string
}

func (l *loggerHook) Levels() []logrus.Level {
    return logrus.AllLevels
}

func (l *loggerHook) Fire(entry *logrus.Entry) error {
    entry.Data["environment"] = gin.Mode()
    entry.Data["type"] = l.typeName
    return nil
}
