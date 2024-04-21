package httpHandler

import (
	"gokomodo-test/internal/interface/container"
)

type handler struct {
	OnboardHandler *onboardHandler
	SellerHandler  *sellerHandler
}

func SetupHandlers(container *container.Container) *handler {
	return &handler{
		OnboardHandler: NewOnboardHandler(container.OnboardService),
		SellerHandler:  NewSellerHandler(container.SellerService, container.OnboardService),
	}
}
