package jobqueue

import (
	"github.com/fireworq/fireworq/jobqueue/logger"
)

// IncomingJob is an interface of incoming jobs.
type IncomingJob interface {
	Category() string
	URL() string
	Payload() string

	NextDelay() uint64 // milliseconds
	Timeout() uint     // seconds
	RetryDelay() uint  // seconds
	RetryCount() uint
}

// Job is an interface of jobs.
type Job interface {
	URL() string
	Payload() string
	Timeout() uint

	RetryCount() uint
	RetryDelay() uint
	FailCount() uint

	ToLoggable() logger.LoggableJob
}

// completedJob : implements the following interfaces
// - Job
// - logger.LoggableJob
type completedJob struct {
	Job
	failed uint
}

func (j *completedJob) FailCount() uint { _ = "STUB: not implemented"; return 0 }

func (j *completedJob) Status() string { _ = "STUB: not implemented"; return "" }

func (j *completedJob) ToLoggable() logger.LoggableJob {
	_ = "STUB: not implemented"
	return *new(logger.LoggableJob)
}

func (j *completedJob) canRetry() bool { _ = "STUB: not implemented"; return false }

type loggableCompletedJob struct {
	logger.LoggableJob
	status    string
	failCount uint
}

func (j *loggableCompletedJob) Status() string { _ = "STUB: not implemented"; return "" }

func (j *loggableCompletedJob) FailCount() uint {
	_ = "STUB: not implemented"

	// nextJob : implements the following interfaces
	// - NextInfo
	return 0
}

type nextJob struct {
	job Job
}

func (j *nextJob) NextDelay() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *nextJob) RetryCount() uint { _ = "STUB: not implemented"; return 0 }

func (j *nextJob) FailCount() uint { _ = "STUB: not implemented"; return 0 }

// NextInfo describes information of a retry.
type NextInfo interface {
	NextDelay() uint64
	RetryCount() uint
	FailCount() uint
}
