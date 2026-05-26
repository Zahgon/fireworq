package web

import (
	"io"

	"github.com/fireworq/fireworq/jobqueue"
	"github.com/fireworq/fireworq/model"
	"github.com/fireworq/fireworq/repository"
	"github.com/fireworq/fireworq/service"
)

// Service is an interface of the application use case service.
type Service interface {
	Stop() <-chan struct{}
	GetJobQueue(qn string) (service.RunningQueue, bool)
	DeleteJobQueue(qn string) error
	AddJobQueue(q *model.Queue) error
	Push(job jobqueue.IncomingJob) (*service.PushResult, error)
}

// Application is an interface of the application.
type Application struct {
	AccessLogWriter   io.Writer
	Version           string
	Service           Service
	QueueRepository   repository.QueueRepository
	RoutingRepository repository.RoutingRepository
}

func (app *Application) newServer() *server { _ = "STUB: not implemented"; return nil }

// Serve starts the application Web server.
func (app *Application) Serve() { _ = "STUB: not implemented"; return }

func shutdownTimeout() uint { _ = "STUB: not implemented"; return 0 }
