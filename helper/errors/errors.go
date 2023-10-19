// Package errors holds custom error definition
package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Error implements the error interface
type Error interface {
	error
	Status() int
	Print() string
}

// CustomError is a custom error type
type CustomError struct {
	// StatusCode is the http status code
	StatusCode int `json:"-"`
	// Message is the error message to display
	Message string `json:"message,omitempty"`
	// Details is the error details for debugging
	Details interface{} `json:"details,omitempty"`
}

// ForbiddenError is mainly used for unauthorized access
// StatusCode is 403
func ForbiddenError(msg string) CustomError {
	return CustomError{
		StatusCode: http.StatusForbidden,
		Message:    msg,
	}
}

// InvalidRequestParsingError is used for invalid parsing of request body
// StatusCode is 400
func InvalidRequestParsingError(err error) CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid Request",
		Details:    err.Error(),
	}
}

// BadRequest is used to indicate that the request is invalid
// StatusCode is 400
func BadRequest(msg string) CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

// InternalDBError is used to indicate that an internal database error occurred
// StatusCode is 500
func InternalDBError(err error) CustomError {
	return CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Database Error",
		Details:    err.Error(),
	}
}

// NoEntityError indicates that the entity is not found
// StatusCode is 404
func NoEntityError(entity string) CustomError {
	parts := strings.Split(entity, "_")
	return CustomError{
		StatusCode: http.StatusNotFound,
		Message:    fmt.Sprintf("requested %s not found", strings.Join(parts, " ")),
	}
}

// ValidationError indicates that the request data is failed to validate
// StatusCode is 400
func ValidationError(msg string) CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

// DBMigrationError is used to indicate that an internal database migration error occurred
func DBMigrationError(err error) CustomError {
	return CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Database Migration Error",
		Details:    err.Error(),
	}
}

func InternalError(err error) CustomError {
	return CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Server Error",
		Details:    err.Error(),
	}
}

func ExternalRequestError(e error) CustomError {
	return CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "External Request Error",
		Details:    e.Error(),
	}
}

func CustomExternalError(status int, msg string, e interface{}) CustomError {
	// Do not set status code here
	return CustomError{
		StatusCode: status,
		Message:    msg,
		Details:    e,
	}
}

func CustomInternalError(msg string, e interface{}) CustomError {
	// Do not set status code here
	return CustomError{
		Message: msg,
		Details: e,
	}
}

// Status returns the status code
func (e CustomError) Status() int {
	return e.StatusCode
}

// Error returns the error message
func (e CustomError) Error() string {
	return e.Message
}

func (e CustomError) Print() string {
	b, _ := json.Marshal(e)
	return string(b)
}
