package job

import (
	"skeleton/logger"
)

func Recovery() {
	if rc := recover(); rc != nil {
		logger.Job.Error(rc)
	}
}
