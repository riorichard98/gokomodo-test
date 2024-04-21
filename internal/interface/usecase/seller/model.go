package seller

type AddProduct struct {
	ProductName string  `json:"product_name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}
