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
