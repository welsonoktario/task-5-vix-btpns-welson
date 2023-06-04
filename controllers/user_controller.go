package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/database"
	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/helpers"
	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/models"
)

type RegisterJsonBody struct {
	Username string
	Email    string
	Password string
}

func Register(c *gin.Context) {
	DB := database.GetDB()
	var body RegisterJsonBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(200, gin.H{
			"status": "fail",
			"msg":    "Please fill the register form",
		})
		return
	}

	hashedPassword, err := helpers.HashPassword(body.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Server error"})
		return
	}

	user := &models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: hashedPassword,
	}

	fmt.Printf("%+v\n", user)

	dbErr := DB.Create(&user).Error

	if dbErr != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status": "fail",
			"msg":    dbErr,
		})
		return
	}

	token, err := helpers.GenerateToken(user.ID)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status": "fail",
			"msg":    err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data": gin.H{
			"user":  &user,
			"token": token,
		},
	})
}

type LoginJsonBody struct {
	Email    string
	Password string
}

func Login(c *gin.Context) {
	DB := database.GetDB()
	var user models.User
	var body LoginJsonBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(200, gin.H{
			"status": "fail",
			"msg":    "Please input email or password",
		})
		return
	}

	res := DB.First(&user, "email = ?", body.Email)

	if res.RowsAffected == 0 {
		c.JSON(200, gin.H{
			"status": "fail",
			"msg":    "Email or password is invalid",
		})
		return
	}

	checkPassword := helpers.CheckPasswordHash(body.Password, user.Password)

	if !checkPassword {
		c.JSON(200, gin.H{
			"status": "fail",
			"msg":    "Email or password is invalid",
		})
		return
	}

	token, err := helpers.GenerateToken(user.ID)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status": "fail",
			"msg":    err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data": gin.H{
			"user":  &user,
			"token": token,
		},
	})
}
