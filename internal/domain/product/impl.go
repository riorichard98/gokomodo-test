package product

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repositoryProduct struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *repositoryProduct {
	if db == nil {
		panic("db is nil")
	}

	return &repositoryProduct{
		db: db,
	}
}

func (r *repositoryProduct) InsertProduct(ctx context.Context, product Product) (err error) {
	query := `
		INSERT INTO product (product_name, description, price, seller_id)
		VALUES ($1, $2, $3, $4)
	`
	_, err = r.db.Exec(ctx, query, product.ProductName, product.Description, product.Price, product.SellerID)
	if err != nil {
		return err
	}
	return nil
}
