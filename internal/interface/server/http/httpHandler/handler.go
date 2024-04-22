package httpHandler

import (
	"gokomodo-test/internal/interface/container"
)

type handler struct {
	OnboardHandler *onboardHandler
	SellerHandler  *sellerHandler
	BuyerHandler   *buyerHandler
}

func SetupHandlers(container *container.Container) *handler {
	return &handler{
		OnboardHandler: NewOnboardHandler(container.OnboardService),
		SellerHandler:  NewSellerHandler(container.SellerService),
		BuyerHandler:   NewBuyerHandler(container.BuyerService),
	}
}
