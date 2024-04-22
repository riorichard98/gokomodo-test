package onboard

import (
	"context"
	"testing"

	buyerD "gokomodo-test/internal/domain/buyer"
	sellerD "gokomodo-test/internal/domain/seller"

	"gokomodo-test/internal/infrastructure/postgre"
	"gokomodo-test/pkg/config"
	"gokomodo-test/pkg/response"
)

var mockService service

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

	mockService = *NewService(buyerRepo, sellerRepo)

	m.Run()
}

func TestLogin_Success(t *testing.T) {
	// Prepare valid login data
	validLoginData := []Login{
		{Email: "rio@gmail.com", Password: "123", Type: "seller"},
		{Email: "rio2@gmail.com", Password: "123", Type: "seller"},
		{Email: "rich@gmail.com", Password: "456", Type: "buyer"},
	}

	for _, loginData := range validLoginData {
		// Call the Login function with the prepared valid login data
		resp, err := mockService.Login(context.Background(), loginData)

		// Check if there's any error during login
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Check if the response status code is success
		if resp.Status != response.CODE_SUCCESS {
			t.Errorf("Expected success response, got %s", resp.Status)
		}

		// You might want to perform additional checks here depending on your response structure.
		// For example, you could check if the returned token is not empty.

		// Example additional check:
		if resp.Data.(Loginresp).Token == "" {
			t.Error("Expected non-empty token in response")
		}
	}
}

func TestLogin_InvalidPassword(t *testing.T) {
	// Prepare invalid login data with incorrect password
	invalidPasswordData := Login{
		Email:    "rio@gmail.com",
		Password: "wrongpassword",
		Type:     "seller",
	}

	// Call the Login function with the prepared invalid login data
	resp, _ := mockService.Login(context.Background(), invalidPasswordData)

	// Check if the response status code is not success
	if resp.Status != response.CODE_BAD_REQUEST {
		t.Errorf("Expected failure response, got %s", resp.Status)
	}

	// Check if the response message indicates invalid password
	expectedErrorMsg := "invalid password"
	if resp.Message != expectedErrorMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMsg, resp.Message)
	}
}

func TestLogin_InvalidEmail(t *testing.T) {
	// Prepare invalid login data with non-existing email
	invalidEmailData := Login{
		Email:    "nonexisting@gmail.com",
		Password: "123",
		Type:     "seller",
	}

	// Call the Login function with the prepared invalid login data
	resp, _ := mockService.Login(context.Background(), invalidEmailData)

	// Check if the response status code is not success
	if resp.Status != response.CODE_BAD_REQUEST {
		t.Errorf("Expected failure response, got %s", resp.Status)
	}

	// Check if the response message indicates user not found
	expectedErrorMsg := "user not found"
	if resp.Message != expectedErrorMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMsg, resp.Message)
	}
}
