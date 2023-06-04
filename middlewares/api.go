package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/helpers"
)

func APIMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "fail",
				"msg":    "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
