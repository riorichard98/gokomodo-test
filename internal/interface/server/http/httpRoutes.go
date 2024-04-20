package http

import (
	"net/http"

	"gokomodo-test/internal/interface/container"
	"gokomodo-test/pkg/config"
	// "gokomodo-test/internal/interface/server/http/httpHandler"

	"gokomodo-test/internal/interface/server/http/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRouter(server *echo.Echo, cont *container.Container) {
	// handler := httpHandler.SetupHandlers(cont)

	// - health check
	server.GET("/ping", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "services up and running... ")
	})

	// onboardRoute := server.Group("/onboard")

	server.Use(middleware.JWTMiddleware(config.GetEnvString("JWT_SECRET_KEY")))

	// buyerRoute := server.Group("/buyer")
	// sellerRoute := server.Group("/seller")

}
