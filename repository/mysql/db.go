//go:generate go-assets-builder -p mysql -o assets.go ../../data/repository

package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // initialize the driver
)

var schema []string

func init() {
	schema = []string{
		"/data/repository/mysql/schema/queue.sql",
		"/data/repository/mysql/schema/queue_throttle.sql",
		"/data/repository/mysql/schema/routing.sql",
		"/data/repository/mysql/schema/config_revision.sql",
	}
}

// Dsn returns the data source name of the storage specified in the
// configuration.
func Dsn() string { _ = "STUB: not implemented"; return "" }

// NewDB creates an instance of DB handler.
func NewDB() (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }
