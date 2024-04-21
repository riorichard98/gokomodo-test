package http

import (
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo/v4"

	"gokomodo-test/internal/interface/container"
	"gokomodo-test/internal/interface/server/http/middleware"
)

func StartHttpService(container *container.Container) {
	e := echo.New() // Create a new Echo instance

	e.Validator = middleware.NewValidator() // Set up validator middleware for request body validation

	// Set the address to listen on for incoming HTTP requests based on the configuration provided in container
	e.Server.Addr = container.Config.Apps.HttpPort

	e.Use(middleware.LoggingMiddleware())        // Use logging middleware to log incoming requests
	e.Use(middleware.SetCorsConfig())            // Set up CORS middleware for handling Cross-Origin Resource Sharing (CORS) headers
	e.HTTPErrorHandler = middleware.ErrorHandler // Set up custom error handler middleware
	SetupRouter(e, container)                    // Set up routes and handlers

	e.Logger.Infof("Starting.....", "App started on port: "+container.Config.Apps.HttpPort) // Log "Starting..." message
	gracehttp.Serve(e.Server)                                                               // Use gracehttp to gracefully serve HTTP requests                                                              // Use gracehttp to gracefully serve HTTP requests
}
