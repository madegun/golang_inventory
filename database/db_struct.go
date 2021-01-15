package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string `json:"firstname" bson: "firstname"`
	LastName    string `json:"lastname" bson: "lastname"`
	Email       string `gorm:"unique_index;nut null" json: "email"`
	AccountType int    `json:"accounttype" bson: "accounttype"`
	Password    string `json:"password" bson:"password"`
}
