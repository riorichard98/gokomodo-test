package order

// Order represents the order table in the database.
type Order struct {
	ID                         string  `db:"id"`
	BuyerID                    string  `db:"buyer_id"`
	SellerID                   string  `db:"seller_id"`
	DeliverySourceAddress      string  `db:"delivery_source_address"`
	DeliveryDestinationAddress string  `db:"delivery_destination_address"`
	Items                      string  `db:"items"`
	Quantity                   int     `db:"quantity"`
	Price                      float64 `db:"price,omitempty"`
	TotalPrice                 float64 `db:"total_price,omitempty"`
	Status                     string  `db:"status,omitempty"`
}

type Total struct {
	Total int `db:"total"`
}
