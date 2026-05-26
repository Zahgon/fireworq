package web

import (
	"encoding/json"
	"net/http"
)

func (app *Application) serveJob(w http.ResponseWriter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// IncomingJob describes a job to be pushed in a queue.
type IncomingJob struct {
	CategoryField string          `json:"category"`
	URLField      string          `json:"url"`
	PayloadField  json.RawMessage `json:"payload"`
	payloadField  string

	RunAfterField   uint `json:"run_after"`   // seconds
	TimeoutField    uint `json:"timeout"`     // seconds
	RetryDelayField uint `json:"retry_delay"` // seconds
	MaxRetriesField uint `json:"max_retries"`
}

// PushResult describes a job pushed to a queue.
type PushResult struct {
	ID        uint64 `json:"id"`
	QueueName string `json:"queue_name"`
	IncomingJob
}

// Category returns the category of the job.
func (job *IncomingJob) Category() string { _ = "STUB: not implemented"; return "" }

// URL returns the URL of the job.
func (job *IncomingJob) URL() string { _ = "STUB: not implemented"; return "" }

// DecodePayload decodes PayloadField of the job.
//
// If PayloadField starts and ends with ", then it is decoded as a
// JSON string.  If PayloadField is "null" then it is decoded to an
// empty string.  Otherwise, the decoded value is the raw string of
// PayloadField.
//
// The decoded value can be retrieved by Payload() method.
func (job *IncomingJob) DecodePayload() error { _ = "STUB: not implemented"; return nil }

// Payload returns a decoded value of PayloadField.
func (job *IncomingJob) Payload() string { _ = "STUB: not implemented"; return "" }

// NextDelay returns the delay for a next try of the job.
func (job *IncomingJob) NextDelay() uint64 { _ = "STUB: not implemented"; return 0 }

// RetryCount returns the max retries of the job.
func (job *IncomingJob) RetryCount() uint { _ = "STUB: not implemented"; return 0 }

// RetryDelay returns the delay for retries of the job.
func (job *IncomingJob) RetryDelay() uint { _ = "STUB: not implemented"; return 0 }

// Timeout returns the timeout of the job.
func (job *IncomingJob) Timeout() uint { _ = "STUB: not implemented"; return 0 }
