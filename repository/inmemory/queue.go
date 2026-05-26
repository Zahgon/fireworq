package inmemory

import (
	"sync"

	"github.com/fireworq/fireworq/model"
	"github.com/fireworq/fireworq/repository"
)

type queueStorage struct {
	sync.RWMutex
	m        map[string]model.Queue
	revision uint64
}

var qs = &queueStorage{m: make(map[string]model.Queue)}

type queueRepository struct{}

// NewQueueRepository creates a new repository.QueueRepository which
// uses in-memory data store.
func NewQueueRepository() repository.QueueRepository {
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

func (r *queueRepository) DeleteByName(name string) error { _ = "STUB: not implemented"; return nil }

func (r *queueRepository) updateRevision() { _ = "STUB: not implemented"; return }

func (r *queueRepository) Revision() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
