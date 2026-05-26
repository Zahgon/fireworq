package mysql

import (
	"database/sql"

	"github.com/fireworq/fireworq/model"
	"github.com/fireworq/fireworq/repository"
)

type queueRepository struct {
	db *sql.DB
}

// NewQueueRepository creates a repository.QueueRepository which uses
// MySQL as a data store.
func NewQueueRepository(db *sql.DB) repository.QueueRepository {
	_ = "STUB: not implemented"
	return *new(repository.QueueRepository)
}

func (r *queueRepository) Add(q *model.Queue) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *queueRepository) FindAll() ([]model.Queue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *queueRepository) FindByName(name string) (*model.Queue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type queueThrottle struct {
	maxDispatchesPerSecond float64
	maxBurstSize           uint
}

func (r *queueRepository) findQueueThrottles(names []string) (map[string]queueThrottle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *queueRepository) DeleteByName(name string) error { _ = "STUB: not implemented"; return nil }

func (r *queueRepository) Revision() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *queueRepository) updateRevision() error { _ = "STUB: not implemented"; return nil }
