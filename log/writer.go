package log

import (
	"io"
	"os"
	"sync"
)

// Writer is an io.Writer with Reopen() method.
type Writer interface {
	io.Writer
	Reopen() error
}

type writer struct {
	io.Writer
}

// New returns a new Writer which inherits w and does nothing on Reopen().
func New(w io.Writer) Writer { _ = "STUB: not implemented"; return *new(Writer) }

func (w *writer) Reopen() error { _ = "STUB: not implemented"; return nil }

type fileWriter struct {
	file *os.File
	mu   sync.RWMutex
}

// OpenFile opens a file and returns a Writer which writes to the file
// and reopens the file on Reopen().
func OpenFile(name string) (Writer, error) { _ = "STUB: not implemented"; return *new(Writer), nil }

func openFile(name string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *fileWriter) Reopen() error { _ = "STUB: not implemented"; return nil }

func (w *fileWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
