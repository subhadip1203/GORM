package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/subhadip1203/GORM/project1/database"
	"github.com/subhadip1203/GORM/project1/models"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	database.ConnectDB()
	models.LoadUserModel()
	models.AddUser()
}
