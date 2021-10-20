package logger

import (
	"bytes"
	"fmt"
	"github.com/rs/zerolog"
	"strings"
)

const (
	colorBlack = iota + 30
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite

	colorBold     = 1
	colorDarkGray = 90
)

func colorize(s interface{}, c int, disabled bool) string {
	if disabled {
		return fmt.Sprintf("%s", s)
	}
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, s)
}

func consoleFormatLevel(app string, noColor bool) zerolog.Formatter {
	return func(i interface{}) string {
		var l, a string
		if ll, ok := i.(string); ok {
			switch ll {
			case zerolog.LevelTraceValue:
				l = colorize("TRACE", colorMagenta, noColor)
				a = colorize(app, colorMagenta, noColor)
			case zerolog.LevelDebugValue:
				l = colorize("DEBUG", colorYellow, noColor)
				a = colorize(app, colorYellow, noColor)
			case zerolog.LevelInfoValue:
				l = colorize("INFO", colorGreen, noColor)
				a = colorize(app, colorGreen, noColor)
			case zerolog.LevelWarnValue:
				l = colorize("WARNING", colorRed, noColor)
				a = colorize(app, colorRed, noColor)
			case zerolog.LevelErrorValue:
				l = colorize(colorize("ERROR", colorRed, noColor), colorBold, noColor)
				a = colorize(colorize(app, colorRed, noColor), colorBold, noColor)
			case zerolog.LevelFatalValue:
				l = colorize(colorize("FATAL", colorRed, noColor), colorBold, noColor)
				a = colorize(colorize(app, colorRed, noColor), colorBold, noColor)
			case zerolog.LevelPanicValue:
				l = colorize(colorize("PANIC", colorRed, noColor), colorBold, noColor)
				a = colorize(colorize(app, colorRed, noColor), colorBold, noColor)
			default:
				l = colorize("LOG", colorBold, noColor)
				a = colorize(app, colorBold, noColor)
			}
		} else {
			l = strings.ToUpper(fmt.Sprintf("%s", i))[0:3]
		}
		b := bytes.NewBufferString(a)
		b.WriteString(":")
		b.WriteString(l)
		return b.String()
	}
}
