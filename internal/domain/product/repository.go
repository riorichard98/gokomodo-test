package product

import "context"

// ProductRepository represents the repository interface for the product package.
type ProductRepository interface {
	// Define repository methods here...
	InsertProduct(ctx context.Context, product Product) (err error)
}
