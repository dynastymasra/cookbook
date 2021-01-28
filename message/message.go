package message

const (
	// Error message related with request value
	// ErrValueCannotEmptyOrNilM error  if value send by client is empty or null
	ErrValueCannotEmptyOrNilM = "the value cannot empty or null"
	// ErrValueNotValidUUIDM error  if value send by client is not valid UUID
	ErrValueNotValidUUIDM = "the value is not valid UUID"

	// Error message related with HTTP request
	// ErrEndpointNotFoundM error  if endpoint requested by client or user not found
	ErrEndpointNotFoundM = "the requested endpoint doesn't exists"
	// ErrMethodNotAllowedCode error  if HTTP method not match
	ErrMethodNotAllowedM = "the http method doesn't match with existing"
	// ErrRequestNotAcceptableM error  if HTTP request header not acceptable
	ErrRequestNotAcceptableM = "the http header not acceptable"
	// ErrUnsupportedMediaTypeM error  if HTTP request header with unsupported media type
	ErrUnsupportedMediaTypeM = "unsupported media type"

	// Error message related with database
	// ErrDatabaseUnavailableM
	ErrDatabaseUnavailableM = "the database connection is failed or unavailable"
	// ErrDatabaseDataNotFoundM error  if data requested by client or user not found in database
	ErrDatabaseDataNotFoundM = "the requested data not found"
	// ErrDatabaseDuplicateDataM error  if data inserted by client or user is duplicate with existing data
	ErrDatabaseDuplicateDataM = "the data has conflict with existing data"
	// ErrDatabaseDataExpectationM error  if result data from database doesn't meet the expected result
	ErrDatabaseDataExpectationM = "the data requested doesn't meet the expected result"

	// ErrUnknownM  if error happen with unknown error code
	ErrUnknownM = "unknown error occur. try again later"
)
