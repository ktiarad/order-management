package middleware

import (
	"log"
	"net/http"
	"order-management/internal/helper"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helper.VerifyToken(ctx)
		log.Default().Println("Verify token :", verifyToken)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":           "UNAUTHENTICATED",
				"additional_info": err.Error(),
			})
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
