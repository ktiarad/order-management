package handler

import (
	"net/http"
	"order-management/internal/model"

	"github.com/gin-gonic/gin"
)

func (x ServiceImpl) SignIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request model.User

		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		httpCode, err := x.AuthenticationServices.SignIn(ctx, request)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, nil)
	}
}

func (x ServiceImpl) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request model.User

		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		token, httpCode, err := x.AuthenticationServices.LogIn(ctx, request)
		if err != nil {
			ctx.JSON(httpCode, err)
			return
		}

		ctx.JSON(httpCode, token)
	}
}
