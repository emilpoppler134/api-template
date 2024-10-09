package http

import (
	"encoding/json"
)

type Response struct {
	StatusCode int
	Body       []byte
}

const (
	StatusOk             = 200
	StatusCreated        = 201
	StatusAccepted       = 202
	StatusNoContent      = 204
	StatusPartialContent = 206

	StatusBadRequest           = 400
	StatusUnauthorized         = 401
	StatusForbidden            = 403
	StatusNotFound             = 404
	StatusMethodNotAllowed     = 405
	StatusConflict             = 409
	StatusPreconditionRequired = 428

	StatusInternalServerError = 500
)

// Constructors
func ConstructResponse(statusCode int, payload interface{}) Response {
	data, err := json.Marshal(payload)
	if err != nil {
		return ConstructError(StatusInternalServerError, "Something went wrong")
	}
	return Response{statusCode, data}
}

func ConstructError(statusCode int, message string) Response {
	payload := map[string]string{"error": message}
	return ConstructResponse(statusCode, payload)
}

// Success responses
func Ok(payload interface{}) Response {
	return ConstructResponse(StatusOk, payload)
}

func Accepted(payload interface{}) Response {
	return ConstructResponse(StatusOk, payload)
}

func Created(payload interface{}) Response {
	return ConstructResponse(StatusCreated, payload)
}

func NoContent(payload interface{}) Response {
	return ConstructResponse(StatusNoContent, payload)
}

func PartialContent(payload interface{}) Response {
	return ConstructResponse(StatusPartialContent, payload)
}

// Error responses
func BadRequest(message string) Response {
	return ConstructError(StatusBadRequest, message)
}

func Unauthorized(message string) Response {
	return ConstructError(StatusUnauthorized, message)
}

func Forbidden(message string) Response {
	return ConstructError(StatusForbidden, message)
}

func NotFound(message string) Response {
	return ConstructError(StatusNotFound, message)
}

func Conflict(message string) Response {
	return ConstructError(StatusConflict, message)
}

func PreconditionRequired(message string) Response {
	return ConstructError(StatusPreconditionRequired, message)
}

func InternalServerError() Response {
	return ConstructError(StatusInternalServerError, "Something went wrong")
}
