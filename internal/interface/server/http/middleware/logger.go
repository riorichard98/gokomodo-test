package middleware

import (
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// LogData is the structure that holds the log information.
type LogData struct {
	Timestamp     string          `json:"timestamp"`
	RequestPath   string          `json:"request_path"`
	RequestMethod string          `json:"request_method"`
	Request       json.RawMessage `json:"request_body"`
	Response      json.RawMessage `json:"response_body"`
}

// LoggingMiddleware creates a middleware to log request and response bodies.
func LoggingMiddleware() echo.MiddlewareFunc {
	return middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		var rawReqBody, rawResBody json.RawMessage
		var err error

		// Prepare the request body for logging
		if len(reqBody) > 0 {
			rawReqBody = json.RawMessage(reqBody)
		} else {
			rawReqBody = json.RawMessage(`{}`)
		}

		// Prepare the response body for logging
		if len(resBody) > 0 {
			rawResBody = json.RawMessage(resBody)
		} else {
			rawResBody = json.RawMessage(`{}`)
		}

		// Prepare log data
		logData := LogData{
			Timestamp:     time.Now().Format(time.RFC3339),
			RequestPath:   c.Request().URL.Path,
			RequestMethod: c.Request().Method,
			Request:       rawReqBody,
			Response:      rawResBody,
		}

		// Marshal the log data into JSON
		logJSON, err := json.Marshal(logData)
		if err != nil {
			c.Logger().Errorf("Error marshaling log data: %v", err)
			return
		}

		// Log the JSON string
		c.Logger().Infof(string(logJSON))
	})
}
