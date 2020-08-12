package server

import (
	"encoding/json"
	"net/http"
)

type StatusProvider interface {
	StatusCode() int
}

type TypeProvider interface {
	Type() string
}

type ServerError struct {
	status  int
	kind    string
	message string
}

func (e ServerError) StatusCode() int {
	return e.status
}

func (e ServerError) Type() string {
	return e.kind
}

func (e ServerError) Error() string {
	return e.message
}

func Error(status int, kind, message string) error {
	return ServerError{
		kind:    kind,
		status:  status,
		message: message,
	}
}

func BadRequestError(message string) error {
	return Error(http.StatusBadRequest, "bad_request", message)
}

func InternalServerError(message string) error {
	return Error(http.StatusInternalServerError, "internal", message)
}

func NotFoundError(message string) error {
	return Error(http.StatusNotFound, "not_found", message)
}

func UnauthorizedError(message string) error {
	return Error(http.StatusUnauthorized, "unauthorized", message)
}

func ForbiddenError(message string) error {
	return Error(http.StatusForbidden, "forbidden", message)
}

func UnprocessableEntityError(message string) error {
	return Error(http.StatusUnprocessableEntity, "unprocessable", message)
}

type serverErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	if e, ok := err.(StatusProvider); ok {
		w.WriteHeader(e.StatusCode())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var body serverErrorResponse

	if e, ok := err.(TypeProvider); ok {
		body.Type = e.Type()
	} else {
		body.Type = "internal"
	}

	body.Message = err.Error()
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(body)
}
