package main

import (
	"microgo/core/handler"
	"microgo/core/infra/database"
	"microgo/core/repository"
	"microgo/core/routes"
	"microgo/core/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectPostgres()
	db := database.GetDB()

	userRepo := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	r := gin.Default()
	routes.SetupRoutes(r, authHandler)

	r.Run(":8080")
}
