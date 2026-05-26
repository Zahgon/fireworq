package mysql

import (
	"database/sql"
	"sync"

	"github.com/fireworq/fireworq/model"
	"github.com/fireworq/fireworq/repository"
)

type routingRepository struct {
	sync.RWMutex
	db       *sql.DB
	routings map[string]string
}

// NewRoutingRepository creates a repository.RoutingRepository which uses
// MySQL as a data store.
func NewRoutingRepository(db *sql.DB) repository.RoutingRepository {
	_ = "STUB: not implemented"
	return *new(repository.RoutingRepository)
}

func (r *routingRepository) Add(jobCategory string, queueName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *routingRepository) FindQueueNameByJobCategory(category string) string {
	_ = "STUB: not implemented"
	return ""
}

func (r *routingRepository) FindAll() ([]model.Routing, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *routingRepository) DeleteByJobCategory(category string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *routingRepository) Revision() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *routingRepository) Reload() error { _ = "STUB: not implemented"; return nil }

func (r *routingRepository) updateRevision() error { _ = "STUB: not implemented"; return nil }
