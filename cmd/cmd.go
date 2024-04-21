package cmd

import (
	"gokomodo-test/internal/interface/container"
	"gokomodo-test/internal/interface/server"
)

func Run() {
	container := container.New()
	server.StartService(container)
}
