package http

import (
	"net/http"

	"gokomodo-test/internal/interface/container"
	"gokomodo-test/internal/interface/server/http/httpHandler"

	"github.com/labstack/echo/v4"
)

func SetupRouter(server *echo.Echo, cont *container.Container) {
	handler := httpHandler.SetupHandlers(cont)

	// - health check
	server.GET("/ping", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "services up and running... ")
	})

	onboardRoute := server.Group("/onboard")
	onboardRoute.POST("/login", handler.OnboardHandler.LoginHandler)

	sellerRoute := server.Group("/sellers")
	sellerRoute.POST("/product", handler.SellerHandler.AddNewProduct)
	sellerRoute.GET("/products", handler.SellerHandler.GetProductList)
	sellerRoute.GET("/orders", handler.SellerHandler.OrderList)
	sellerRoute.POST("/orders", handler.SellerHandler.AcceptOrder)

	buyerRoute := server.Group("/buyers")
	buyerRoute.GET("/products", handler.BuyerHandler.GetAllProduct)
	buyerRoute.POST("/order", handler.BuyerHandler.OrderProduct)
	buyerRoute.GET("/order", handler.BuyerHandler.GetAllOrder)
}
