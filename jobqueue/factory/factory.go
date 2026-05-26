package factory

import (
	"github.com/fireworq/fireworq/jobqueue"
	"github.com/fireworq/fireworq/model"
)

// IncomingJob imitates IncomingJob in jobqueue package: factory
// package is intended to be used as a jobqueue package (by import
// jobqueue ".../fireworq/jobqueue/factory" since the only reason for
// having a separate package is to avoid cyclic import with a driver
// package such as mysql.
type IncomingJob = jobqueue.IncomingJob

// JobQueue imitates JobQueue in jobqueue package: factory package is
// intended to be used as a jobqueue package (by import jobqueue
// ".../fireworq/jobqueue/factory" since the only reason for having a
// separate package is to avoid cyclic import with a driver package
// such as mysql.
type JobQueue = jobqueue.JobQueue

// NewImpl creates a new jobqueue.Impl instance according to the value
// of "driver" configuration.
func NewImpl(q *model.Queue) jobqueue.Impl { _ = "STUB: not implemented"; return *new(jobqueue.Impl) }

// Start creates and starts a new JobQueue instance whose
// implementation is decided by the value of "driver" configuration.
func Start(q *model.Queue) JobQueue { _ = "STUB: not implemented"; return *new(JobQueue) }
