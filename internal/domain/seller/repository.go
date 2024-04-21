package seller

import "context"

// SellerRepository represents the repository interface for the seller package.
type SellerRepository interface {
	// Define repository methods here...
	FindByEmail(ctx context.Context, email string) (seller Seller, err error)
	FindByID(ctx context.Context, id string) (seller Seller, err error)
}
