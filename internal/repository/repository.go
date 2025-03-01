package repository

import (
	"context"
	"database/sql"
	"order-management/internal/model"
)

type Repository interface {
	GetCustomerDetail(
		/*req*/ ctx context.Context, id int) (
		/*res*/ customer model.Customer, err error,
	)
	GetCustomerByName(
		/*req*/ ctx context.Context, name string) (
		/*res*/ customers model.Customers, err error,
	)
	GetAllCustomers(
		/*req*/ ctx context.Context, request model.GetAllCustomersRequest) (
		/*res*/ customers model.Customers, err error,
	)
	CreateCustomer(
		/*req*/ ctx context.Context, request model.CreateCustomerRequest) (
		/*res*/ err error,
	)
	UpdateCustomer(
		/*req*/ ctx context.Context, request model.Customer) (
		/*res*/ err error,
	)
	DeleteCustomer(
		/*req*/ ctx context.Context, id int) (
		/*res*/ err error,
	)

	GetOrderDetail(
		/*req*/ ctx context.Context, id int) (
		/*res*/ order model.Order, err error,
	)
	GetOrderByTitle(
		/*req*/ ctx context.Context, title string) (
		/*res*/ orders model.Orders, err error,
	)
	GetAllOrders(
		/*req*/ ctx context.Context, request model.GetAllOrdersRequest) (
		/*res*/ orders model.Orders, err error,
	)
	CreateOrder(
		/*req*/ ctx context.Context, request model.Order) (
		/*res*/ err error,
	)
	UpdateStatus(
		/*req*/ ctx context.Context, request model.Order) (
		/*res*/ err error,
	)
	DeleteOrder(
		/*req*/ ctx context.Context, id int) (
		/*res*/ err error,
	)

	CreateUser(
		/*req*/ ctx context.Context, request model.User) (
		/*res*/ err error,
	)
	GetUserByEmail(
		/*req*/ ctx context.Context, email string) (
		/*res*/ user model.User, err error,
	)
}

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repo{
		db: db,
	}
}
