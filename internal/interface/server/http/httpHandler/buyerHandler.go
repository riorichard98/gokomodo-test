package httpHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"gokomodo-test/internal/interface/usecase/buyer"
	"gokomodo-test/pkg/utils"
)

type buyerHandler struct {
	buyerService buyer.BuyerService
}

func NewBuyerHandler(
	buyerService buyer.BuyerService,
) *buyerHandler {
	if buyerService == nil {
		panic("buyer service in buyer handler is nil")
	}

	return &buyerHandler{
		buyerService: buyerService,
	}
}

func (b *buyerHandler) GetAllProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()

	// Get the JWT claims from the context
	userId := utils.ClaimJWT(c.Get("user"))

	limit := c.QueryParam("limit")
	page := c.QueryParam("page")

	resp, _ := b.buyerService.GetAllProduct(ctx, userId, page, limit)
	return c.JSON(http.StatusOK, resp)
}

func (b *buyerHandler) OrderProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()

	req := buyer.OrderProductReq{}
	if err = utils.Validate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request format")
	}

	// Get the JWT claims from the context
	userId := utils.ClaimJWT(c.Get("user"))

	resp, _ := b.buyerService.OrderProduct(ctx, userId, req)
	return c.JSON(http.StatusOK, resp)
}

func (b *buyerHandler) GetAllOrder(c echo.Context) (err error) {
	ctx := c.Request().Context()

	// Get the JWT claims from the context
	userId := utils.ClaimJWT(c.Get("user"))

	limit := c.QueryParam("limit")
	page := c.QueryParam("page")

	resp, _ := b.buyerService.ListOrder(ctx, userId, page, limit)
	return c.JSON(http.StatusOK, resp)
}