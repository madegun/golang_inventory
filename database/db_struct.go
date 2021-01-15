package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `gorm:"unique_index;nut null" json: "email"`
	AccountType int    `json:"accounttype"`
	Password    string `json:"password"`
}

type Donator struct {
	gorm.Model
	DeviceID    uint
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phonenumber"`
}

/* Status
0 - Delivered
1 - Scrapped
*/
type Delivery struct {
	gorm.Model
	DeviceID  uint
	Status    bool   `json:"status"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

/* WorkCond
0 - Working
1 - For parts
2 - Broken / for recycling
*/
type Device struct {
	gorm.Model
	DevID       string   `json:"devid"`
	DevModel    string   `json:"devmodel"`
	Accessories string   `json:"accessories"`
	Aspect      string   `json:"aspect"`
	Donator     Donator  `json:"donator"`
	Specs       string   `json:"specs"`
	WorkCond    int      `json:"workcond"`
	Hostname    string   `json:"hostname"`
	RemoteName  string   `json:"remotename"`
	Delivery    Delivery `json:"delivery"`
}
