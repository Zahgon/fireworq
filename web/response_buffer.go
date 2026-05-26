package web

import (
	"bytes"
	"net/http"
)

type responseBuffer struct {
	buf    bytes.Buffer
	status int
	header http.Header
}

func (rb *responseBuffer) Write(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (rb *responseBuffer) WriteHeader(status int) { _ = "STUB: not implemented"; return }

func (rb *responseBuffer) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (rb *responseBuffer) WriteTo(w http.ResponseWriter) error {
	_ = "STUB: not implemented"
	return nil
}
