package web

import (
	"net/http"
)

type handler func(w http.ResponseWriter, req *http.Request) error

func (f handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}
