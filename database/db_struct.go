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
0 - Received
1 - In work
2 - Delivered
3 - Scrapped
*/
type Delivery struct {
	gorm.Model
	DeviceID  uint
	Status    int    `json:"status"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

/* WorkCond
0 - Working
1 - For parts
2 - Broken / for recycling
3 - To be tested
*/
type Device struct {
	gorm.Model
	DevID       string   `json:"devid"`
	DevModel    string   `json:"devmodel"`
	Accessories string   `json:"accessories"`
	Condition   string   `json:"aspect"`
	Donator     Donator  `json:"donator"`
	Specs       string   `json:"specs"`
	WorkCond    int      `json:"workcond"`
	Hostname    string   `json:"hostname"`
	RemoteName  string   `json:"remotename"`
	Delivery    Delivery `json:"delivery"`
	Remarks     string   `json:"remarks"`
}
