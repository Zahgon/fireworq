package inmemory

import (
	"sync"

	"github.com/fireworq/fireworq/model"
	"github.com/fireworq/fireworq/repository"
)

type routingStorage struct {
	sync.RWMutex
	m        map[string]string
	revision uint64
}

var rs = &routingStorage{m: make(map[string]string)}

type routingRepository struct{}

// NewRoutingRepository creates a new repository.RoutingRepository
// which uses in-memory data store.
func NewRoutingRepository() repository.RoutingRepository {
	_ = "STUB: not implemented"
	return *new(repository.RoutingRepository)
}

func (r *routingRepository) Add(jobCategory string, queueName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *routingRepository) FindAll() ([]model.Routing, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *routingRepository) FindQueueNameByJobCategory(category string) string {
	_ = "STUB: not implemented"
	return ""
}

func (r *routingRepository) DeleteByJobCategory(category string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *routingRepository) updateRevision() { _ = "STUB: not implemented"; return }

func (r *routingRepository) Revision() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *routingRepository) Reload() error { _ = "STUB: not implemented"; return nil }
