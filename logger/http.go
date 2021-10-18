package logger

import (
	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var HTTP *logrus.Logger

func init() {
	HTTP = logrus.New()
	HTTP.SetLevel(logrus.TraceLevel)
	HTTP.SetFormatter(&formatter.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		HideKeys:        true,
		FieldsOrder:     []string{"type"},
	})
	HTTP.AddHook(&loggerHook{typeName: "http"})
}
