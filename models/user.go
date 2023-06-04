package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id" valid:"optional"`
	Username  string    `gorm:"not null" json:"username"`
	Email     string    `gorm:"not null" json:"email" valid:"email~Please enter a valid email address"`
	Password  string    `json:"-" gorm:"not null; minlength" valid:"minstringlength(6)~Password length must greater or equal than 6"`
	Photos    []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos" valid:"optional"`
	CreatedAt time.Time `json:"-" valid:"optional"`
	UpdatedAt time.Time `json:"-" valid:"optional"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, validationError := govalidator.ValidateStruct(user)

	if validationError != nil {
		return validationError
	}

	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, validationError := govalidator.ValidateStruct(user)

	if validationError != nil {
		return validationError
	}

	return nil
}
