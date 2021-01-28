package message

const (
	// Error code related with request value
	// ErrValueCannotEmptyOrNil error  if value send by client is empty or null
	ErrValueCannotEmptyOrNil = "10000"
	// ErrValueNotValidUUID error  if value send by client is not valid UUID
	ErrValueNotValidUUID = "10001"

	// Error code related with database
	// ErrDatabaseUnavailable error if database is unavailable
	ErrDatabaseUnavailable = "30000"
	// ErrDatabaseDataNotFound error if data requested by user or client not found
	ErrDatabaseDataNotFound = "30001"
	// ErrDatabaseDuplicate error if data inserted by user or client is duplicate with existing data
	ErrDatabaseDuplicate = "30002"
	// ErrDatabaseDataExpectation error if result data from database doesn't meet the expectation
	ErrDatabaseDataExpectation = "30003"

	// Error code related with HTTP request
	ErrInvalidValue    = "40400"
	ErrUnauthorized    = "40401"
	ErrPaymentRequired = "40402"
	ErrForbidden       = "40403"
	// ErrEndpointNotFound error if endpoint is not found
	ErrEndpointNotFound = "40404"
	// ErrMethodNotAllowed error if http method not match
	ErrMethodNotAllowed = "40405"
	// ErrRequestNotAcceptable error if http header not acceptable
	ErrRequestNotAcceptable = "40406"
	ErrProxyAuthRequired    = "40407"
	ErrRequestTimeout       = "40408"
	ErrConflict             = "40409"
	ErrDataMissing          = "40410"
	ErrLengthRequired       = "40411"
	ErrPrecondition         = "40412"
	ErrDataToLarge          = "40413"
	ErrURITooLong           = "40414"
	// ErrUnsupportedMediaType error if http request unsupported media type
	ErrUnsupportedMediaType = "40415"
	ErrRangeTooLong         = "40416"
	ErrExpectation          = "40417"
	ErrMisdirectedRequest   = "40421"
	ErrUnprocessableEntity  = "40422"
	ErrDataLocked           = "40423"
	ErrFailedDependency     = "40424"
	ErrTooEarly             = "40425"
	ErrUpgradeRequired      = "40426"
	ErrPreconditionRequired = "40428"
	ErrTooManyRequests      = "40429"
	ErrHeaderTooLarge       = "40431"
	// ErrInternalServer error if server error or panic, cannot process the request
	ErrInternalServer        = "40500"
	ErrNotImplemented        = "40501"
	ErrBadGateway            = "40502"
	ErrServiceUnavailable    = "40503"
	ErrGateway               = "40504"
	ErrNotSupported          = "40505"
	ErrVariant               = "40506"
	ErrInsufficientStorage   = "40507"
	ErrLoop                  = "40508"
	ErrNotExtended           = "40510"
	ErrNetworkAuthentication = "40511"

	// ErrUnknown if error happen with unknown mapping error
	ErrUnknown = "99999"
)
