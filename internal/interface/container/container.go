package container

import (
	"gokomodo-test/pkg/config"
	// "gokomodo-test/internal/infrastructure/postgre"
)

type Container struct {
	Config *config.DefaultConfig
	DB     *config.DB
}

// to validate all necesseries dependencies to be injected
func (c *Container) Validate() *Container {
	if c.Config == nil {
		panic("config is nil")
	}
	if c.DB == nil {
		panic("db is nil")
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

	// db := postgre.NewPgSql(dbConf)

	container := &Container{
		Config: defConfig,
		DB:     &dbConf,
	}

	container.Validate()
	return container
}
