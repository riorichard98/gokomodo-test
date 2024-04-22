package seller

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repositorySeller struct {
	db *pgxpool.Pool
}

func NewSellerRepository(db *pgxpool.Pool) *repositorySeller {
	if db == nil {
		panic("db is nil")
	}

	return &repositorySeller{
		db: db,
	}
}

func (r *repositorySeller) FindByEmail(ctx context.Context, email string) (seller Seller, err error) {
	query := "select * from seller where email = $1"
	row := r.db.QueryRow(ctx, query, email)

	seller = Seller{}
	err = row.Scan(&seller.ID, &seller.Email, &seller.Name, &seller.Password, &seller.AlamatPickup)
	if err != nil {
		return Seller{}, err
	}

	return seller, nil
}

func (r *repositorySeller) FindByID(ctx context.Context, id string) (seller Seller, err error) {
	query := "select * from seller where id = $1"
	row := r.db.QueryRow(ctx, query, id)

	seller = Seller{}
	err = row.Scan(&seller.ID, &seller.Email, &seller.Name, &seller.Password, &seller.AlamatPickup)

	return
}
