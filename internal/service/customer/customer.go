package customers

import (
	"context"
	"log"
	"net/http"
	"order-management/internal/model"
	"order-management/internal/repository"
)

type CustomerServices interface {
	GetCustomerDetail(
		/*req*/ ctx context.Context, id int) (
		/*res*/ customer model.Customer, httpCode int, err error,
	)
	GetCustomerByName(
		/*req*/ ctx context.Context, name string) (
		/*res*/ customers model.Customers, httpCode int, err error,
	)
	GetAllCustomers(
		/*req*/ ctx context.Context, request model.GetAllCustomersRequest) (
		/*res*/ customers model.Customers, httpCode int, err error,
	)
	CreateCustomer(
		/*req*/ ctx context.Context, request model.CreateCustomerRequest) (
		/*res*/ httpCode int, err error,
	)
	UpdateCustomer(
		/*req*/ ctx context.Context, request model.Customer) (
		/*res*/ httpCode int, err error,
	)
	DeleteCustomer(
		/*req*/ ctx context.Context, id int) (
		/*res*/ httpCode int, err error,
	)
}

type Service struct {
	Repo repository.Repository
}

func New(repo repository.Repository) CustomerServices {
	return &Service{
		Repo: repo,
	}
}

func (x *Service) GetCustomerDetail(
	/*req*/ ctx context.Context, id int) (
	/*res*/ customer model.Customer, httpCode int, err error,
) {
	customer, err = x.Repo.GetCustomerDetail(ctx, id)
	if err != nil {
		return customer, http.StatusInternalServerError, err
	}

	return customer, http.StatusOK, nil
}

func (x *Service) GetCustomerByName(
	/*req*/ ctx context.Context, name string) (
	/*res*/ customers model.Customers, httpCode int, err error,
) {
	customers, err = x.Repo.GetCustomerByName(ctx, name)
	if err != nil {
		log.Fatal(err)
		return customers, http.StatusInternalServerError, err
	}

	return customers, http.StatusOK, nil
}

func (x *Service) GetAllCustomers(
	/*req*/ ctx context.Context, request model.GetAllCustomersRequest) (
	/*res*/ customers model.Customers, httpCode int, err error,
) {
	customers, err = x.Repo.GetAllCustomers(ctx, request)
	if err != nil {
		return customers, http.StatusInternalServerError, err
	}

	return customers, http.StatusOK, nil
}

func (x *Service) CreateCustomer(
	/*req*/ ctx context.Context, request model.CreateCustomerRequest) (
	/*res*/ httpCode int, err error,
) {
	err = x.Repo.CreateCustomer(ctx, request)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (x *Service) UpdateCustomer(
	/*req*/ ctx context.Context, request model.Customer) (
	/*res*/ httpCode int, err error,
) {
	err = x.Repo.UpdateCustomer(ctx, request)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (x *Service) DeleteCustomer(
	/*req*/ ctx context.Context, id int) (
	/*res*/ httpCode int, err error,
) {
	err = x.Repo.DeleteCustomer(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
