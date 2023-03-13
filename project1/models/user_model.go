package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/subhadip1203/GORM/project1/database"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name      string
	Age       int8      `gorm:"default:18"`
	CreatedAt time.Time `json:"-"`
}

func LoadUserModel() {
	db = database.GetDB()
	db.AutoMigrate(&User{})
}

func AddUser() {
	user := User{Name: "Jinzhu", Age: 18}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println("db error")
	}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Json error")
	}
	fmt.Println(string(jsonUser))

}
