package web

import (
	"net/http"
)

type clientError interface {
	error
	clientError() string
	httpStatus() int
}

type simpleClientError int

func (e simpleClientError) httpStatus() int { _ = "STUB: not implemented"; return 0 }

func (e simpleClientError) Error() string { _ = "STUB: not implemented"; return "" }

func (e simpleClientError) clientError() string { _ = "STUB: not implemented"; return "" }

func (e simpleClientError) WithDetail(detail string) *detailedClientError {
	_ = "STUB: not implemented"
	return nil
}

type detailedClientError struct {
	status  int
	message string
}

func (e *detailedClientError) httpStatus() int { _ = "STUB: not implemented"; return 0 }

func (e *detailedClientError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *detailedClientError) clientError() string { _ = "STUB: not implemented"; return "" }

type serverError interface {
	error
	serverError() string
	httpStatus() int
}

type simpleServerError int

func (e simpleServerError) httpStatus() int { _ = "STUB: not implemented"; return 0 }

func (e simpleServerError) Error() string { _ = "STUB: not implemented"; return "" }

func (e simpleServerError) serverError() string { _ = "STUB: not implemented"; return "" }

func (e simpleServerError) WithDetail(detail string) *detailedServerError {
	_ = "STUB: not implemented"
	return nil
}

type detailedServerError struct {
	status  int
	message string
}

func (e *detailedServerError) httpStatus() int { _ = "STUB: not implemented"; return 0 }

func (e *detailedServerError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *detailedServerError) serverError() string { _ = "STUB: not implemented"; return "" }

const (
	errMethodNotAllowed    = simpleClientError(http.StatusMethodNotAllowed)
	errNotFound            = simpleClientError(http.StatusNotFound)
	errBadRequest          = simpleClientError(http.StatusBadRequest)
	errNotImplemented      = simpleServerError(http.StatusNotImplemented)
	errInternalServerError = simpleServerError(http.StatusInternalServerError)
)
