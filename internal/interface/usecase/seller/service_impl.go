package seller

import (
	"context"
	"strconv"

	"gokomodo-test/internal/domain/order"
	"gokomodo-test/internal/domain/product"
	"gokomodo-test/internal/domain/seller"

	"gokomodo-test/pkg/response"
	"gokomodo-test/pkg/utils"
)

type service struct {
	productRepo product.ProductRepository
	sellerRepo  seller.SellerRepository
	orderRepo   order.OrderRepository
}

func NewService(
	productRepo product.ProductRepository,
	sellerRepo seller.SellerRepository,
	orderRepo order.OrderRepository,
) *service {
	// to validate that repository is properly injected
	if productRepo == nil {
		panic("product repository is nil")
	}

	if sellerRepo == nil {
		panic("seller repository is nil")
	}

	if orderRepo == nil {
		panic("order repository is nil")
	}

	return &service{
		productRepo: productRepo,
		sellerRepo:  sellerRepo,
		orderRepo:   orderRepo,
	}
}

func (s *service) validateSeller(ctx context.Context, sellerId string) (validSeller bool) {
	validSeller = true
	seller, err := s.sellerRepo.FindByID(ctx, sellerId)
	if err != nil && err.Error() != "no rows in result set" {
		utils.PrintError(err)
		validSeller = false
		return
	}
	if seller.ID == "" {
		validSeller = false
	}
	return
}

func (s *service) AddNewProduct(ctx context.Context, productData AddProduct, sellerId string) (resp response.DefaultResponse, err error) {
	if !s.validateSeller(ctx, sellerId) {
		resp = response.ErrorResponse(response.CODE_UNAUTHORIZED, response.MESSAGE_UNAUTHORIZED)
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

func (s *service) GetListProduct(ctx context.Context, sellerId, page, limit string) (resp response.DefaultResponse, err error) {
	if !s.validateSeller(ctx, sellerId) {
		resp = response.ErrorResponse(response.CODE_UNAUTHORIZED, response.MESSAGE_UNAUTHORIZED)
		return
	}
	pageNum, _ := strconv.Atoi(page)
	if pageNum < 1 {
		pageNum = 1
	}
	limitNum, _ := strconv.Atoi(limit)
	if limitNum < 1 {
		limitNum = 10
	}
	products, total, err := s.productRepo.GetProducts(ctx, sellerId, limitNum, ((pageNum - 1) * limitNum))
	if err != nil {
		utils.PrintError(err)
		resp = response.ErrorServerResponse()
		return
	}
	items := ItemList{
		Items: products,
		Total: total,
	}
	resp = response.CreateResponse(response.CODE_SUCCESS, response.MESSAGE_SUCCESS, items)
	return
}

func (s *service) ListOrder(ctx context.Context, sellerId, page, limit string) (resp response.DefaultResponse, err error) {
	if !s.validateSeller(ctx, sellerId) {
		resp = response.ErrorResponse(response.CODE_UNAUTHORIZED, response.MESSAGE_UNAUTHORIZED)
		return
	}
	pageNum, _ := strconv.Atoi(page)
	if pageNum < 1 {
		pageNum = 1
	}
	limitNum, _ := strconv.Atoi(limit)
	if limitNum < 1 {
		limitNum = 10
	}
	orders, total, err := s.orderRepo.FindBySellerOrBuyer(ctx, sellerId, "", limitNum, ((pageNum - 1) * limitNum))
	if err != nil {
		utils.PrintError(err)
		resp = response.ErrorServerResponse()
		return
	}
	items := ItemList{
		Items: orders,
		Total: total,
	}
	resp = response.CreateResponse(response.CODE_SUCCESS, response.MESSAGE_SUCCESS, items)
	return
}

func (s *service) AcceptOrder(ctx context.Context, sellerId, orderId string) (resp response.DefaultResponse, err error) {
	if !s.validateSeller(ctx, sellerId) {
		resp = response.ErrorResponse(response.CODE_UNAUTHORIZED, response.MESSAGE_UNAUTHORIZED)
		return
	}
	if !utils.IsValidUUID(orderId) {
		resp = response.ErrorResponse(response.CODE_BAD_REQUEST, "invalid product id")
		return
	}
	orderFound, err := s.orderRepo.FindByID(ctx, orderId)
	if err != nil && err.Error() != "no rows in result set" {
		utils.PrintError(err)
		resp = response.ErrorServerResponse()
		return
	}
	if orderFound.ID == "" {
		resp = response.ErrorResponse(response.CODE_BAD_REQUEST, "order not found")
		return
	}
	err = s.orderRepo.UpdateOrderStatus(ctx, orderId, "Accepted")
	if err != nil {
		utils.PrintError(err)
		resp = response.ErrorServerResponse()
		return
	}
	resp = response.CreateResponse(response.CODE_SUCCESS, response.MESSAGE_SUCCESS, nil)
	return
}
