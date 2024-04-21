package httpHandler

import (
	"gokomodo-test/internal/interface/container"
)

type handler struct {
}

func SetupHandlers(container *container.Container) *handler {
	return &handler{
		// FileUploadHandler: NewUploadHandler(container.FileUploadService),
	}
}
