package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_sslmode := os.Getenv("DB_SSL")
	db_timezone := os.Getenv("DB_TIME_ZONE")

	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s  sslmode=%s TimeZone=%s", db_host, db_user, db_pass, db_name, db_port, db_sslmode, db_timezone)

	dbConnection, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dbURL,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatalf("Database not connected")
	} else {
		sqlDB, err := dbConnection.DB()
		if err != nil {
			log.Fatalf("Database connection failed")
		} else {
			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)
			fmt.Println("DB connected")
			DB = dbConnection
		}
	}
}

func GetDB() *gorm.DB {
	return DB
}
