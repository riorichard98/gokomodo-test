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
		insert into product (product_name, description, price, seller_id)
		values ($1, $2, $3, $4)
	`
	_, err = r.db.Exec(ctx, query, product.ProductName, product.Description, product.Price, product.SellerID)
	return
}

func (r *repositoryProduct) GetProducts(ctx context.Context, sellerID string, limit, offset int) (products []Product, total int, err error) {
	var query string
	var totalQuery string
	var args []interface{}
	var countArgs []interface{}

	// Check if sellerID is empty
	if sellerID == "" {
		query = `
        select * from product limit $1 offset $2`
		totalQuery = `select count(*) as total from product`
		args = []interface{}{limit, offset}
	} else {
		query = `
        select * from product where seller_id = $1 limit $2 offset $3`
		totalQuery = `select count(*) as total from product where seller_id = $1`
		args = []interface{}{sellerID, limit, offset}
		countArgs = []interface{}{sellerID}
	}

	// Fetch products
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	products = []Product{}
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.ProductName, &p.Description, &p.Price, &p.SellerID)
		if err != nil {
			return
		}
		products = append(products, p)
	}

	// Fetch total count
	row := r.db.QueryRow(ctx, totalQuery, countArgs...)
	totalRow := Total{}
	err = row.Scan(&totalRow.Total)
	total = totalRow.Total
	return
}

func (r *repositoryProduct) FindByID(ctx context.Context, id string) (product Product, err error) {
	query := "select * from product where id = $1"
	row := r.db.QueryRow(ctx, query, id)

	product = Product{}
	err = row.Scan(&product.ID, &product.ProductName, &product.Description, &product.Price, &product.SellerID)

	return
}
