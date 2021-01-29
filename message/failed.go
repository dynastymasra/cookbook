package message

// FailedMessage format with title, message, and Code
type FailedMessage struct {
	Code    interface{} `json:"code"`
	Title   string      `json:"title"`
	Message interface{} `json:"message"`
}

// FailedUnsupportedMediaType is failed message if client request with unsupported media type
func FailedUnsupportedMediaType() FailedMessage {
	return FailedMessage{
		Title:   "HEADER_MEDIA_TYPE",
		Message: ErrUnsupportedMediaTypeM,
		Code:    ErrUnsupportedMediaType,
	}
}

// FailedRequestNotAcceptable is failed message if client http header not acceptable
func FailedRequestNotAcceptable() FailedMessage {
	return FailedMessage{
		Title:   "REQUEST_NOT_ACCEPTABLE",
		Message: ErrRequestNotAcceptableM,
		Code:    ErrRequestNotAcceptable,
	}
}

// FailedEndpointNotFound is failed message if client request to endpoint not exist
func FailedEndpointNotFound() FailedMessage {
	return FailedMessage{
		Title:   "NOT_FOUND",
		Message: ErrEndpointNotFoundM,
		Code:    ErrEndpointNotFound,
	}
}

// FailedMethodNotAllowed is failed message if client request with not allowed method
func FailedMethodNotAllowed() FailedMessage {
	return FailedMessage{
		Title:   "METHOD_NOT_ALLOWED",
		Message: ErrMethodNotAllowedM,
		Code:    ErrMethodNotAllowed,
	}
}

// FailedUnauthorized is failed message if client is unauthenticated
func FailedUnauthorized() FailedMessage {
	return FailedMessage{
		Title:   "UNAUTHORIZED",
		Message: ErrUnauthorizedM,
		Code:    ErrUnauthorized,
	}
}

// FailedForbidden is failed message if the client does not have access rights to the content
func FailedForbidden() FailedMessage {
	return FailedMessage{
		Title:   "FORBIDDEN",
		Message: ErrForbiddenM,
		Code:    ErrForbidden,
	}
}
