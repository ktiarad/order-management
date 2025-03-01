package orders

import (
	"context"
	"net/http"
	"order-management/internal/model"
	"order-management/internal/repository"
)

type OrderServices interface {
	GetOrderDetail(
		/*req*/ ctx context.Context, id int) (
		/*res*/ order model.Order, httpCode int, err error,
	)
	GetOrderByTitle(
		/*req*/ ctx context.Context, title string) (
		/*res*/ orders model.Orders, httpCode int, err error,
	)
	GetAllOrders(
		/*req*/ ctx context.Context, request model.GetAllOrdersRequest) (
		/*res*/ orders model.Orders, httpCode int, err error,
	)
	CreateOrder(
		/*req*/ ctx context.Context, request model.Order) (
		/*res*/ httpCode int, err error,
	)
	UpdateOrderStatus(
		/*req*/ ctx context.Context, request model.Order) (
		/*res*/ httpCode int, err error,
	)
	DeleteOrder(
		/*req*/ ctx context.Context, id int) (
		/*res*/ httpCode int, err error,
	)
}

type Service struct {
	Repo repository.Repository
}

func New(repo repository.Repository) OrderServices {
	return &Service{
		Repo: repo,
	}
}

func (x *Service) GetOrderDetail(
	/*req*/ ctx context.Context, id int) (
	/*res*/ order model.Order, httpCode int, err error,
) {
	order, err = x.Repo.GetOrderDetail(ctx, id)
	if err != nil {
		return order, http.StatusInternalServerError, err
	}

	return order, http.StatusOK, nil
}

func (x *Service) GetOrderByTitle(
	/*req*/ ctx context.Context, title string) (
	/*res*/ orders model.Orders, httpCode int, err error,
) {
	orders, err = x.Repo.GetOrderByTitle(ctx, title)
	if err != nil {
		return orders, http.StatusInternalServerError, err
	}

	return orders, http.StatusOK, nil
}

func (x *Service) GetAllOrders(
	/*req*/ ctx context.Context, request model.GetAllOrdersRequest) (
	/*res*/ orders model.Orders, httpCode int, err error,
) {
	orders, err = x.Repo.GetAllOrders(ctx, request)
	if err != nil {
		return orders, http.StatusInternalServerError, err
	}

	return orders, http.StatusOK, nil
}

func (x *Service) CreateOrder(
	/*req*/ ctx context.Context, request model.Order) (
	/*res*/ httpCode int, err error,
) {
	err = x.Repo.CreateOrder(ctx, request)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (x *Service) UpdateOrderStatus(
	/*req*/ ctx context.Context, request model.Order) (
	/*res*/ httpCode int, err error,
) {
	err = x.Repo.UpdateStatus(ctx, request)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (x *Service) DeleteOrder(
	/*req*/ ctx context.Context, id int) (
	/*res*/ httpCode int, err error,
) {
	err = x.Repo.DeleteOrder(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
