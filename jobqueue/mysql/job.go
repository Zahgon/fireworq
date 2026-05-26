package mysql

import (
	"github.com/fireworq/fireworq/jobqueue"
	"github.com/fireworq/fireworq/jobqueue/logger"
)

// incomingJob : implements the following interfaces
// - jobqueue.IncomingJob
// - jobqueue.Job
// - logger.LoggableJob
type incomingJob struct {
	jobqueue.IncomingJob
	id uint64
}

func (j *incomingJob) ID() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *incomingJob) FailCount() uint { _ = "STUB: not implemented"; return 0 }

func (j *incomingJob) Status() string { _ = "STUB: not implemented"; return "" }

func (j *incomingJob) CreatedAt() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *incomingJob) NextDelay() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *incomingJob) NextTry() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *incomingJob) ToLoggable() logger.LoggableJob {
	_ = "STUB: not implemented"

	// job : implements the following interfaces
	// - jobqueue.Job
	// - logger.LoggableJob
	return *new(logger.LoggableJob)
}

type job struct {
	id         uint64
	category   string
	url        string
	payload    string
	status     string
	createdAt  uint64 // milliseconds
	nextTry    uint64 // milliseconds
	timeout    uint   // seconds
	retryDelay uint   // seconds
	retryCount uint
	failCount  uint
}

func (j *job) ID() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *job) Category() string { _ = "STUB: not implemented"; return "" }

func (j *job) URL() string { _ = "STUB: not implemented"; return "" }

func (j *job) Payload() string { _ = "STUB: not implemented"; return "" }

func (j *job) NextTry() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *job) RetryCount() uint { _ = "STUB: not implemented"; return 0 }

func (j *job) RetryDelay() uint { _ = "STUB: not implemented"; return 0 }

func (j *job) FailCount() uint { _ = "STUB: not implemented"; return 0 }

func (j *job) Timeout() uint { _ = "STUB: not implemented"; return 0 }

func (j *job) Status() string { _ = "STUB: not implemented"; return "" }

func (j *job) CreatedAt() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *job) ToLoggable() logger.LoggableJob {
	_ = "STUB: not implemented"
	return *new(logger.LoggableJob)
}
