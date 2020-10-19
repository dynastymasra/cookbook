package message

type Code int

const (
	// ErrValueCannotEmptyOrNilCode error code if value send by client is empty or null
	ErrValueCannotEmptyOrNilCode Code = 10000
	// ErrValueNotValidUUIDCode error code if value send by client is not valid UUID
	ErrValueNotValidUUIDCode Code = 10001

	// ErrEndpointNotFoundCode error code if endpoint is not found
	ErrEndpointNotFoundCode Code = 20404
	// ErrMethodNotAllowedCode error code if http method not match
	ErrMethodNotAllowedCode Code = 20405

	// ErrDatabaseUnavailableCode error code if database is unavailable
	ErrDatabaseUnavailableCode Code = 30000
	// ErrDatabaseDataNotFoundCode error code if data requested by user or client not found
	ErrDatabaseDataNotFoundCode Code = 30001
	// ErrDatabaseDuplicateCode error code if data inserted by user or client is duplicate with existing data
	ErrDatabaseDuplicateCode Code = 30002

	ErrInvalidValueCode         Code = 40400
	ErrUnauthorizedCode         Code = 40401
	ErrPaymentRequiredCode      Code = 40402
	ErrForbiddenCode            Code = 40403
	ErrReadRequestBodyCode      Code = 40406
	ErrProxyAuthRequiredCode    Code = 40407
	ErrRequestTimeoutCode       Code = 40408
	ErrConflictCode             Code = 40409
	ErrDataMissingCode          Code = 40410
	ErrLengthRequiredCode       Code = 40411
	ErrPreconditionCode         Code = 40412
	ErrDataToLargeCode          Code = 40413
	ErrURITooLongCode           Code = 40414
	ErrUnsupportedMediaTypeCode Code = 40415
	ErrRangeTooLongCode         Code = 40416
	ErrExpectationCode          Code = 40417
	ErrMisdirectedRequestCode   Code = 40421
	ErrUnprocessableEntityCode  Code = 40422
	ErrDataLockedCode           Code = 40423
	ErrFailedDependencyCode     Code = 40424
	ErrTooEarlyCode             Code = 40425
	ErrUpgradeRequiredCode      Code = 40426
	ErrPreconditionRequiredCode Code = 40428
	ErrTooManyRequestsCode      Code = 40429
	ErrHeaderTooLargeCode       Code = 40431

	ErrInternalServiceCode       Code = 50500
	ErrNotImplementedCode        Code = 50501
	ErrBadGatewayCode            Code = 50502
	ErrServiceUnavailableCode    Code = 50503
	ErrGatewayCode               Code = 50504
	ErrNotSupportedCode          Code = 50505
	ErrVariantCode               Code = 50506
	ErrInsufficientStorageCode   Code = 50507
	ErrLoopCode                  Code = 50508
	ErrNotExtendedCode           Code = 50510
	ErrNetworkAuthenticationCode Code = 50511

	// ErrUnknownCode code if error happen with unknown mapping error code
	ErrUnknownCode = 99999
)
