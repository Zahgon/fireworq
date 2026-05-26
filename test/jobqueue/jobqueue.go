package jqtest

import (
	"testing"

	"github.com/fireworq/fireworq/jobqueue"
)

type job struct {
	category   string
	url        string
	payload    string
	retryCount uint
	retryDelay uint
	timeout    uint
}

func (j *job) Category() string { _ = "STUB: not implemented"; return "" }

func (j *job) URL() string { _ = "STUB: not implemented"; return "" }

func (j *job) Payload() string { _ = "STUB: not implemented"; return "" }

func (j *job) NextDelay() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *job) NextTry() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *job) RetryCount() uint { _ = "STUB: not implemented"; return 0 }

func (j *job) RetryDelay() uint { _ = "STUB: not implemented"; return 0 }

func (j *job) Timeout() uint { _ = "STUB: not implemented"; return 0 }

const retryCount = 3

func newTestJob(category, url, data string) jobqueue.IncomingJob {
	_ = "STUB: not implemented"
	return *new(jobqueue.IncomingJob)
}

type nextJob struct {
	jobqueue.Job
	nextDelay uint64
}

func (j *nextJob) NextDelay() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *nextJob) NextTry() uint64 { _ = "STUB: not implemented"; return 0 }

func (j *nextJob) RetryCount() uint { _ = "STUB: not implemented"; return 0 }

func (j *nextJob) FailCount() uint { _ = "STUB: not implemented"; return 0 }

// Subtest is an interface of a test function where the queue is
// assumed to be empty before running the test.
type Subtest func(t *testing.T, jq jobqueue.Impl)

// SubtestRunner is an interface of a function that runs multiple
// subtests.  An instance of this interface is required to be a
// function which runs specified subtests with truncating data store
// contents before running each subtest.
type SubtestRunner func(t *testing.T, db, q string, tests []Subtest)

// TestSubtests runs predefined subtests by the specified runner.
func TestSubtests(t *testing.T, runner SubtestRunner) { _ = "STUB: not implemented"; return }

func subtestActive(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestEmpty(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestPush1(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestPop1(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestPopOrder(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestPopPartially(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestPopMulti(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestDelete1(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestDeletePartially(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestDeleteMulti(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestUpdate1(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestUpdatePartially(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestUpdateMulti(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestAsyncPop1(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestAsyncDelete1(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }

func subtestAsyncUpdate1(t *testing.T, jq jobqueue.Impl) { _ = "STUB: not implemented"; return }
