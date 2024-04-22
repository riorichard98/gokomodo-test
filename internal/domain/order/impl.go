package order

import (
	"context"

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

func (r *repositoryOrder) InsertOrder(ctx context.Context, newOrder Order) (order Order, err error) {
	query := `
        INSERT INTO "order" (buyer_id, seller_id, delivery_source_address, delivery_destination_address, items, quantity, price, total_price, status)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id, buyer_id, seller_id, delivery_source_address, delivery_destination_address, items, quantity, price, total_price, status
    `
	row := r.db.QueryRow(ctx, query, newOrder.BuyerID, newOrder.SellerID, newOrder.DeliverySourceAddress, newOrder.DeliveryDestinationAddress, newOrder.Items, newOrder.Quantity, newOrder.Price, newOrder.TotalPrice, newOrder.Status)
	order = Order{}
	err = row.Scan(&order.ID, &order.BuyerID, &order.SellerID, &order.DeliverySourceAddress, &order.DeliveryDestinationAddress, &order.Items, &order.Quantity, &order.Price, &order.TotalPrice, &order.Status)
	return
}

func (r *repositoryOrder) FindBySellerOrBuyer(ctx context.Context, sellerId, buyerId string, limit, offset int) (orders []Order, total int, err error) {
	var query string
	var totalQuery string
	var id string

	// Check if buyerId is not empty
	if buyerId != "" {
		query = `
        select * from "order" where buyer_id = $1 limit $2 offset $3 `
		totalQuery = `select count(*) as total from "order" where buyer_id = $1`
		id = buyerId
	} else if sellerId != "" {
		query = `
        select * from "order" where seller_id = $1 limit $2 offset $3 `
		totalQuery = `select count(*) as total from "order" where seller_id = $1`
		id = sellerId
	}

	// Fetch orders
	rows, err := r.db.Query(ctx, query, id, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	orders = []Order{}
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.ID, &order.BuyerID, &order.SellerID, &order.DeliverySourceAddress, &order.DeliveryDestinationAddress, &order.Items, &order.Quantity, &order.Price, &order.TotalPrice, &order.Status)
		if err != nil {
			return
		}
		orders = append(orders, order)
	}
	// Fetch total count
	row := r.db.QueryRow(ctx, totalQuery, id)
	totalRow := Total{}
	err = row.Scan(&totalRow.Total)
	total = totalRow.Total
	return
}

func (r *repositoryOrder) UpdateOrderStatus(ctx context.Context, orderID, status string) (err error) {
	query := `
		update "order"
		set status = $1
		where id = $2
	`
	_, err = r.db.Exec(ctx, query, status, orderID)

	return
}

func (r *repositoryOrder) FindByID(ctx context.Context, id string) (order Order, err error) {
	query := `select * from "order" where id = $1`
	row := r.db.QueryRow(ctx, query, id)

	order = Order{}
	err = row.Scan(&order.ID, &order.BuyerID, &order.SellerID, &order.DeliverySourceAddress, &order.DeliveryDestinationAddress, &order.Items, &order.Quantity, &order.Price, &order.TotalPrice, &order.Status)

	return
}
