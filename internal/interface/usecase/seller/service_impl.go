package seller

import (
	"context"

	"gokomodo-test/internal/domain/product"
	"gokomodo-test/internal/domain/seller"
	"gokomodo-test/pkg/response"
	"gokomodo-test/pkg/utils"
)

type service struct {
	productRepo product.ProductRepository
	sellerRepo  seller.SellerRepository
}

func NewService(
	productRepo product.ProductRepository,
	sellerRepo seller.SellerRepository,
) *service {
	// to validate that repository is properly injected
	if productRepo == nil {
		panic("product repository is nil")
	}

	if sellerRepo == nil {
		panic("seller repository is nil")
	}

	return &service{
		productRepo: productRepo,
		sellerRepo:  sellerRepo,
	}
}

func (s *service) AddNewProduct(ctx context.Context, productData AddProduct, sellerId string) (resp response.DefaultResponse, err error) {
	seller, err := s.sellerRepo.FindByID(ctx, sellerId)
	if err != nil && err.Error() != "no rows in result set" {
		utils.PrintError(err)
		resp = response.ErrorServerResponse()
		return
	}
	if seller.ID == "" {
		resp = response.ErrorResponse(response.CODE_UNAUTHORIZED, response.CODE_UNAUTHORIZED)
		return
	}
	newProduct := product.Product{
		ProductName: productData.ProductName,
		Description: productData.Description,
		Price:       productData.Price,
		SellerID:    sellerId,
	}

	err = s.productRepo.InsertProduct(ctx, newProduct)
	if err != nil {
		utils.PrintError(err)
		resp = response.ErrorServerResponse()
		return
	}
	resp = response.CreateResponse(response.CODE_SUCCESS, response.MESSAGE_SUCCESS, nil)
	return
}
