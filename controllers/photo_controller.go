package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/database"
	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/helpers"
	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/models"
)

func AllPhotos(c *gin.Context) {
	DB := database.GetDB()
	var photos []models.Photo
	userID, err := helpers.ExtractTokenID(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    "Unauthorized",
		})
		return
	}

	dbErr := DB.Find(&photos).Where("user_id = ?", userID).Error

	if dbErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    dbErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   &photos,
	})
}

func FindPhoto(c *gin.Context) {
	DB := database.GetDB()
	id := c.Param("photo")
	var photo models.Photo

	err := DB.First(&photo, id).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   &photo,
	})
}

func AddPhoto(c *gin.Context) {
	DB := database.GetDB()
	userID, err := helpers.ExtractTokenID(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    "Unauthorized",
		})
		return
	}

	file, err := c.FormFile("photo_file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    "Please provide photo file",
			"error":  err.Error(),
		})
		return
	}

	filename := "/storage/photos/" + uuid.NewString() + filepath.Ext(file.Filename)
	destination := filepath.Dir(filename)

	errFile := c.SaveUploadedFile(file, destination)

	if errFile != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    "Server error",
			"error":  errFile.Error,
		})
		return
	}

	photo := models.Photo{
		Title:    c.PostForm("title"),
		Caption:  c.PostForm("caption"),
		PhotoUrl: destination,
		UserID:   userID,
	}

	dbErr := DB.Create(&photo).Error

	if dbErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    "Server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data":   &photo,
	})
}

type UpdatePhotoJsonBody struct {
	Title   string
	Caption string
}

func UpdatePhoto(c *gin.Context) {
	DB := database.GetDB()
	id := c.Params.ByName("photo")
	var photo models.Photo
	var body UpdatePhotoJsonBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(200, gin.H{
			"status": "fail",
			"msg":    "Please input email or password",
		})
		return
	}

	userID, err := helpers.ExtractTokenID(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
			"msg":    "Unauthorized",
		})
		return
	}

	dbErr := DB.First(&photo, id).Where("user_id = ?", userID).Error

	if dbErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    "Photo not found",
		})
		return
	}

	photo.Title = body.Title
	photo.Caption = body.Caption

	updateErr := DB.Save(&photo).Error

	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    "Server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func DeletePhoto(c *gin.Context) {
	DB := database.GetDB()
	id := c.Params.ByName("photo")
	userID, err := helpers.ExtractTokenID(c)
	var photo models.Photo

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
			"msg":    "Unauthorized",
		})
		return
	}

	dbErr := DB.First(&photo, id).Where("user_id = ?", userID).Error

	if dbErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    "Photo not found",
		})
		return
	}

	deleteErr := DB.Delete(&photo).Error

	if deleteErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    "Server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
