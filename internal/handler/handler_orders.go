package handler

import (
	"errors"
	"net/http"
	"order-management/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (x ServiceImpl) GetOrderDetail() gin.HandlerFunc {
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

		results, httpCode, err := x.OrderServices.GetOrderDetail(ctx, idInt)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, results)
	}
}

func (x ServiceImpl) GetOrderByTitle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := ctx.Param("title")

		if title == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("title should not empty"))
			return
		}

		results, httpCode, err := x.OrderServices.GetOrderByTitle(ctx, title)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, results)
	}
}

func (x ServiceImpl) GetAllOrders() gin.HandlerFunc {
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

		request := model.GetAllOrdersRequest{
			Limit: limitInt,
			Page:  pageInt,
		}

		results, httpCode, err := x.OrderServices.GetAllOrders(ctx, request)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, results)
	}
}

func (x ServiceImpl) CreateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request model.Order

		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		httpCode, err := x.OrderServices.CreateOrder(ctx, request)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, nil)
	}
}

func (x ServiceImpl) UpdateOrderStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request model.Order

		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		httpCode, err := x.OrderServices.UpdateOrderStatus(ctx, request)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, nil)
	}
}

func (x ServiceImpl) DeleteOrder() gin.HandlerFunc {
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

		httpCode, err := x.OrderServices.DeleteOrder(ctx, idInt)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, nil)
	}
}
