package seller

import (
	"context"

	"gokomodo-test/pkg/response"
)

type SellerService interface {
	AddNewProduct(ctx context.Context, productData AddProduct, sellerId string) (resp response.DefaultResponse, err error)
}
