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
