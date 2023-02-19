package constants

const (
	// Success code
	StatusCodeSuccess = "1000"

	// Error that is caused by input
	StatusCodeGenericBadRequestError    = "2000"
	StatusCodeMissingRequiredParameters = "2100"
	StatusCodeInvalidParameters         = "2200"
	StatusCodeGenericNotFoundError      = "2300"
	StatusCodeUnprocessableEntity       = "2400"

	// Error that is caused by our own code
	StatusCodeGenericInternalError = "5000"
	StatusCodeDatabaseError        = "5100"
)
