package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var HTTP zerolog.Logger

func init() {
	HTTP = log.Output(zerolog.ConsoleWriter{
		Out:         os.Stderr,
		FormatLevel: consoleFormatLevel("Http", false),
		TimeFormat:  "2006-01-02 15:04:05",
	}).With().Timestamp().Logger()
}
