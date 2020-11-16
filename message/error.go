package message

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dynastymasra/cookbook"

	"github.com/go-playground/validator/v10"
)

type InternalError struct {
	Code string
	Msg  error
}

func NewInternalError(err error) *InternalError {
	return &InternalError{
		Msg: err,
	}
}

func (i *InternalError) Error() string {
	return i.Msg.Error()
}

func IsInternalError(err error) bool {
	if _, ok := err.(*InternalError); ok {
		return true
	}
	return false
}

// ErrorMessage format with error code and title
type ErrorMessage struct {
	Code  Code   `json:"code,omitempty"`
	Title string `json:"title,omitempty"`
	Error error  `json:"message,omitempty"`
}

// NewErrorMessage create new error message format
func NewErrorMessage(code Code, title string, err error) *ErrorMessage {
	return &ErrorMessage{
		Code:  code,
		Title: title,
		Error: err,
	}
}

func (e Code) ErrorMessage() *ErrorMessage {
	switch e {
	case ErrValueCannotEmptyOrNilCode:
		return NewErrorMessage(ErrValueCannotEmptyOrNilCode, "empty or null", fmt.Errorf("%v", ErrValueCannotEmptyOrNilM))
	case ErrValueNotValidUUIDCode:
		return NewErrorMessage(ErrValueNotValidUUIDCode, "uuid", fmt.Errorf("%v", ErrValueNotValidUUIDM))
	case ErrEndpointNotFoundCode:
		return NewErrorMessage(ErrEndpointNotFoundCode, "not found", fmt.Errorf("%v", ErrEndpointNotFoundM))
	case ErrMethodNotAllowedCode:
		return NewErrorMessage(ErrMethodNotAllowedCode, "method", fmt.Errorf("%v", ErrMethodNotAllowedM))
	case ErrRequestNotAcceptableCode:
		return NewErrorMessage(ErrRequestNotAcceptableCode, "request", fmt.Errorf("%v", ErrRequestNotAcceptableM))
	case ErrUnsupportedMediaTypeCode:
		return NewErrorMessage(ErrUnsupportedMediaTypeCode, "media type", fmt.Errorf("%v", ErrUnsupportedMediaTypeM))
	case ErrDatabaseUnavailableCode:
		return NewErrorMessage(ErrDatabaseUnavailableCode, "database", fmt.Errorf("%v", ErrDatabaseUnavailableM))
	case ErrDatabaseDataNotFoundCode:
		return NewErrorMessage(ErrDatabaseDataNotFoundCode, "not found", fmt.Errorf("%v", ErrDatabaseDataNotFoundM))
	case ErrDatabaseDuplicateCode:
		return NewErrorMessage(ErrDatabaseDuplicateCode, "conflict", fmt.Errorf("%v", ErrDatabaseDuplicateDataM))
	case ErrDatabaseDataExpectationCode:
		return NewErrorMessage(ErrDatabaseDataExpectationCode, "data", fmt.Errorf("%v", ErrDatabaseDataExpectationM))
	default:
		return NewErrorMessage(ErrUnknownCode, "unknown", fmt.Errorf("%v", ErrUnknownM))
	}
}

func (e Code) HTTPErrorMessage() int {
	switch e {
	case ErrValueCannotEmptyOrNilCode:
		return http.StatusBadRequest
	case ErrValueNotValidUUIDCode:
		return http.StatusBadRequest
	case ErrEndpointNotFoundCode:
		return http.StatusNotFound
	case ErrMethodNotAllowedCode:
		return http.StatusMethodNotAllowed
	case ErrRequestNotAcceptableCode:
		return http.StatusNotAcceptable
	case ErrUnsupportedMediaTypeCode:
		return http.StatusUnsupportedMediaType
	case ErrDatabaseUnavailableCode:
		return http.StatusServiceUnavailable
	case ErrDatabaseDataNotFoundCode:
		return http.StatusNotFound
	case ErrDatabaseDuplicateCode:
		return http.StatusConflict
	case ErrDatabaseDataExpectationCode:
		return http.StatusPreconditionFailed
	default:
		return http.StatusNotImplemented
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
	var code Code
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
		code = ErrRequestNotAcceptableCode
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

func ErrorMessageToJSONList(msg []ErrorMessage) []cookbook.JSON {
	var res []cookbook.JSON

	for _, v := range msg {
		res = append(res, cookbook.JSON{
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
	var code Code
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

func ParseValidator(err error) []cookbook.JSON {
	var res []cookbook.JSON

	switch e := err.(type) {
	case validator.ValidationErrors:
		for _, ve := range e {
			field := strings.ToLower(ve.Field())
			res = append(res, cookbook.JSON{
				"code":    ErrInvalidValueCode,
				"title":   field,
				"message": fmt.Sprintf("Error field validation for '%s' failed on the '%s' tag", field, ve.Tag()),
			})
		}
	default:
		res = append(res, cookbook.JSON{
			"code":    ErrUnknownCode,
			"title":   "Unknown",
			"message": err.Error(),
		})
	}

	return res
}

// HTTPDataNotFound return new client error message
func HTTPDataNotFound(title string, err error) *ClientError {
	return NewClientError(http.StatusNotFound, *NewErrorMessage(ErrDatabaseDataNotFoundCode, title, err))
}

// HTTPDataDuplicate return new client error message
func HTTPDataDuplicate(title string, err error) *ClientError {
	return NewClientError(http.StatusConflict, *NewErrorMessage(ErrDatabaseDuplicateCode, title, err))
}

// HTTPDataExpectationFailed return new server error message
func HTTPDataExpectationFailed(title string, err error) *ClientError {
	return NewClientError(http.StatusPreconditionFailed, *NewErrorMessage(ErrDatabaseDataExpectationCode, title, err))
}

// HTTPDatabaseUnavailable return new server error message
func HTTPDatabaseUnavailable(title string, err error) *ServerError {
	return NewServerError(http.StatusServiceUnavailable, *NewErrorMessage(ErrDatabaseUnavailableCode, title, err))
}
