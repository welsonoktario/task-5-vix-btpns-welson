package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"not null" json:"username"`
	Email     string    `gorm:"not null" json:"email" valid:"email~Please enter a valid email address"`
	Password  string    `gorm:"not null; minlength" valid:"minstringlength(6)~Password length must greater or equal than 6"`
	Photos    []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADEE;" json:"photos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
