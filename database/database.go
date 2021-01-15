package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("main.db"), &gorm.Config{})

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Device{})
	DB.AutoMigrate(&Delivery{})
	DB.AutoMigrate(&Donator{})

	if err != nil {
		panic("Can't connect to the db")
	}

}
