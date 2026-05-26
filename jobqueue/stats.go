package jobqueue

import (
	"github.com/paulbellamy/ratecounter"
)

// Stats describes queue statistics.
type Stats struct {
	TotalPushes            int64 `json:"total_pushes"`
	TotalPops              int64 `json:"total_pops"`
	TotalSuccesses         int64 `json:"total_successes"`
	TotalFailures          int64 `json:"total_failures"`
	TotalPermanentFailures int64 `json:"total_permanent_failures"`
	TotalCompletes         int64 `json:"total_completes"`
	TotalElapsed           int64 `json:"total_elapsed"`
	PushesPerSecond        int64 `json:"pushes_per_second"`
	PopsPerSecond          int64 `json:"pops_per_second"`
}

type stats struct {
	totalPushes            int64
	totalPops              int64
	totalSuccesses         int64
	totalFailures          int64
	totalPermanentFailures int64
	totalCompletes         int64
	totalElapsed           int64
	pushesPerSecond        *ratecounter.RateCounter
	popsPerSecond          *ratecounter.RateCounter
}

func newStats() *stats { _ = "STUB: not implemented"; return nil }

func (s *stats) push(num int64) { _ = "STUB: not implemented"; return }

func (s *stats) pop(num int64) { _ = "STUB: not implemented"; return }

func (s *stats) succeed(num int64) { _ = "STUB: not implemented"; return }

func (s *stats) fail(num int64) { _ = "STUB: not implemented"; return }

func (s *stats) permanentlyFail(num int64) { _ = "STUB: not implemented"; return }

func (s *stats) complete(num int64) { _ = "STUB: not implemented"; return }

func (s *stats) elapsed(t int64) { _ = "STUB: not implemented"; return }

func (s *stats) export() *Stats { _ = "STUB: not implemented"; return nil }
