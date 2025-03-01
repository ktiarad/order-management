package service

import (
	"order-management/internal/repository"
	"order-management/internal/service/authentication"
	customers "order-management/internal/service/customer"
	orders "order-management/internal/service/order"
)

type Services struct {
	customers.CustomerServices
	orders.OrderServices
	authentication.AuthenticationServices
}

func NewService(
	repo repository.Repository,
) Services {
	customerServices := customers.New(repo)
	orderServices := orders.New(repo)
	authenticationServices := authentication.New(repo)

	return Services{
		customerServices,
		orderServices,
		authenticationServices,
	}
}
