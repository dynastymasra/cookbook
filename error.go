package cookbook

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	ErrResourceNotExistM = "the requested resource doesn't exists"

	ErrDatabaseUnavailableM   = "the database connection is failed"
	ErrDatabaseDuplicateDataM = "the data has conflict with existing data"
	ErrDatabaseDataNotFoundM  = "the data requested not found"

	ErrDatabaseUnavailableCode  = 30000
	ErrDatabaseDataNotFoundCode = 30001
	ErrDatabaseDuplicateCode    = 30002

	ErrInvalidValueCode         = 40400
	ErrUnauthorizedCode         = 40401
	ErrPaymentRequiredCode      = 40402
	ErrForbiddenCode            = 40403
	ErrEndpointNotFoundCode     = 40404
	ErrMethodNotAllowedCode     = 40405
	ErrReadRequestBodyCode      = 40406
	ErrProxyAuthRequiredCode    = 40407
	ErrRequestTimeoutCode       = 40408
	ErrConflictCode             = 40409
	ErrDataMissingCode          = 40410
	ErrLengthRequiredCode       = 40411
	ErrPreconditionCode         = 40412
	ErrDataToLargeCode          = 40413
	ErrURITooLongCode           = 40414
	ErrUnsupportedMediaTypeCode = 40415
	ErrRangeTooLongCode         = 40416
	ErrExpectationCode          = 40417
	ErrMisdirectedRequestCode   = 40421
	ErrUnprocessableEntityCode  = 40422
	ErrDataLockedCode           = 40423
	ErrFailedDependencyCode     = 40424
	ErrTooEarlyCode             = 40425
	ErrUpgradeRequiredCode      = 40426
	ErrPreconditionRequiredCode = 40428
	ErrTooManyRequestsCode      = 40429
	ErrHeaderTooLargeCode       = 40431

	ErrInternalServiceCode       = 50500
	ErrNotImplementedCode        = 50501
	ErrBadGatewayCode            = 50502
	ErrServiceUnavailableCode    = 50503
	ErrGatewayCode               = 50504
	ErrNotSupportedCode          = 50505
	ErrVariantCode               = 50506
	ErrInsufficientStorageCode   = 50507
	ErrLoopCode                  = 50508
	ErrNotExtendedCode           = 50510
	ErrNetworkAuthenticationCode = 50511

	ErrUnknownCode = 99999
)

var (
	// ErrEndpointNotFound is a message error if endpoint url not found
	ErrEndpointNotFound = ErrorMessage{
		Code:  ErrEndpointNotFoundCode,
		Title: "Endpoint",
		Error: errors.New(ErrResourceNotExistM),
	}

	// ErrMethodNotAllowed is a message error if http method is not allow to access the resource
	ErrMethodNotAllowed = ErrorMessage{
		Code:  ErrMethodNotAllowedCode,
		Title: "Method",
		Error: errors.New(ErrResourceNotExistM),
	}

	// ErrDatabaseUnavailable a message to inform if database connection failed or refuse
	ErrDatabaseUnavailable = ErrorMessage{
		Code:  ErrDatabaseUnavailableCode,
		Title: "Database",
		Error: errors.New(ErrDatabaseUnavailableM),
	}

	// ErrDatabaseDataNotFound error message if data from database doesn't exist
	ErrDatabaseDataNotFound = ErrorMessage{
		Code:  ErrDatabaseDataNotFoundCode,
		Title: "Not Found",
		Error: errors.New(ErrDatabaseDataNotFoundM),
	}

	// ErrDatabaseDuplicate error message if the data inserted to database has duplicate
	ErrDatabaseDuplicate = ErrorMessage{
		Code:  ErrDatabaseDuplicateCode,
		Title: "Duplicate",
		Error: errors.New(ErrDatabaseDuplicateDataM),
	}

	ErrorResult = map[int]ErrorMessage{
		ErrEndpointNotFoundCode: ErrEndpointNotFound,
		ErrMethodNotAllowedCode: ErrMethodNotAllowed,

		ErrDatabaseUnavailableCode:  ErrDatabaseUnavailable,
		ErrDatabaseDataNotFoundCode: ErrDatabaseDataNotFound,
		ErrDatabaseDuplicateCode:    ErrDatabaseDuplicate,
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

func HTTPToClientError(status int, title, body string) *ClientError {
	var code int
	switch status {
	case http.StatusBadRequest:
		code = ErrInvalidValueCode
	case http.StatusUnauthorized:
		code = ErrUnauthorizedCode
	case http.StatusPaymentRequired:
		code = ErrPaymentRequiredCode
	case http.StatusForbidden:
		code = ErrForbiddenCode
	case http.StatusNotFound:
		code = ErrEndpointNotFoundCode
	case http.StatusMethodNotAllowed:
		code = ErrMethodNotAllowedCode
	case http.StatusNotAcceptable:
		code = ErrReadRequestBodyCode
	case http.StatusProxyAuthRequired:
		code = ErrProxyAuthRequiredCode
	case http.StatusRequestTimeout:
		code = ErrRequestTimeoutCode
	case http.StatusConflict:
		code = ErrConflictCode
	case http.StatusGone:
		code = ErrDataMissingCode
	case http.StatusLengthRequired:
		code = ErrLengthRequiredCode
	case http.StatusPreconditionFailed:
		code = ErrPreconditionCode
	case http.StatusRequestEntityTooLarge:
		code = ErrDataToLargeCode
	case http.StatusRequestURITooLong:
		code = ErrURITooLongCode
	case http.StatusUnsupportedMediaType:
		code = ErrUnsupportedMediaTypeCode
	case http.StatusRequestedRangeNotSatisfiable:
		code = ErrRangeTooLongCode
	case http.StatusExpectationFailed:
		code = ErrExpectationCode
	case http.StatusMisdirectedRequest:
		code = ErrMisdirectedRequestCode
	case http.StatusUnprocessableEntity:
		code = ErrUnprocessableEntityCode
	case http.StatusLocked:
		code = ErrDataLockedCode
	case http.StatusFailedDependency:
		code = ErrFailedDependencyCode
	case http.StatusTooEarly:
		code = ErrTooEarlyCode
	case http.StatusUpgradeRequired:
		code = ErrUpgradeRequiredCode
	case http.StatusPreconditionRequired:
		code = ErrPreconditionRequiredCode
	case http.StatusTooManyRequests:
		code = ErrTooManyRequestsCode
	case http.StatusRequestHeaderFieldsTooLarge:
		code = ErrHeaderTooLargeCode
	default:
		code = ErrUnknownCode
	}

	return &ClientError{
		HTTPCode: status,
		Message: []ErrorMessage{
			{
				Code:  code,
				Title: title,
				Error: errors.New(body),
			},
		},
	}
}

func ErrorMessageToJSONList(msg []ErrorMessage) []JSON {
	var res []JSON

	for _, v := range msg {
		res = append(res, JSON{
			"code":    v.Code,
			"title":   v.Title,
			"message": v.Error.Error(),
		})
	}

	return res
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

func HTTPtoServerError(status int, title, body string) *ServerError {
	var code int
	switch status {
	case http.StatusInternalServerError:
		code = ErrInternalServiceCode
	case http.StatusNotImplemented:
		code = ErrNotImplementedCode
	case http.StatusBadGateway:
		code = ErrBadGatewayCode
	case http.StatusServiceUnavailable:
		code = ErrServiceUnavailableCode
	case http.StatusGatewayTimeout:
		code = ErrGatewayCode
	case http.StatusHTTPVersionNotSupported:
		code = ErrNotSupportedCode
	case http.StatusVariantAlsoNegotiates:
		code = ErrVariantCode
	case http.StatusInsufficientStorage:
		code = ErrInsufficientStorageCode
	case http.StatusLoopDetected:
		code = ErrLoopCode
	case http.StatusNotExtended:
		code = ErrNotExtendedCode
	case http.StatusNetworkAuthenticationRequired:
		code = ErrNetworkAuthenticationCode
	default:
		code = ErrUnknownCode
	}

	return &ServerError{
		HTTPCode: status,
		Message: ErrorMessage{
			Code:  code,
			Title: title,
			Error: errors.New(body),
		},
	}
}

func ParseValidator(err error) []JSON {
	var res []JSON

	switch e := err.(type) {
	case validator.ValidationErrors:
		for _, ve := range e {
			field := strings.ToLower(ve.Field())
			res = append(res, JSON{
				"code":    ErrInvalidValueCode,
				"title":   field,
				"message": fmt.Sprintf("Error field validation for '%s' failed on the '%s' tag", field, ve.Tag()),
			})
		}
	default:
		res = append(res, JSON{
			"code":    ErrUnknownCode,
			"title":   "Unknown",
			"message": err.Error(),
		})
	}

	return res
}
