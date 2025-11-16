package main

import (
	"microgo/core/usecase"
	"microgo/infrastructure/controller"
	"microgo/infrastructure/database"
	"microgo/infrastructure/repository"
	"microgo/infrastructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	database.ConnectPostgres()
	db := database.GetDB()

	userRepo := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := controller.NewAuthHandler(authUsecase)

	r := gin.Default()
  	r.Use(cors.New(cors.Config{
        	AllowOrigins:     []string{"http://localhost:8080/","http://localhost:8081/", "*"},
        	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        	ExposeHeaders:    []string{"Content-Length"},
        	AllowCredentials: true,
        	MaxAge: 12 * time.Hour,
    	}))
	routes.SetupRoutes(r, authHandler)

	r.Run(":8081")
}
