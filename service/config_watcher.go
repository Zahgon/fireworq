package service

type configWatcher struct {
	revision     func() (uint64, error)
	reload       func()
	stopC        chan struct{}
	stoppedC     chan struct{}
	lastRevision uint64
}

func newConfigWatcher(revision func() (uint64, error), reload func()) *configWatcher {
	_ = "STUB: not implemented"
	return nil
}

func (w *configWatcher) start(interval uint) { _ = "STUB: not implemented"; return }

func (w *configWatcher) stop() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (w *configWatcher) loop(interval uint) { _ = "STUB: not implemented"; return }
