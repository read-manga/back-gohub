package routes

import (
	"microgo/infrastructure/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authHandler *controller.AuthHandler) {
	v1 := r.Group("/auth")
	{
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)
		v1.PUT("/update", authHandler.UpdateUser)
	}
}
