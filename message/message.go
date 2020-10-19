package message

type Message string

const (
	// ErrValueCannotEmptyOrNilM error message if value send by client is empty or null
	ErrValueCannotEmptyOrNilM Message = "the value cannot empty or null"
	// ErrValueNotValidUUIDM error message if value send by client is not valid UUID
	ErrValueNotValidUUIDM Message = "the value is not valid UUID"

	// ErrEndpointNotFoundM error message if endpoint requested by client or user not found
	ErrEndpointNotFoundM Message = "the requested endpoint doesn't exists"
	// ErrMethodNotAllowedCode error message if HTTP method not match
	ErrMethodNotAllowedM Message = "the http method doesn't match with existing"

	// ErrDatabaseUnavailableM
	ErrDatabaseUnavailableM Message = "the database connection is failed or unavailable"
	// ErrDatabaseDataNotFoundM error message if data requested by client or user not found in database
	ErrDatabaseDataNotFoundM Message = "the requested data not found"
	// ErrDatabaseDuplicateDataM error message if data inserted by client or user is duplicate with existing data
	ErrDatabaseDuplicateDataM Message = "the data has conflict with existing data"

	// ErrUnknownM message if error happen with unknown error code
	ErrUnknownM Message = "unknown error occur. try again later"
)
