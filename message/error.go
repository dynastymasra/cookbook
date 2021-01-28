package message

// FailedMessage format with title, message, and Code
type FailedMessage struct {
	Title   string
	Message interface{}
	Code    interface{}
}
