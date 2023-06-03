package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/database"
	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/helpers"
	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/models"
)

func StoreUser(c *gin.Context) {
	DB := database.GetDB()
	hashedPassword, err := helpers.HashPassword(c.Param("password"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Server error"})
	}

	user := &models.User{
		Username: c.Param("username"),
		Email:    c.Param("email"),
		Password: hashedPassword,
	}

	res := DB.Create(&user)

	if res.Error != nil {
		c.AbortWithStatusJSON(500, res.Error)
	}

	c.JSON(http.StatusCreated, &user)
}
