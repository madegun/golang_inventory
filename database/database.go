package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})

	DB.AutoMigrate(&User{})

	if err != nil {
		panic("Can't connect to the db")
	}

}

type User struct {
	gorm.Model
	FirstName   string `json:"firstname" bson: "firstname"`
	LastName    string `json:"lastname" bson: "lastname"`
	Email       string `gorm:"unique_index;nut null" json: "email"`
	AccountType int    `json:"accounttype" bson: "accounttype"`
	Password    string `json:"password" bson:"password"`
}
