package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HttpError interface {
	Code(int) HttpError
	Send(error) error
	NotFound() error
	Conflict() error
	Forbidden() error
	BadRequest() error
	Unauthorized() error
	InternalServerError() error
	NotImplemented() error
	SessionExpired() error
}

type httpError struct {
	err    error
	status int
}

func (h *httpError) Code(status int) HttpError {
	h.status = status
	return h
}

func (h *httpError) Send(err error) error {
	return echo.NewHTTPError(h.status, ErrorResponse(err.Error()))
}

func (h *httpError) BadRequest() error {
	h.status = http.StatusBadRequest
	return h.Send(h.err)
}

func (h *httpError) Unauthorized() error {
	h.status = http.StatusUnauthorized
	return h.Send(h.err)
}

func (h *httpError) Forbidden() error {
	h.status = http.StatusForbidden
	return h.Send(h.err)
}

func (h *httpError) NotFound() error {
	h.status = http.StatusNotFound
	return h.Send(h.err)
}

func (h *httpError) InternalServerError() error {
	h.status = http.StatusInternalServerError
	return h.Send(h.err)
}

func (h *httpError) Conflict() error {
	h.status = http.StatusConflict
	return h.Send(h.err)
}

func (h *httpError) NotImplemented() error {
	h.status = http.StatusNotImplemented
	return h.Send(h.err)
}

func (h *httpError) SessionExpired() error {
	h.status = 419
	return h.Send(h.err)
}

func HTTPError(err error) HttpError {
	return &httpError{
		err:    err,
		status: http.StatusInternalServerError,
	}
}
