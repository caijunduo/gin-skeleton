package logger

import (
	"github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var Job *logrus.Logger

func init() {
	Job = logrus.New()
	Job.SetLevel(logrus.TraceLevel)
	Job.SetFormatter(&formatter.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		HideKeys:        true,
		FieldsOrder:     []string{"type"},
	})
	Job.AddHook(&loggerHook{typeName: "job"})
}
