package dispatcher

import (
	"github.com/fireworq/fireworq/dispatcher/kicker"
	"github.com/fireworq/fireworq/dispatcher/worker"
	"github.com/fireworq/fireworq/jobqueue"
	"github.com/fireworq/fireworq/model"

	"github.com/rs/zerolog"

	"golang.org/x/time/rate"
)

const defaultMinBufferSize = 1000

// Init initializes global parameters of dispatchers by configuration values.
//
// Configuration keys prefixed by "dispatch_" are considered.
func Init() {
	_ = "STUB: not implemented"

	// Config contains information to create a dispatcher instance.
	return
}

type Config struct {
	MinBufferSize uint
	Kicker        kicker.Config
	Worker        worker.Config
}

// Start creates and starts a new dispatcher instance with the current
// configuration.
//
// The instance watches a queue specified by q in a way specified by
// m.
func (cfg Config) Start(q JobQueue, m *model.Queue) Dispatcher {
	_ = "STUB: not implemented"
	return *new(Dispatcher)
}

// Dispatcher is an interface of dispatchers for some queue.
type Dispatcher interface {
	Stats() *Stats
	PollingInterval() uint
	MaxWorkers() uint
	MaxDispatchesPerSecond() float64
	MaxBurstSize() int
	Ping()
	Stop() <-chan struct{}
}

// Start creates and starts a new dispatcher instance with the default
// configuration.
func Start(q JobQueue, m *model.Queue) Dispatcher {
	_ = "STUB: not implemented"
	return *new(Dispatcher)
}

type dispatcher struct {
	jobqueue  JobQueue
	kicker    kicker.Kicker
	worker    worker.Worker
	kick      chan struct{}
	stop      chan struct{}
	stopped   chan struct{}
	jobBuffer chan jobqueue.Job
	sem       chan struct{}
	limiter   *rate.Limiter
	logger    zerolog.Logger
}

func (d *dispatcher) Kick() { _ = "STUB: not implemented"; return }

func (d *dispatcher) Ping() { _ = "STUB: not implemented"; return }

func (d *dispatcher) Stats() *Stats { _ = "STUB: not implemented"; return nil }

func (d *dispatcher) PollingInterval() uint { _ = "STUB: not implemented"; return 0 }

func (d *dispatcher) MaxWorkers() uint { _ = "STUB: not implemented"; return 0 }

func (d *dispatcher) MaxDispatchesPerSecond() float64 { _ = "STUB: not implemented"; return 0 }

func (d *dispatcher) MaxBurstSize() int { _ = "STUB: not implemented"; return 0 }

func (d *dispatcher) Stop() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (d *dispatcher) loop() { _ = "STUB: not implemented"; return }

func (d *dispatcher) popJobs() { _ = "STUB: not implemented"; return }

// JobQueue is an interface of a queue which can be watched by
// dispatchers.
type JobQueue interface {
	Pop(limit uint) ([]jobqueue.Job, error)
	Complete(job jobqueue.Job, res *jobqueue.Result)
	Name() string
}

// Stats contains statistics of a dispatcher.
type Stats struct {
	OutstandingJobs int64 `json:"outstanding_jobs"`
	TotalWorkers    int64 `json:"total_workers"`
	IdleWorkers     int64 `json:"idle_workers"`
}
