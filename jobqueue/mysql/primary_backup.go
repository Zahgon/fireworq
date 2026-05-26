package mysql

import (
	"github.com/fireworq/fireworq/jobqueue"
	"github.com/fireworq/fireworq/model"
)

type primaryBackupJobQueue struct {
	*jobQueue
	activator *activator
}

// NewPrimaryBackup creates a jobqueue.Impl which uses MySQL as a data
// store and restricts only one node to be active in a cluster.
//
// Inactive nodes become backup nodes, which will be active when the
// active node dies.
func NewPrimaryBackup(definition *model.Queue, dsn string) jobqueue.Impl {
	_ = "STUB: not implemented"
	return *new(jobqueue.Impl)
}

func (q *primaryBackupJobQueue) Start() { _ = "STUB: not implemented"; return }

func (q *primaryBackupJobQueue) Stop() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (q *primaryBackupJobQueue) IsActive() bool { _ = "STUB: not implemented"; return false }

func (q *primaryBackupJobQueue) Pop(limit uint) ([]jobqueue.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *primaryBackupJobQueue) Node() (*jobqueue.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Strip the port number

// activation interface

func (q *primaryBackupJobQueue) queueName() string { _ = "STUB: not implemented"; return "" }

func (q *primaryBackupJobQueue) getDsn() string { _ = "STUB: not implemented"; return "" }
