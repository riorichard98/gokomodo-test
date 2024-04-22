package order

import (
	"context"
)

// OrderRepository represents the repository interface for the order package.
type OrderRepository interface {
	// Define repository methods here...
	InsertOrder(ctx context.Context, newOrder Order) (order Order, err error)
	FindBySellerOrBuyer(ctx context.Context, sellerId, buyerId string, limit, offset int) (orders []Order, total int, err error)
	UpdateOrderStatus(ctx context.Context, orderID, status string) (err error)
	FindByID(ctx context.Context, id string) (order Order, err error)
}
