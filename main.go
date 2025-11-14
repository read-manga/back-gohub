package main

import (
	"microgo/core/usecase"
	"microgo/infrastructure/controller"
	"microgo/infrastructure/database"
	"microgo/infrastructure/repository"
	"microgo/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectPostgres()
	db := database.GetDB()

	userRepo := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := controller.NewAuthHandler(authUsecase)

	r := gin.Default()
	routes.SetupRoutes(r, authHandler)

	r.Run(":8080")
}
