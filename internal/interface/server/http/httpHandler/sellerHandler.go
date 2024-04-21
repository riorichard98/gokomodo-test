package httpHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"gokomodo-test/internal/interface/usecase/onboard"
	"gokomodo-test/internal/interface/usecase/seller"
	"gokomodo-test/pkg/utils"
)

type sellerHandler struct {
	sellerService  seller.SellerService
	onboardService onboard.OnboardService
}

func NewSellerHandler(
	sellerService seller.SellerService,
	onboardService onboard.OnboardService,
) *sellerHandler {
	if sellerService == nil {
		panic("seller service in seller handler is nil")
	}

	if onboardService == nil {
		panic("onboard service in onboard handler is nil")
	}

	return &sellerHandler{
		sellerService:  sellerService,
		onboardService: onboardService,
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
