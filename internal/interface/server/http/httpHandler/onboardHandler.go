package httpHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"gokomodo-test/internal/interface/usecase/onboard"
	"gokomodo-test/pkg/utils"
)

type onboardHandler struct {
	onboardService onboard.OnboardService
}

func NewOnboardHandler(
	onboardService onboard.OnboardService,
) *onboardHandler {
	if onboardService == nil {
		panic("onboard service in onboard handler is nil")
	}

	return &onboardHandler{
		onboardService: onboardService,
	}
}

func (o *onboardHandler) LoginHandler(c echo.Context) (err error) {
	ctx := c.Request().Context()

	req := onboard.Login{}

	if err = utils.Validate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request format")
	}

	resp, _ := o.onboardService.Login(ctx, req)
	return c.JSON(http.StatusOK, resp)
}
