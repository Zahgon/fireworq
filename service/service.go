package service

import (
	"sync"

	jobqueue "github.com/fireworq/fireworq/jobqueue/factory"
	"github.com/fireworq/fireworq/model"
	"github.com/fireworq/fireworq/repository"
)

// PushResult is information of pushed job except those in
// jobqueue.IncomingJob itself.
type PushResult struct {
	ID        uint64
	QueueName string
}

// Service is an application use case service that manages running
// queues.
type Service struct {
	defaultQueueName string
	queue            repository.QueueRepository
	routing          repository.RoutingRepository
	runningQueues    map[string]RunningQueue
	mu               sync.Mutex
	muJob            sync.RWMutex
	queueW           *configWatcher
	routingW         *configWatcher
}

// NewService creates a new Service instance.
func NewService(repos *repository.Repositories) *Service { _ = "STUB: not implemented"; return nil }

// Stop stops all the running queues.
//
// This method should not be called more than once in the whole
// application.
func (s *Service) Stop() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// GetJobQueue returns a RunningQueue of name qn.  The second return
// value is false ff no queue is found.
//
// This method is goroutine safe.
func (s *Service) GetJobQueue(qn string) (RunningQueue, bool) {
	_ = "STUB: not implemented"
	return *new(RunningQueue), false
}

func (s *Service) getJobQueue(qn string) (RunningQueue, bool) {
	_ = "STUB: not implemented"
	return *new(RunningQueue), false
}

// DeleteJobQueue stops a running queue of name qn and removes it from
// the queue definition list.
//
// This method is goroutine safe.
func (s *Service) DeleteJobQueue(qn string) error { _ = "STUB: not implemented"; return nil }

// AddJobQueue defines a new queue and starts it.
//
// This method is goroutine safe.
func (s *Service) AddJobQueue(q *model.Queue) error { _ = "STUB: not implemented"; return nil }

// When throttling is configured, we use the fixed polling interval.
const throttleQueuePollingInterval = 100

func (s *Service) addJobQueue(q *model.Queue) error { _ = "STUB: not implemented"; return nil }

// Push pushes a job to a queue.  The target queue is determined by
// the category of the job and defined routings.
func (s *Service) Push(job jobqueue.IncomingJob) (*PushResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This happens when the queue definition is not in the cache
// but in the data store, which means it has been defined
// through another node.

func (s *Service) startup() { _ = "STUB: not implemented"; return }

func (s *Service) reloadQueues() { _ = "STUB: not implemented"; return }

func (s *Service) reloadRoutings() { _ = "STUB: not implemented"; return }

func (s *Service) initDefaultQueue(queueName string) error { _ = "STUB: not implemented"; return nil }

func (s *Service) putJobQueue(q *model.Queue) RunningQueue {
	_ = "STUB: not implemented"
	return *new(RunningQueue)
}

func (s *Service) deactivateQueues() { _ = "STUB: not implemented"; return }

func (s *Service) destroyQueues() { _ = "STUB: not implemented"; return }

func defaultPollingInterval() uint { _ = "STUB: not implemented"; return 0 }

func defaultMaxWorkers() uint { _ = "STUB: not implemented"; return 0 }

func configRefreshInterval() uint { _ = "STUB: not implemented"; return 0 }
