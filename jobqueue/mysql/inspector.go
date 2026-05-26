package mysql

import (
	"database/sql"

	"github.com/fireworq/fireworq/jobqueue"
)

type scanner interface {
	Scan(args ...interface{}) error
}

type inspector struct {
	db  *sql.DB
	sql *sqls
}

func (i *inspector) Delete(jobID uint64) error { _ = "STUB: not implemented"; return nil }

func (i *inspector) Find(jobID uint64) (*jobqueue.InspectedJob, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *inspector) FindAllGrabbed(limit uint, cursor string, order jobqueue.SortOrder) (*jobqueue.InspectedJobs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *inspector) FindAllWaiting(limit uint, cursor string, order jobqueue.SortOrder) (*jobqueue.InspectedJobs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *inspector) FindAllDeferred(limit uint, cursor string, order jobqueue.SortOrder) (*jobqueue.InspectedJobs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *inspector) findAllAsc(status string, minTime int64, maxTime int64, limit uint, cursor string) (*jobqueue.InspectedJobs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no job to report

// Emulate `ORDER BY next_try ASC, job_id ASC`, which causes
// `using filesort` together with `SELECT ~ WHERE ~ IN`.

func (i *inspector) findAllDesc(status string, minTime int64, maxTime int64, limit uint, cursor string) (*jobqueue.InspectedJobs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no job to report

// Emulate `ORDER BY next_try DESC, job_id DESC`, which causes
// `using filesort` together with `SELECT ~ WHERE ~ IN`.

func (i *inspector) scan(s scanner) (*jobqueue.InspectedJob, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
