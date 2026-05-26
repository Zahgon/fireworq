package service

import (
	"github.com/fireworq/fireworq/dispatcher"
	"github.com/fireworq/fireworq/jobqueue"
	"github.com/fireworq/fireworq/model"
)

// RunningQueue is an interface of a running queue, which is a job
// queue and its dispatcher combined.
type RunningQueue interface {
	jobqueue.JobQueue
	PollingInterval() uint
	MaxWorkers() uint
	WorkerStats() *dispatcher.Stats
	Deactivate() <-chan struct{}
}

type runningQueue struct {
	jobqueue.JobQueue
	dispatcher dispatcher.Dispatcher
}

func startJobQueue(q *model.Queue) *runningQueue { _ = "STUB: not implemented"; return nil }

func (q *runningQueue) Deactivate() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (q *runningQueue) Stop() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (q *runningQueue) Push(job jobqueue.IncomingJob) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (q *runningQueue) PollingInterval() uint { _ = "STUB: not implemented"; return 0 }

func (q *runningQueue) MaxWorkers() uint { _ = "STUB: not implemented"; return 0 }

func (q *runningQueue) WorkerStats() *dispatcher.Stats { _ = "STUB: not implemented"; return nil }
