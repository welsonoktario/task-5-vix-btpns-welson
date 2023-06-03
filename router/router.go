package router

import (
	"github.com/gin-gonic/gin"

	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/controllers"
)

func InitRoutes() *gin.Engine {
	route := gin.Default()

	userRoute := route.Group("/api/users")
	{
		userRoute.POST("/register", controllers.StoreUser)
	}

	return route
}
