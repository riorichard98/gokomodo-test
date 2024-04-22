package buyer

import (
	"context"

	"gokomodo-test/pkg/response"
)

type BuyerService interface {
	GetAllProduct(ctx context.Context, buyerId, page, limit string) (resp response.DefaultResponse, err error)
	OrderProduct(ctx context.Context, buyerId string, orderData OrderProductReq) (resp response.DefaultResponse, err error)
	ListOrder(ctx context.Context, buyerId, page, limit string) (resp response.DefaultResponse, err error)
}
