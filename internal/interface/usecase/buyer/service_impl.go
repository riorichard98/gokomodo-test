package buyer

import (
	"context"
	"fmt"
	"strconv"

	"gokomodo-test/internal/domain/buyer"
	"gokomodo-test/internal/domain/order"
	"gokomodo-test/internal/domain/product"
	"gokomodo-test/internal/domain/seller"
	"gokomodo-test/pkg/response"
	"gokomodo-test/pkg/utils"
)

type service struct {
	productRepo product.ProductRepository
	buyerRepo   buyer.BuyerRepository
	sellerRepo  seller.SellerRepository
	orderRepo   order.OrderRepository
}

func NewService(
	productRepo product.ProductRepository,
	buyerRepo buyer.BuyerRepository,
	sellerRepo seller.SellerRepository,
	orderRepo order.OrderRepository,
) *service {
	// to validate that repository is properly injected
	if productRepo == nil {
		panic("product repository is nil")
	}

	if buyerRepo == nil {
		panic("buyer repository is nil")
	}

	if sellerRepo == nil {
		panic("seller repository is nil")
	}

	if orderRepo == nil {
		panic("order repository is nil")
	}

	return &service{
		productRepo: productRepo,
		buyerRepo:   buyerRepo,
		sellerRepo:  sellerRepo,
		orderRepo:   orderRepo,
	}
}

func (s *service) validateBuyer(ctx context.Context, buyerId string) (validBuyer bool, buyer buyer.Buyer) {
	validBuyer = true
	buyer, err := s.buyerRepo.FindByID(ctx, buyerId)
	if err != nil && err.Error() != "no rows in result set" {
		utils.PrintError(err)
		validBuyer = false
		return
	}
	if buyer.ID == "" {
		validBuyer = false
	}
	return
}

func (s *service) GetAllProduct(ctx context.Context, buyerId, page, limit string) (resp response.DefaultResponse, err error) {
	validBuyer, _ := s.validateBuyer(ctx, buyerId)
	if !validBuyer {
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
	products, total, err := s.productRepo.GetProducts(ctx, "", limitNum, ((pageNum - 1) * limitNum))
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

func (s *service) OrderProduct(ctx context.Context, buyerId string, orderData OrderProductReq) (resp response.DefaultResponse, err error) {
	validBuyer, buyer := s.validateBuyer(ctx, buyerId)
	if !validBuyer {
		resp = response.ErrorResponse(response.CODE_UNAUTHORIZED, response.MESSAGE_UNAUTHORIZED)
		return
	}
	if !utils.IsValidUUID(orderData.ProductId) {
		resp = response.ErrorResponse(response.CODE_BAD_REQUEST, "invalid product id")
		return
	}
	product, err := s.productRepo.FindByID(ctx, orderData.ProductId)
	if err != nil && err.Error() != "no rows in result set" {
		utils.PrintError(err)
		resp = response.ErrorServerResponse()
		return
	}
	if product.ID == "" {
		resp = response.ErrorResponse(response.CODE_BAD_REQUEST, "product not found")
		return
	}
	seller, err := s.sellerRepo.FindByID(ctx, product.SellerID)
	if err != nil && err.Error() != "no rows in result set" {
		utils.PrintError(err)
		resp = response.ErrorServerResponse()
		return
	}
	if seller.ID == "" {
		resp = response.ErrorResponse(response.CODE_BAD_REQUEST, "seller not found")
		return
	}
	newOrder := order.Order{
		BuyerID:                    buyerId,
		SellerID:                   seller.ID,
		DeliverySourceAddress:      seller.AlamatPickup,
		DeliveryDestinationAddress: buyer.AlamatPengiriman,
		Items:                      fmt.Sprintf("%dx %s", orderData.Quantity, product.ProductName),
		Quantity:                   orderData.Quantity,
		Price:                      product.Price,
		TotalPrice:                 (product.Price * float64(orderData.Quantity)),
		Status:                     "Pending",
	}
	orderMade, err := s.orderRepo.InsertOrder(ctx, newOrder)
	if err != nil {
		utils.PrintError(err)
		resp = response.ErrorServerResponse()
		return
	}
	resp = response.CreateResponse(response.CODE_SUCCESS, response.MESSAGE_SUCCESS, orderMade)
	return
}

func (s *service) ListOrder(ctx context.Context, buyerId, page, limit string) (resp response.DefaultResponse, err error) {
	validBuyer, _ := s.validateBuyer(ctx, buyerId)
	if !validBuyer {
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
	orders, total, err := s.orderRepo.FindBySellerOrBuyer(ctx, "", buyerId, limitNum, ((pageNum - 1) * limitNum))
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
