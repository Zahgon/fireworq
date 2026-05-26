package web

import (
	"net/http"

	"github.com/fireworq/fireworq/dispatcher"
	"github.com/fireworq/fireworq/jobqueue"
)

func (app *Application) serveQueueList(w http.ResponseWriter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (app *Application) serveQueueListStats(w http.ResponseWriter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (app *Application) serveQueue(w http.ResponseWriter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (app *Application) serveQueueNode(w http.ResponseWriter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (app *Application) serveQueueStats(w http.ResponseWriter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (app *Application) serveQueueGrabbed(w http.ResponseWriter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (app *Application) serveQueueWaiting(w http.ResponseWriter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (app *Application) serveQueueDeferred(w http.ResponseWriter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (app *Application) serveQueueJobs(find func(jobqueue.Inspector, uint, string, jobqueue.SortOrder) (*jobqueue.InspectedJobs, error), w http.ResponseWriter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (app *Application) serveQueueFailed(w http.ResponseWriter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (app *Application) serveQueueJob(w http.ResponseWriter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (app *Application) serveQueueFailedJob(w http.ResponseWriter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// JobqueueStats is an alias to pointer type of jobqueue.Stats.
type JobqueueStats = *jobqueue.Stats

// DispatcherStats is an alias to pointer type of dispatcher.Stats.
type DispatcherStats = *dispatcher.Stats

// Stats contains queue statistics and worker statistics.
type Stats struct {
	JobqueueStats
	DispatcherStats
	ActiveNodes int64 `json:"active_nodes"`
}
