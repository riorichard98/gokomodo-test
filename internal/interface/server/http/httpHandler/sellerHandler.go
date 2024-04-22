package httpHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"gokomodo-test/internal/interface/usecase/seller"
	"gokomodo-test/pkg/utils"
)

type sellerHandler struct {
	sellerService seller.SellerService
}

func NewSellerHandler(
	sellerService seller.SellerService,
) *sellerHandler {
	if sellerService == nil {
		panic("seller service in seller handler is nil")
	}

	return &sellerHandler{
		sellerService: sellerService,
	}
}

func (s *sellerHandler) AddNewProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()

	req := seller.AddProduct{}

	if err = utils.Validate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request format")
	}

	// Get the JWT claims from the context
	userId := utils.ClaimJWT(c.Get("user"))

	resp, _ := s.sellerService.AddNewProduct(ctx, req, userId)
	return c.JSON(http.StatusOK, resp)
}

func (s *sellerHandler) GetProductList(c echo.Context) (err error) {
	ctx := c.Request().Context()

	// Get the JWT claims from the context
	userId := utils.ClaimJWT(c.Get("user"))

	limit := c.QueryParam("limit")
	page := c.QueryParam("page")

	resp, _ := s.sellerService.GetListProduct(ctx, userId, page, limit)
	return c.JSON(http.StatusOK, resp)
}

func (s *sellerHandler) OrderList(c echo.Context) (err error) {
	ctx := c.Request().Context()

	// Get the JWT claims from the context
	userId := utils.ClaimJWT(c.Get("user"))

	limit := c.QueryParam("limit")
	page := c.QueryParam("page")

	resp, _ := s.sellerService.ListOrder(ctx, userId, page, limit)
	return c.JSON(http.StatusOK, resp)
}

func (s *sellerHandler) AcceptOrder(c echo.Context) (err error) {
	ctx := c.Request().Context()

	// Get the JWT claims from the context
	userId := utils.ClaimJWT(c.Get("user"))

	orderId := c.QueryParam("order")

	resp, _ := s.sellerService.AcceptOrder(ctx, userId, orderId)
	return c.JSON(http.StatusOK, resp)
}
