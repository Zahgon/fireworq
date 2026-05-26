package logger

import (
	"os"

	logwriter "github.com/fireworq/fireworq/log"

	"github.com/rs/zerolog"
)

var (
	// Writer is a log writer which logger.Info() or logger.Debug()
	// writes to.
	Writer = logwriter.New(os.Stdout)

	tag    string
	logger *zerolog.Logger = func() *zerolog.Logger {
		l := zerolog.Nop()
		return &l
	}()
)

// Init initializes global parameters of logger by configuration values.
//
// Configuration keys prefixed by "queue_log_" are considered.
func Init() { _ = "STUB: not implemented"; return }

func put(event *zerolog.Event, queue string, action string, j LoggableJob, msg string) {
	_ = "STUB: not implemented"
	return
}

// Info writes an INFO level log entry of a job action.
func Info(queue string, action string, j LoggableJob, msg string) {
	_ = "STUB: not implemented"
	return
}

// Debug writes a DEBUG level log entry of a job action.
func Debug(queue string, action string, j LoggableJob, msg string) {
	_ = "STUB: not implemented"
	return
}

// LoggableJob defines fields of a job to be written into the log.
type LoggableJob interface {
	Category() string
	URL() string
	Payload() string

	ID() uint64
	Status() string

	NextTry() uint64
	RetryCount() uint
	RetryDelay() uint
	FailCount() uint
	Timeout() uint

	CreatedAt() uint64
}

// Elapsed returns elapsed since the job is created in millisecond.
func Elapsed(j LoggableJob) int64 { _ = "STUB: not implemented"; return 0 }
