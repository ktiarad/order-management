package handler

import (
	"errors"
	"net/http"
	"order-management/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (x ServiceImpl) GetCustomerDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("id should not empty"))
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, err)
		}

		results, httpCode, err := x.CustomerServices.GetCustomerDetail(ctx, idInt)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, results)
	}
}

func (x ServiceImpl) GetCustomerByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		if name == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("name should not empty"))
			return
		}

		results, httpCode, err := x.CustomerServices.GetCustomerByName(ctx, name)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, results)
	}
}

func (x ServiceImpl) GetAllCustomers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limit := ctx.Param("limit")
		page := ctx.Param("page")

		if limit == "" {
			limit = "0"
		}
		if page == "" {
			page = "0"
		}

		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New(("invalid limit")))
		}

		pageInt, err := strconv.Atoi(page)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New(("invalid page")))
		}

		request := model.GetAllCustomersRequest{
			Limit: limitInt,
			Page:  pageInt,
		}

		results, httpCode, err := x.CustomerServices.GetAllCustomers(ctx, request)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, results)
	}
}

func (x ServiceImpl) CreateCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request model.CreateCustomerRequest

		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		httpCode, err := x.CustomerServices.CreateCustomer(ctx, request)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, nil)
	}
}

func (x ServiceImpl) UpdateCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request model.Customer

		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		httpCode, err := x.CustomerServices.UpdateCustomer(ctx, request)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, nil)
	}
}

func (x ServiceImpl) DeleteCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("id should not empty"))
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, err)
		}

		httpCode, err := x.CustomerServices.DeleteCustomer(ctx, idInt)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, nil)
	}
}
