package model

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Customers []Customer

type CreateCustomerRequest struct {
	Name string `json:"name"`
}

type GetAllCustomersRequest struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
