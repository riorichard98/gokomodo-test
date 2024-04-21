package container

import (
	buyerD "gokomodo-test/internal/domain/buyer"
	productD "gokomodo-test/internal/domain/product"
	sellerD "gokomodo-test/internal/domain/seller"

	"gokomodo-test/internal/infrastructure/postgre"
	"gokomodo-test/pkg/config"

	"gokomodo-test/internal/interface/usecase/onboard"
	"gokomodo-test/internal/interface/usecase/seller"
)

type Container struct {
	Config         *config.DefaultConfig
	DB             *config.DB
	OnboardService onboard.OnboardService
	SellerService  seller.SellerService
}

// to validate all necesseries dependencies to be injected
func (c *Container) Validate() *Container {
	if c.Config == nil {
		panic("config is nil")
	}
	if c.DB == nil {
		panic("db is nil")
	}
	if c.OnboardService == nil {
		panic("Onboard service is nil")
	}
	return c
}

func New() *Container {
	config.LoadEnv()

	// app default configuration
	defConfig := &config.DefaultConfig{
		Apps: config.Apps{
			Name: config.GetEnvString("APP_NAME"),
			// Address:  config.GetString("ADDRESS"), don't really need it for a moment
			HttpPort: config.GetEnvString("PORT"),
		},
	}

	dbConf := config.DB{
		Host:        config.GetEnvString("PSQL_HOST"),
		User:        config.GetEnvString("PSQL_USERNAME"),
		Password:    config.GetEnvString("PSQL_PASSWORD"),
		Name:        config.GetEnvString("PSQL_DBNAME"),
		Port:        config.GetEnvInteger("PSQL_PORT"),
		Timeout:     config.GetEnvInteger("PSQL_TIMEOUT"),
		MaxPoolSize: config.GetEnvInteger("PSQL_MAXPOOL_SIZE"),
		MinPoolSize: config.GetEnvInteger("PSQL_MINPOOL_SIZE"),
	}

	db := postgre.NewPgSql(dbConf)

	// repositories
	buyerRepo := buyerD.NewBuyerRepository(db)
	sellerRepo := sellerD.NewSellerRepository(db)
	productRepo := productD.NewProductRepository(db)

	// usecases
	onboardService := onboard.NewService(buyerRepo, sellerRepo)
	sellerService := seller.NewService(productRepo, sellerRepo)

	container := &Container{
		Config:         defConfig,
		DB:             &dbConf,
		OnboardService: onboardService,
		SellerService:  sellerService,
	}

	container.Validate()
	return container
}
