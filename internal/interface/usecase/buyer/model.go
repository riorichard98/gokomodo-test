package buyer

type ItemList struct {
	Items interface{} `json:"items"`
	Total int         `json:"total"`
}

type OrderProductReq struct {
	ProductId string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,numeric"`
}
