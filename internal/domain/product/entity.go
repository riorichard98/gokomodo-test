package product

// Product represents the product table in the database.
type Product struct {
	ID          string  `db:"id"`
	ProductName string  `db:"product_name"`
	Description string  `db:"description,omitempty"`
	Price       float64 `db:"price"`
	SellerID    string  `db:"seller_id"`
}

type Total struct {
	Total int `db:"total"`
}
