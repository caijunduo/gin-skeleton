package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var Job zerolog.Logger

func init() {
	Job = log.Output(zerolog.ConsoleWriter{
		Out:         os.Stderr,
		FormatLevel: consoleFormatLevel("Job", false),
		TimeFormat:  "2006-01-02 15:04:05",
	}).With().Timestamp().Logger()
}
