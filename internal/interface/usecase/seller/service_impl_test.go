package seller

import (
	"context"
	"fmt"
	"testing"

	buyerD "gokomodo-test/internal/domain/buyer"
	orderD "gokomodo-test/internal/domain/order"
	productD "gokomodo-test/internal/domain/product"
	sellerD "gokomodo-test/internal/domain/seller"

	"gokomodo-test/internal/infrastructure/postgre"
	"gokomodo-test/pkg/config"
	"gokomodo-test/pkg/response"
)

// possible required data
var mockService service
var orderId string
var sellerId string
var productId string
var ctx = context.Background()

func TestMain(m *testing.M) {
	config.LoadEnv("../../../../.env")
	dbTestConf := config.DB{
		Host:        config.GetEnvString("PSQL_HOST"),
		User:        config.GetEnvString("PSQL_USERNAME"),
		Password:    config.GetEnvString("PSQL_PASSWORD"),
		Name:        config.GetEnvString("PSQL_DBNAME"),
		Port:        config.GetEnvInteger("PSQL_PORT"),
		Timeout:     config.GetEnvInteger("PSQL_TIMEOUT"),
		MaxPoolSize: config.GetEnvInteger("PSQL_MAXPOOL_SIZE"),
		MinPoolSize: config.GetEnvInteger("PSQL_MINPOOL_SIZE"),
	}

	db := postgre.NewPgSql(dbTestConf)

	sellerRepo := sellerD.NewSellerRepository(db)
	productRepo := productD.NewProductRepository(db)
	orderRepo := orderD.NewOrderRepository(db)
	buyerRepo := buyerD.NewBuyerRepository(db)

	mockService = *NewService(productRepo, sellerRepo, orderRepo)

	seller, _ := sellerRepo.FindByEmail(ctx, "rio@gmail.com")
	sellerId = seller.ID
	buyer, _ := buyerRepo.FindByEmail(ctx, "rich@gmail.com")
	productRepo.InsertProduct(ctx, productD.Product{
		ProductName: "Product unit testing(seller)",
		Description: "for unit testing",
		Price:       3003.98,
		SellerID:    seller.ID,
	})
	products, _, _ := productRepo.GetProducts(ctx, "", 10, 0)
	for _, p := range products {
		if p.ProductName == "Product unit testing(seller)" {
			productId = p.ID
		}
	}
	order, _ := orderRepo.InsertOrder(ctx, orderD.Order{
		BuyerID:                    buyer.ID,
		SellerID:                   sellerId,
		DeliverySourceAddress:      seller.AlamatPickup,
		DeliveryDestinationAddress: buyer.AlamatPengiriman,
		Items:                      fmt.Sprintf("%dx %s", 100, "Product unit testing(seller)"),
		Quantity:                   100,
		Price:                      3003.98,
		TotalPrice:                 (100 * 3003.98),
		Status:                     "Pending",
	})
	orderId = order.ID
	m.Run()
}

func TestAddAndListProduct_Success(t *testing.T) {
	// Define the new product to be added
	newProduct := AddProduct{
		ProductName: "New Product",
		Description: "Description of the new product",
		Price:       99.99,
	}

	// Add the new product
	resp, err := mockService.AddNewProduct(ctx, newProduct, sellerId)
	if err != nil {
		t.Errorf("Error adding product: %v", err)
	}

	// Check the response for adding the product
	if resp.Status != response.CODE_SUCCESS {
		t.Errorf("Expected success response, got: %v", resp)
	}

	// Get the list of products
	page := "1"
	limit := "10"
	respList, err := mockService.GetListProduct(ctx, sellerId, page, limit)
	if err != nil {
		t.Errorf("Error getting product list: %v", err)
	}

	// Check the response for getting the product list
	if respList.Status != response.CODE_SUCCESS {
		t.Errorf("Expected success response for getting product list, got: %v", respList)
	}

	// Check if the newly added product exists in the list
	found := false
	if productList, ok := respList.Data.(ItemList); ok {
		for _, item := range productList.Items.([]productD.Product) {
			if item.ProductName == newProduct.ProductName && item.Description == newProduct.Description && item.Price == newProduct.Price {
				found = true
				break
			}
		}
	}

	// Assert that the newly added product exists in the list
	if !found {
		t.Errorf("Newly added product not found in the product list")
	}
}

func TestAcceptAndGetListOrder_Success(t *testing.T) {
	// Accept the order
	resp, err := mockService.AcceptOrder(ctx, sellerId, orderId)
	if err != nil {
		t.Errorf("Error accepting order: %v", err)
	}

	// Check the response for accepting the order
	if resp.Status != response.CODE_SUCCESS {
		t.Errorf("Expected success response for accepting order, got: %v", resp)
	}

	// Get the list of orders
	page := "1"
	limit := "10"
	respList, err := mockService.ListOrder(ctx, sellerId, page, limit)
	if err != nil {
		t.Errorf("Error getting order list: %v", err)
	}

	// Check the response for getting the order list
	if respList.Status != response.CODE_SUCCESS {
		t.Errorf("Expected success response for getting order list, got: %v", respList)
	}

	// Check if the accepted order exists in the list with status "Accepted"
	found := false
	if orderList, ok := respList.Data.(ItemList); ok {
		for _, item := range orderList.Items.([]orderD.Order) {
			if item.ID == orderId && item.Status == "Accepted" {
				found = true
				break
			}
		}
	}

	// Assert that the accepted order exists in the list with status "Accepted"
	if !found {
		t.Errorf("Accepted order not found in the order list with status 'Accepted'")
	}
}
