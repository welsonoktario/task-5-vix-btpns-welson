package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/models"
)

var (
	DB  *gorm.DB
	err error
)

func GetDB() *gorm.DB {
	return DB
}

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbDSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	DB, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
}

func MigrateDatabase() {
	err = DB.AutoMigrate(&models.User{}, &models.Photo{})

	if err != nil {
		return
	}
}
