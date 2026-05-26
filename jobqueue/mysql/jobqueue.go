package mysql

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql" // initialize the driver
	"github.com/rs/zerolog"

	"github.com/fireworq/fireworq/jobqueue"
	"github.com/fireworq/fireworq/model"
)

// Dsn returns the data source name of the storage specified in the
// configuration.
func Dsn() string { _ = "STUB: not implemented"; return "" }

type jobQueue struct {
	name    string
	dsn     string
	sql     *sqls
	db      *sql.DB
	dbPop   *sql.DB
	mu      sync.RWMutex
	stopped uint32
	logger  zerolog.Logger
}

// New creates a jobqueue.Impl which uses MySQL as a data store.
func New(definition *model.Queue, dsn string) jobqueue.Impl {
	_ = "STUB: not implemented"
	return *new(jobqueue.Impl)
}

func newJobQueue(definition *model.Queue, dsn string) *jobQueue {
	_ = "STUB: not implemented"
	return nil
}

func (q *jobQueue) Start() { _ = "STUB: not implemented"; return }

func (q *jobQueue) Stop() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (q *jobQueue) IsActive() bool { _ = "STUB: not implemented"; return false }

func (q *jobQueue) Push(j jobqueue.IncomingJob) (jobqueue.Job, error) {
	_ = "STUB: not implemented"
	return *new(jobqueue.Job), nil
}

func (q *jobQueue) Pop(limit uint) ([]jobqueue.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. Pre-SELECT jobs to grab.  We should not `SELECT ~ FOR UPDATE`
// here because it blocks `Push()` due to a gap lock.

// no job to grab

// to avoid gap locks

// 2. SELECT jobs to grab FOR UPDATEing their status.

// The number of jobs may reduce when they are grabbed in another
// thread right after the preselection.  This is unlikely to
// happen to a single dispatcher, though.

// 3. UPDATE the status of jobs.

// Emulate `ORDER BY next_try ASC`, which causes `using filesort`
// together with `SELECT ~ WHERE ~ IN`.

func (q *jobQueue) Delete(completedJob jobqueue.Job) { _ = "STUB: not implemented"; return }

func (q *jobQueue) Update(completedJob jobqueue.Job, next jobqueue.NextInfo) {
	_ = "STUB: not implemented"
	return
}

func (q *jobQueue) Recover() { _ = "STUB: not implemented"; return }

func (q *jobQueue) Inspector() jobqueue.Inspector {
	_ = "STUB: not implemented"
	return *new(jobqueue.Inspector)
}

func (q *jobQueue) FailureLog() jobqueue.FailureLog {
	_ = "STUB: not implemented"
	return *new(jobqueue.FailureLog)
}

func (q *jobQueue) Node() (*jobqueue.Node, error) { _ = "STUB: not implemented"; return nil, nil }

// Strip the port number

func (q *jobQueue) connect() { _ = "STUB: not implemented"; return }

// Restrict connections to prevent connection ID from being changed.

func (q *jobQueue) disconnect() { _ = "STUB: not implemented"; return }
