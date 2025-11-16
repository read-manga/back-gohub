
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


		v1.POST("/me", authHandler.UserFind)
		v1.PUT("/update", authHandler.UpdateUser)
		v1.PUT("/updateRepo", authHandler.UpdateRepository)
	
	}
}

