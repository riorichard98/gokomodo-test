package seller

import (
	"context"

	"gokomodo-test/pkg/response"
)

type SellerService interface {
	AddNewProduct(ctx context.Context, productData AddProduct, sellerId string) (resp response.DefaultResponse, err error)
	GetListProduct(ctx context.Context, sellerId, page, limit string) (resp response.DefaultResponse, err error)
	ListOrder(ctx context.Context, sellerId, page, limit string) (resp response.DefaultResponse, err error)
	AcceptOrder(ctx context.Context, sellerId, orderId string) (resp response.DefaultResponse, err error)
}
