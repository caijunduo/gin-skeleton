package job

import (
	"skeleton/logger"
)

func Recovery() {
	if rc := recover(); rc != nil {
		logger.Job.Panic().Stack().Err(rc.(error))
	}
}
