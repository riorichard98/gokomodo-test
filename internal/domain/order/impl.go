package order

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type repositoryOrder struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *repositoryOrder {
	if db == nil {
		panic("db is nil")
	}

	return &repositoryOrder{
		db: db,
	}
}
