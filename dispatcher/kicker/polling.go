package kicker

// PollingKicker is a builder of a Kicker which kicks a Kickable
// repeatedly on some interval.
type PollingKicker struct {
	Interval uint
}

// NewKicker creates a new polling kicker instance.
func (cfg *PollingKicker) NewKicker() Kicker { _ = "STUB: not implemented"; return *new(Kicker) }

type pollingKicker struct {
	interval uint
	started  uint32
	stop     chan struct{}
	stopped  chan struct{}
}

func (k *pollingKicker) Start(kickable Kickable) { _ = "STUB: not implemented"; return }

func (k *pollingKicker) Stop() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (k *pollingKicker) Ping() {
	_ = "STUB: not implemented"
	// ignore; do nothing
	return
}

func (k *pollingKicker) PollingInterval() uint { _ = "STUB: not implemented"; return 0 }

func (k *pollingKicker) loop(kickable Kickable) { _ = "STUB: not implemented"; return }
