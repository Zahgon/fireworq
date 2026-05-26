package worker

import (
	"github.com/fireworq/fireworq/jobqueue"

	"github.com/rs/zerolog"
)

var defaultUserAgent string

// HTTPInit initializes global parameters of HTTP workers by
// configuration values.
//
// Configuration keys prefixed by "dispatch_" are considered.
func HTTPInit() { _ = "STUB: not implemented"; return }

// HTTPWorker is a worker which handles a job as an HTTP POST request
// to the URL specified by the job.
type HTTPWorker struct {
	UserAgent string
	Logger    *zerolog.Logger
}

// NewWorker creates a new HTTP worker instance which inherits the
// configurations of the current one.
func (worker *HTTPWorker) NewWorker() Worker { _ = "STUB: not implemented"; return *new(Worker) }

// Work makes a POST request to job.URL and returns the result.
func (worker *HTTPWorker) Work(job jobqueue.Job) *jobqueue.Result {
	_ = "STUB: not implemented"
	return nil
}
