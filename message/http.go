package message

import "net/http"

func HTTPStatusCode(code interface{}) int {
	switch code {
	case ErrValueCannotEmptyOrNil:
		return http.StatusBadRequest
	case ErrValueNotValidUUID:
		return http.StatusBadRequest
	case ErrEndpointNotFound:
		return http.StatusNotFound
	case ErrMethodNotAllowed:
		return http.StatusMethodNotAllowed
	case ErrRequestNotAcceptable:
		return http.StatusNotAcceptable
	case ErrUnsupportedMediaType:
		return http.StatusUnsupportedMediaType
	case ErrDatabaseUnavailable:
		return http.StatusServiceUnavailable
	case ErrDatabaseDataNotFound:
		return http.StatusNotFound
	case ErrDatabaseDuplicate:
		return http.StatusConflict
	case ErrDatabaseDataExpectation:
		return http.StatusPreconditionFailed
	default:
		return http.StatusNotImplemented
	}
}
