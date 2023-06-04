package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/database"
	"github.com/welsonoktario/task-5-vix-btpns-welsonoktario/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// govalidator.SetFieldsRequiredByDefault(true)

	database.ConnectDatabase()
	database.MigrateDatabase()

	router := router.InitRoutes()
	router.MaxMultipartMemory = 8 << 20
	router.Static("/assets/photos", "./storage/photos")

	router.Run(":8000")
}
