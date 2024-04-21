package buyer

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repositoryBuyer struct {
	db *pgxpool.Pool
}

func NewBuyerRepository(db *pgxpool.Pool) *repositoryBuyer {
	if db == nil {
		panic("db is nil")
	}

	return &repositoryBuyer{
		db: db,
	}
}

func (r *repositoryBuyer) FindByEmail(ctx context.Context, email string) (buyer Buyer, err error) {
	query := "select * from buyer where email = $1"
	row := r.db.QueryRow(ctx, query, email)

	buyer = Buyer{}
	err = row.Scan(&buyer.ID, &buyer.Email, &buyer.Name, &buyer.Password, &buyer.AlamatPengiriman)
	if err != nil {
		return Buyer{}, err
	}

	return buyer, nil
}

func (r *repositoryBuyer) FindByID(ctx context.Context, id string) (buyer Buyer, err error) {
	query := "select * from buyer where id = $1"
	row := r.db.QueryRow(ctx, query, id)

	buyer = Buyer{}
	err = row.Scan(&buyer.ID, &buyer.Email, &buyer.Name, &buyer.Password, &buyer.AlamatPengiriman)
	if err != nil {
		return Buyer{}, err
	}

	return buyer, nil
}
