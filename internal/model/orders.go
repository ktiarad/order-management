package model

type Order struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	CustomerID  int     `json:"customer_id"`
	Status      string  `json:"status"`
	Price       float64 `json:"price"`
}

type Orders []Order

type GetAllOrdersRequest struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
