package buyer

import (
	"context"
	"testing"

	buyerD "gokomodo-test/internal/domain/buyer"
	orderD "gokomodo-test/internal/domain/order"
	productD "gokomodo-test/internal/domain/product"
	sellerD "gokomodo-test/internal/domain/seller"

	"gokomodo-test/internal/infrastructure/postgre"
	"gokomodo-test/pkg/config"
	// "gokomodo-test/pkg/response"
)

// possible required data
var mockService service
var productId string
var buyerId string
var ctx = context.Background()

// setup for required data
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

	buyerRepo := buyerD.NewBuyerRepository(db)
	sellerRepo := sellerD.NewSellerRepository(db)
	productRepo := productD.NewProductRepository(db)
	orderRepo := orderD.NewOrderRepository(db)

	mockService = *NewService(productRepo, buyerRepo, sellerRepo, orderRepo)

	seller, _ := sellerRepo.FindByEmail(ctx, "rio@gmail.com")
	buyer, _ := buyerRepo.FindByEmail(ctx, "rich@gmail.com")
	buyerId = buyer.ID
	productRepo.InsertProduct(ctx, productD.Product{
		ProductName: "Product unit testing",
		Description: "for unit testing",
		Price:       3003.98,
		SellerID:    seller.ID,
	})
	products, _, _ := productRepo.GetProducts(ctx, "", 10, 0)
	for _, p := range products {
		if p.ProductName == "Product unit testing" {
			productId = p.ID
		}
	}
	m.Run()
}

func TestListProduct_Success(t *testing.T) {
	// Call the GetAllProduct function
	resp, err := mockService.GetAllProduct(ctx, buyerId, "1", "10")

	// Check if there's any error
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if resp.Data is of type ItemList
	itemList, ok := resp.Data.(ItemList)
	if !ok {
		t.Errorf("Expected resp.Data to be of type ItemList, got %T", resp.Data)
	}

	// Check if Items field of ItemList is a slice
	itemsSlice, ok := itemList.Items.([]productD.Product)
	if !ok {
		t.Errorf("Expected Items field to be a slice, got %T", itemList.Items)
	}

	// Check the length of the items slice
	if len(itemsSlice) < 1 {
		t.Errorf("Expected at least one item, got %d", len(itemsSlice))
	}
}
func TestOrderProductAndOrderList_Success(t *testing.T) {
	resp, err := mockService.OrderProduct(ctx, buyerId, OrderProductReq{
		ProductId: productId,
		Quantity:  10,
	})

	// Check if there's any error
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	orderMade, ok := resp.Data.(orderD.Order)
	if !ok {
		t.Errorf("Expected resp.Data to be of type orderD.Order, got %T", resp.Data)
	}

	if orderMade.ID == "" {
		t.Errorf("Expected order made id is already exist")
	}

	resp2, err := mockService.ListOrder(ctx, buyerId, "1", "10")

	// Check if there's any error
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if resp.Data is of type ItemList
	itemList, ok := resp2.Data.(ItemList)
	if !ok {
		t.Errorf("Expected resp.Data to be of type ItemList, got %T", resp.Data)
	}

	// Check if Items field of ItemList is a slice
	itemsSlice, ok := itemList.Items.([]orderD.Order)
	if !ok {
		t.Errorf("Expected Items field to be a slice, got %T", itemList.Items)
	}

	// Check the length of the items slice
	if len(itemsSlice) < 1 {
		t.Errorf("Expected at least one item, got %d", len(itemsSlice))
	}

}
