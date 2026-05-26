package mysql

import (
	"database/sql"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog"
)

var (
	activatorLockWaitTimeout    = 10 * time.Second
	activatorActivationInterval = 1 * time.Second
)

type activator struct {
	queueName string
	cancel    atomic.Value
	stoppedC  chan struct{}
	stopped   uint32
	active    int32
	db        *sql.DB
	dsn       string
	logger    zerolog.Logger
}

type activation interface {
	queueName() string
	getDsn() string
}

func startActivator(q activation, onActivating func()) *activator {
	_ = "STUB: not implemented"
	return nil
}

func (a *activator) stop() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (a *activator) isActive() bool { _ = "STUB: not implemented"; return false }

func (a *activator) lockName() string { _ = "STUB: not implemented"; return "" }

// File private methods

func (a *activator) loop(onActivating func()) { _ = "STUB: not implemented"; return }

func (a *activator) activate(onActivating func()) (shouldRetry bool) {
	_ = "STUB: not implemented"
	return false
}

// This should not happen since sql.Open() won't try to
// connect to the DB and won't fail unless the DB driver name
// is invalid.  We just try again in case sql.Open() changes
// the behavior in future.

// Make sure that the queue is active since we have the lock.
// Without doing this, the queue won't be activated if the
// former call of getLock() failed to receive packets from the
// DB but `GET_LOCK` had actually been accepted by the DB.

// `GET_LOCK` timed out; just try again.

// This doesn't seem to happen to `GET_LOCK`:
// } else if e, ok := err.(*mysqldriver.MySQLError); ok && e.Number == 1205 {
// 	// Lock wait timeout (`lock_wait_timeout`) exceeded.
// 	// Just try again later.
// 	a.logger.Debug().Msg(err.Error())

// Connection failed (maybe DB server down).
// Try to reconnect later.

func (a *activator) connect() error { _ = "STUB: not implemented"; return nil }

func (a *activator) disconnect() { _ = "STUB: not implemented"; return }

func (a *activator) hasLock() bool { _ = "STUB: not implemented"; return false }

func (a *activator) getLock() error { _ = "STUB: not implemented"; return nil }

// Got NULL; means running out of memory or the thread was killed.

type lockError struct{}

func (err *lockError) Error() string { _ = "STUB: not implemented"; return "" }

type lockTimeoutError struct{}

func (err *lockTimeoutError) Error() string { _ = "STUB: not implemented"; return "" }

type stoppedError struct{}

func (err *stoppedError) Error() string { _ = "STUB: not implemented"; return "" }
