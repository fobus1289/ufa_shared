package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	ctx echo.Context
}

func Response(ctx echo.Context) *response {
	return &response{
		ctx,
	}
}

// Method for sending a response with a specific status code and data
func (r *response) Send(status int, res any) error {
	return r.ctx.JSON(status, res)
}

// Method for sending a response with a 200 OK status code
func (r *response) OK(res any) error {
	return r.Send(http.StatusOK, res)
}

// Method for sending a response with a 201 Created status code
func (r *response) Created(res any) error {
	return r.Send(http.StatusCreated, res)
}

// Method for sending a response with a 204 No Content status code
func (r *response) NoContent() error {
	return r.Send(http.StatusNoContent, nil)
}

// Method for sending a response with a 400 Bad Request status code
func (r *response) BadRequest(res any) error {
	return r.Send(http.StatusBadRequest, res)
}

// Method for sending a response with a 401 Unauthorized status code
func (r *response) Unauthorized(res any) error {
	return r.Send(http.StatusUnauthorized, res)
}

// Method for sending a response with a 403 Forbidden status code
func (r *response) Forbidden(res any) error {
	return r.Send(http.StatusForbidden, res)
}

// Method for sending a response with a 404 Not Found status code
func (r *response) NotFound(res any) error {
	return r.Send(http.StatusNotFound, res)
}

// Method for sending a response with a 500 Internal Server Error status code
func (r *response) InternalServerError(res any) error {
	return r.Send(http.StatusInternalServerError, res)
}

// Method for sending a response with a 409 Conflict status code
func (r *response) Conflict(res any) error { return r.Send(http.StatusConflict, res) }

// Method for sending a response with a 501 Not Implemented status code
func (r *response) NotImplemented() error { return r.Send(http.StatusNotImplemented, nil) }

// Method for sending a response with a 501 Not Implemented status code
func (r *response) SessionExpired() error { return r.Send(419, nil) }
