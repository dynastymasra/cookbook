package cookbook

import "errors"

const (
	ErrResourceNotExist = "the requested resource doesn't exists"
	ErrPrimaryDatabase  = "the database connection is failed"

	ErrEndpointNotFoundCode = 4404
	ErrMethodNotAllowedCode = 4405

	ErrDatabaseUnavailable = 5100
)

var (
	// ErrEndpointNotFound is a message error if endpoint url not found
	ErrEndpointNotFound = ErrorMessage{
		Code:  ErrEndpointNotFoundCode,
		Title: "Endpoint",
		Error: errors.New(ErrResourceNotExist),
	}

	// ErrMethodNotAllowed is a message error if http method is not allow to access the resource
	ErrMethodNotAllowed = ErrorMessage{
		Code:  ErrMethodNotAllowedCode,
		Title: "Method",
		Error: errors.New(ErrResourceNotExist),
	}

	ErrDatabaseConnectionFailed = ErrorMessage{
		Code:  ErrDatabaseUnavailable,
		Title: "Database",
		Error: errors.New(ErrPrimaryDatabase),
	}
)

// ErrorMessage format with error code and title
type ErrorMessage struct {
	Code  int
	Title string
	Error error
}
