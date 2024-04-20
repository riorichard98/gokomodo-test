package middleware

import (
	"fmt"
	"gokomodo-test/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ErrorHandler handles errors and sends a JSON response
func ErrorHandler(err error, c echo.Context) {
	// Log the error
	c.Logger().Error(err)

	// Check if the response has already been sent
	if c.Response().Committed {
		return
	}

	// Default internal server error message
	status := response.CODE_INTERNAL_SERVER_ERROR
	message := response.MESSAGE_INTERNAL_SERVER_ERROR

	// Handle specific types of HTTP errors
	if he, ok := err.(*echo.HTTPError); ok {
		message = fmt.Sprintf("%v", he.Message)
	}

	// Prepare the error response
	errorResponse := response.ErrorResponse(status,message)

	// Send JSON response
	c.JSON(http.StatusInternalServerError, errorResponse)
}
