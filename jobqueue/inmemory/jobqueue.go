package inmemory

import (
	"sync"

	"github.com/fireworq/fireworq/jobqueue"
	"github.com/fireworq/fireworq/jobqueue/logger"
)

type jobQueue struct {
	sync.Mutex
	queue *queue
}

// New creates a jobqueue.Impl which uses in-memory data store.
func New() jobqueue.Impl { _ = "STUB: not implemented"; return *new(jobqueue.Impl) }

func (q *jobQueue) Start() { _ = "STUB: not implemented"; return }

func (q *jobQueue) Stop() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (q *jobQueue) Push(j jobqueue.IncomingJob) (jobqueue.Job, error) {
	_ = "STUB: not implemented"
	return *new(jobqueue.Job), nil
}

func (q *jobQueue) Pop(limit uint) ([]jobqueue.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *jobQueue) Delete(job jobqueue.Job) {
	_ = "STUB: not implemented"
	// Do nothing; the job is deleted from the queue on Pop().
	return
}

func (q *jobQueue) Update(completedJob jobqueue.Job, next jobqueue.NextInfo) {
	_ = "STUB: not implemented"
	return
}

func (q *jobQueue) IsActive() bool { _ = "STUB: not implemented"; return false }

type job struct {
	jobqueue.IncomingJob
	id         uint64
	createdAt  uint64
	nextTry    uint64
	retryCount uint
	failCount  uint
}

func newJob(j jobqueue.IncomingJob) *job { _ = "STUB: not implemented"; return nil }

func (j *job) ID() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *job) CreatedAt() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *job) Status() string { _ = "STUB: not implemented"; return "" }

func (j *job) NextTry() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *job) RetryCount() uint { _ = "STUB: not implemented"; return 0 }

func (j *job) FailCount() uint { _ = "STUB: not implemented"; return 0 }

func (j *job) ToLoggable() logger.LoggableJob {
	_ = "STUB: not implemented"
	return *new(logger.LoggableJob)
}

type queue []*job

func (q queue) Len() int { _ = "STUB: not implemented"; return 0 }

func (q queue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (q queue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (q *queue) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (q *queue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

var lastID uint64
