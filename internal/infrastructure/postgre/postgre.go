package postgre

import (
	"context"
	"fmt"
	"gokomodo-test/pkg/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPgSql creates a new connection to a PostgreSQL database
func NewPgSql(db config.DB) *pgxpool.Pool {
	ctx := context.Background()

	// Construct the connection string
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		db.User, db.Password, db.Host, db.Port, db.Name)

	// Create a new pool configuration
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		panic(fmt.Sprintf("unable to parse connection string: %v", err.Error()))
	}
	// Set additional configuration options
	config.MaxConns = int32(db.MaxPoolSize)
	config.MinConns = int32(db.MinPoolSize)
	config.ConnConfig.ConnectTimeout = time.Duration(db.Timeout) * time.Second

	// Create a new connection pool with the configured options
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic(fmt.Sprintf("unable to create connection pool: %v", err.Error()))
	}

	// Ping the connection
	if err := pool.Ping(ctx); err != nil {
		panic(fmt.Sprintf("unable to connect with postgresql: %v", err.Error()))
	}

	fmt.Println("postgre database is connected")
	return pool
}
