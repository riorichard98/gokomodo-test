package response

const (
	// Success represents the HTTP status code for successful responses (2xx)
	CODE_SUCCESS    = "0200"
	MESSAGE_SUCCESS = "Success"

	// BadRequest represents the HTTP status code for bad request responses (400)
	CODE_BAD_REQUEST    = "0400"
	MESSAGE_BAD_REQUEST = "Bad Request"

	// Unauthenticated represents the HTTP status code for unauthenticated responses (401)
	CODE_UNAUTHENTICATED    = "0401"
	MESSAGE_UNAUTHENTICATED = "Unauthenticated"

	// Unauthorized represents the HTTP status code for unauthorized responses (403)
	CODE_UNAUTHORIZED    = "0403"
	MESSAGE_UNAUTHORIZED = "Unauthorized"

	// InternalServerError represents the HTTP status code for internal server error responses (500)
	CODE_INTERNAL_SERVER_ERROR = "0500"
	MESSAGE_INTERNAL_SERVER_ERROR = "Internal Server Error"
)
