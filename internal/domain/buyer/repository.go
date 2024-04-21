package buyer

import "context"

type BuyerRepository interface {
	FindByEmail(ctx context.Context, email string) (buyer Buyer, err error)
	FindByID(ctx context.Context, id string) (buyer Buyer, err error)
}
