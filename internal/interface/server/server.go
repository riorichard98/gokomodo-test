package server

import (
	"gokomodo-test/internal/interface/container"
	"gokomodo-test/internal/interface/server/http"
)

func StartService(cont *container.Container) {
	http.StartHttpService(cont)
}
