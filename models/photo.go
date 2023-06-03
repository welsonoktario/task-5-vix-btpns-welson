package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
	// User     User   `json:"user"`
}

func (photo *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, validationError := govalidator.ValidateStruct(photo)

	if validationError != nil {
		return validationError
	}

	return nil
}

func (photo *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, validationError := govalidator.ValidateStruct(photo)

	if validationError != nil {
		return validationError
	}

	return nil
}
