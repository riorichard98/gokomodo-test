package middleware

import (
    "net/http"

    "github.com/go-playground/validator/v10"
    "github.com/labstack/echo/v4"
)

// CustomValidator holds a validator instance
type CustomValidator struct {
    validator *validator.Validate
}

// NewValidator creates a new instance of CustomValidator
func NewValidator() *CustomValidator {
    return &CustomValidator{
        validator: validator.New(),
    }
}

// Validate validates the request body using the validator instance
func (cv *CustomValidator) Validate(i interface{}) error {
    if err := cv.validator.Struct(i); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return nil
}
