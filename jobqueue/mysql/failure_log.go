package mysql

import (
	"database/sql"

	"github.com/fireworq/fireworq/jobqueue"
)

type failureLog struct {
	db  *sql.DB
	sql *sqls
}

func (l *failureLog) Add(failed jobqueue.Job, result *jobqueue.Result) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *failureLog) Delete(failureID uint64) error { _ = "STUB: not implemented"; return nil }

func (l *failureLog) Find(failureID uint64) (*jobqueue.FailedJob, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *failureLog) FindAll(limit uint, cursor string) (*jobqueue.FailedJobs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *failureLog) FindAllRecentFailures(limit uint, cursor string) (*jobqueue.FailedJobs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *failureLog) findAllByQuery(query string, limit uint, cursor string) (*jobqueue.FailedJobs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *failureLog) scan(s scanner) (*jobqueue.FailedJob, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
