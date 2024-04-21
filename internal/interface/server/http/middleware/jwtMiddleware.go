package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware(secretKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Check if the request path starts with "/onboard" or is "/ping"
			if strings.HasPrefix(c.Request().URL.Path, "/onboard") || c.Request().URL.Path == "/ping" {
				// Allow access to endpoints starting with "/onboard" and the "/ping" endpoint without requiring a token
				return next(c)
			}

			// For other endpoints, perform JWT authentication
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
			})(next)(c)
		}
	}
}
