package web

import (
	"io"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func graceful(server *http.Server, shutdownTimeout time.Duration) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

type server struct {
	addrs       []net.Addr
	makeHandler func(h http.Handler) http.Handler
	mux         *mux.Router
}

func newServer(out io.Writer) *server { _ = "STUB: not implemented"; return nil }

func (s *server) start() (*http.Server, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *server) startGracefully(shutdownTimeout time.Duration) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (s *server) handle(pattern string, h func(http.ResponseWriter, *http.Request) error) {
	_ = "STUB: not implemented"
	return
}
