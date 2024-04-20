package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware(secretKey string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:        &jwt.StandardClaims{},
		SigningKey:    []byte(secretKey),
		SigningMethod: "HS256",
		ErrorHandler: func(err error) error {
			if strings.Contains(err.Error(), "Token is expired") {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token is expired")
			}

			// Handle other JWT validation errors
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or missing token")
		},
	})
}
