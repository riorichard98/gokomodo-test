package http

import (
	"fmt"
	"os"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"gokomodo-test/internal/interface/container"
	"gokomodo-test/internal/interface/server/http/middleware"
	"gokomodo-test/pkg/config"
)

func StartHttpService(container *container.Container) {
	e := echo.New() // Create a new Echo instance

	e.Validator = middleware.NewValidator() // Set up validator middleware for request body validation

	// Set the address to listen on for incoming HTTP requests based on the configuration provided in container
	e.Server.Addr = container.Config.Apps.HttpPort

	e.Use(middleware.LoggingMiddleware())        // Use logging middleware to log incoming requests
	e.Use(middleware.SetCorsConfig())            // Set up CORS middleware for handling Cross-Origin Resource Sharing (CORS) headers
	e.Use(middleware.JWTMiddleware(config.GetEnvString("JWT_SECRET_KEY")))
	e.HTTPErrorHandler = middleware.ErrorHandler // Set up custom error handler middleware
	SetupRouter(e, container)                    // Set up routes and handlers

	// Configure the Echo logger
	e.Logger.SetLevel(log.INFO)   // Set log level to INFO or DEBUG for more verbose output
	e.Logger.SetOutput(os.Stdout) // Set output writer to os.Stdout to output log messages to the terminal

	fmt.Printf("Starting..... App started on port %s\n", container.Config.Apps.HttpPort) // Log "Starting..." message with fmt.Printf
	gracehttp.Serve(e.Server)                                                            // Use gracehttp to gracefully serve HTTP requests                                                              // Use gracehttp to gracefully serve HTTP requests
}
