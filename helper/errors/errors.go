// Package errors holds custom error definition
package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error implements the error interface
type Error interface {
	error
	Status() int
	Print() string
}

// CustomError is a custom error type
type CustomError struct {
	// statusCode is the http status code
	statusCode int
	// Message is the error message to display
	Message string `json:"message,omitempty"`
	// Details is the error details for debugging
	Details interface{} `json:"details,omitempty"`
}

// ForbiddenError is mainly used for unauthorized access
// statusCode is 403
func ForbiddenError(msg string) CustomError {
	return CustomError{
		statusCode: http.StatusForbidden,
		Message:    msg,
	}
}

// InvalidRequestParsingError is used for invalid parsing of request body
// statusCode is 400
func InvalidRequestParsingError(err error) CustomError {
	return CustomError{
		statusCode: http.StatusBadRequest,
		Message:    "Invalid request",
		Details:    err.Error(),
	}
}

// InternalDBError is used to indicate that an internal database error occurred
// statusCode is 500
func InternalDBError(err error) CustomError {
	return CustomError{
		statusCode: http.StatusInternalServerError,
		Message:    "Internal database error",
		Details:    err.Error(),
	}
}

// NoEntityError indicates that the entity is not found
// statusCode is 404
func NoEntityError(entity string) CustomError {
	return CustomError{
		statusCode: http.StatusNotFound,
		Message:    fmt.Sprintf("Requested %s not found", entity),
	}
}

// ValidationError indicates that the request data is failed to validate
// statusCode is 400
func ValidationError(msg string) CustomError {
	return CustomError{
		statusCode: http.StatusBadRequest,
		Message:    "Invalid input data",
		Details:    msg,
	}
}

// DBMigrationError is used to indicate that an internal database migration error occurred
func DBMigrationError(err error) CustomError {
	return CustomError{
		statusCode: http.StatusInternalServerError,
		Message:    "Database migration error",
		Details:    err.Error(),
	}
}

func InternalError(err error) CustomError {
	return CustomError{
		statusCode: http.StatusInternalServerError,
		Message:    "Internal server error",
		Details:    err.Error(),
	}
}

// Status returns the status code
func (e CustomError) Status() int {
	return e.statusCode
}

// Error returns the error message
func (e CustomError) Error() string {
	return e.Message
}

func (e CustomError) Print() string {
	b, _ := json.Marshal(e)
	return string(b)
}
