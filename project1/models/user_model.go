package models

import (
	"github.com/subhadip1203/GORM/project1/database"
	"gorm.io/gorm"
)

func LoadUserModel() {
	db := database.GetDB()
	type User struct {
		gorm.Model
		Name string
		Age  int8 `gorm:"default:18"`
	}
	db.AutoMigrate(&User{})
}
