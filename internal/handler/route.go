package handler

import (
	"order-management/internal/middleware"
	"order-management/internal/service"

	"github.com/gin-gonic/gin"
)

type ServiceImpl struct {
	service.Services
}

func SetRoute(
	router *gin.Engine,
	service service.Services,
) {
	handler := ServiceImpl{
		Services: service,
	}

	// Authentication
	router.POST("signin", handler.SignIn())
	router.POST("login", handler.Login())

	// Customers
	router.GET("customers/detail/:id", middleware.Authentication(), handler.GetCustomerDetail())
	router.GET("customers/search/:name", middleware.Authentication(), handler.GetCustomerByName())
	router.GET("customers/list/:limit/:page", middleware.Authentication(), handler.GetAllCustomers())
	router.POST("customers", middleware.Authentication(), handler.CreateCustomer())
	router.PUT("customers", middleware.Authentication(), handler.UpdateCustomer())
	router.DELETE("customers/:id", middleware.Authentication(), handler.DeleteCustomer())

	// Orders
	router.GET("orders/detail/:id", middleware.Authentication(), handler.GetOrderDetail())
	router.GET("orders/search/:title", middleware.Authentication(), handler.GetOrderByTitle())
	router.GET("orders/list/:limit/:page", middleware.Authentication(), handler.GetAllOrders())
	router.POST("orders", middleware.Authentication(), handler.CreateOrder())
	router.PUT("orders/status", middleware.Authentication(), handler.UpdateOrderStatus())
	router.DELETE("orders/:id", middleware.Authentication(), handler.DeleteOrder())
}
