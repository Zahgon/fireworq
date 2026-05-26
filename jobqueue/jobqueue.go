package jobqueue

import (
	"github.com/fireworq/fireworq/model"
)

// Impl is an interface of a job queue implementation.
type Impl interface {
	Start()
	Stop() <-chan struct{}
	Push(job IncomingJob) (Job, error)
	Pop(limit uint) ([]Job, error)
	Delete(job Job)
	Update(job Job, next NextInfo)
	IsActive() bool
}

// JobQueue is an interface of a job queue.
type JobQueue interface {
	Stop() <-chan struct{}
	Push(job IncomingJob) (uint64, error)
	Pop(limit uint) ([]Job, error)
	Complete(job Job, res *Result)

	Name() string

	IsActive() bool
	Node() (*Node, error)
	Stats() *Stats

	Inspector() (Inspector, bool)
	FailureLog() (FailureLog, bool)
}

// Start returns a job queue.
func Start(definition *model.Queue, q Impl) JobQueue {
	_ = "STUB: not implemented"
	return *new(JobQueue)
}

type jobQueue struct {
	name       string
	maxWorkers uint
	impl       Impl
	stats      *stats
}

func (q *jobQueue) Name() string { _ = "STUB: not implemented"; return "" }

func (q *jobQueue) Stop() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (q *jobQueue) Push(j IncomingJob) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (q *jobQueue) Pop(limit uint) ([]Job, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *jobQueue) Complete(job Job, res *Result) { _ = "STUB: not implemented"; return }

func (q *jobQueue) IsActive() bool { _ = "STUB: not implemented"; return false }

func (q *jobQueue) Node() (*Node, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *jobQueue) Stats() *Stats { _ = "STUB: not implemented"; return nil }

func (q *jobQueue) Inspector() (Inspector, bool) {
	_ = "STUB: not implemented"
	return *new(Inspector), false
}

func (q *jobQueue) FailureLog() (FailureLog, bool) {
	_ = "STUB: not implemented"
	return *new(FailureLog), false
}

// InactiveError is an error returned when Pop() is called on an
// inactive queue.
type InactiveError struct{}

func (e *InactiveError) Error() string { _ = "STUB: not implemented"; return "" }

// ConnectionClosedError is an error returned when Pop() is called but
// connection to a remote store has been lost.
type ConnectionClosedError struct{}

func (e *ConnectionClosedError) Error() string { _ = "STUB: not implemented"; return "" }
