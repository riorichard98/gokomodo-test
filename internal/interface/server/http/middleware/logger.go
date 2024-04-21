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
		logData := LogData{
			Timestamp:     time.Now().Format(time.RFC3339),
			RequestPath:   c.Request().URL.Path,
			RequestMethod: c.Request().Method,
			Request:       json.RawMessage(reqBody),
			Response:      json.RawMessage(resBody),
		}

		logJSON, err := json.Marshal(logData)
		if err != nil {
			// Handle JSON marshaling error.
			c.Logger().Errorf("Error marshaling log data: %v", err)
			return
		}

		// Print the JSON string or send it to a logging system.
		c.Logger().Info(string(logJSON))
	})
}
