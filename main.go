package main

import (
	"net/http"
	"order-management/internal/handler"
	"order-management/internal/repository"
	"order-management/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := repository.ConnectDB()
	defer db.Close()

	repo := repository.NewRepo(db)

	router := gin.Default()

	services := service.NewService(repo)
	handler.SetRoute(router, services)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
