package product

import "context"

// ProductRepository represents the repository interface for the product package.
type ProductRepository interface {
	// Define repository methods here...
	InsertProduct(ctx context.Context, product Product) (err error)
	GetProducts(ctx context.Context, sellerID string, limit, offset int) (products []Product, total int, err error)
	FindByID(ctx context.Context, id string) (product Product, err error)
}
