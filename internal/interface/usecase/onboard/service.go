package onboard

import (
	"context"

	"gokomodo-test/pkg/response"
)

type OnboardService interface {
	Login(ctx context.Context, loginData Login) (resp response.DefaultResponse, err error)
}
