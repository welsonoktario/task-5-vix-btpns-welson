package router

import (
	"github.com/gin-gonic/gin"

	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/controllers"
	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/middlewares"
)

func InitRoutes() *gin.Engine {
	route := gin.Default()

	apiRoute := route.Group("/api")

	userRoute := apiRoute.Group("/users")
	{
		userRoute.POST("/login", controllers.Login)
		userRoute.POST("/register", controllers.Register)
	}

	photoRoute := apiRoute.Group("/photos", middlewares.APIMiddleware())
	{
		photoRoute.GET("/", controllers.AllPhotos)
		photoRoute.POST("/", controllers.AddPhoto)
		photoRoute.GET("/:photo", controllers.FindPhoto)
		photoRoute.PUT("/:photo", controllers.UpdatePhoto)
		photoRoute.DELETE("/:photo", controllers.DeletePhoto)
	}

	return route
}
