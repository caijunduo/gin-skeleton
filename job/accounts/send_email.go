package accountJob

import (
	"skeleton/job"
	"skeleton/logger"
)

type SendEmail struct{}

func (s SendEmail) Run() {
	defer job.Recovery()
	logger.Job.Info("Job Send Email")
}
