package cookbook

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	ErrResourceNotExist  = "the requested resource doesn't exists"
	ErrPrimaryDatabase   = "the database connection is failed"
	ErrErrDuplicateDataM = "the data has duplicated"

	ErrEndpointNotFoundCode = 40404
	ErrMethodNotAllowedCode = 40405
	ErrReadRequestBodyCode  = 40406
	ErrDuplicateDataCode    = 40409

	ErrInternalServiceCode     = 50000
	ErrDatabaseUnavailableCode = 50100
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

	// ErrDatabaseConnectionFailed is a message to inform if database connection failed or refuse
	ErrDatabaseConnectionFailed = ErrorMessage{
		Code:  ErrDatabaseUnavailableCode,
		Title: "Database",
		Error: errors.New(ErrPrimaryDatabase),
	}

	// ErrErrDuplicateData error message if the data inserted to database has duplicate
	ErrErrDuplicateData = ErrorMessage{
		Code:  ErrDuplicateDataCode,
		Title: "Duplicate",
		Error: errors.New(ErrErrDuplicateDataM),
	}
)

// ErrorMessage format with error code and title
type ErrorMessage struct {
	Code  int
	Title string
	Error error
}

func NewErrorMessage(code int, title string, err error) *ErrorMessage {
	return &ErrorMessage{
		Code:  code,
		Title: title,
		Error: err,
	}
}

type ClientError struct {
	HTTPCode int
	Message  []ErrorMessage
}

func (c *ClientError) Error() string {
	return fmt.Sprintf("%+v", c.Message)
}

func NewClientError(code int, msg ...ErrorMessage) *ClientError {
	return &ClientError{
		HTTPCode: code,
		Message:  msg,
	}
}

func FromClientError(err error) *ClientError {
	if msg, ok := err.(*ClientError); ok {
		return msg
	}

	return &ClientError{
		HTTPCode: http.StatusTeapot,
		Message: []ErrorMessage{
			{
				Code:  ErrInternalServiceCode,
				Title: "Unknown",
				Error: err,
			},
		},
	}
}

type ServerError struct {
	HTTPCode int
	Message  ErrorMessage
}

func (s *ServerError) Error() string {
	return s.Message.Error.Error()
}

func NewServerError(code int, msg ErrorMessage) *ServerError {
	return &ServerError{
		HTTPCode: code,
		Message:  msg,
	}
}

func FromServerError(err error) *ServerError {
	if msg, ok := err.(*ServerError); ok {
		return msg
	}

	return &ServerError{
		HTTPCode: http.StatusInternalServerError,
		Message: ErrorMessage{
			Code:  ErrInternalServiceCode,
			Title: "Unknown",
			Error: err,
		},
	}
}
