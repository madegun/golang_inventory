package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xyn/golang_inventory/database"
)

func GetDevice(c *fiber.Ctx) error {
	var device database.Device
	database.DB.First(&device, "DevID = ?", c.Params("id"))
	return c.JSON(fiber.Map{"error": 0, "message": "N/A"})
}

func SetDevice(c *fiber.Ctx) error {
	var device database.Device
	database.DB.First(&device, "DevID = ?", c.Params("id"))
	if device.DevID != "" {
		return c.JSON(fiber.Map{"error": 1, "message": "This ID is already registered!"})
	}
	database.DB.Create(&database.Device{
		DevID:    c.Params("id"),
		DevModel: c.FormValue("devmodel"),
		Delivery: database.Delivery{
			Status: 0,
		},
		Donator: database.Donator{
			FirstName:   c.FormValue("donator_firstname"),
			LastName:    c.FormValue("donator_lastname"),
			PhoneNumber: c.FormValue("donator_phonenumber"),
		},
		Condition:   c.FormValue("condition"),
		Accessories: c.FormValue("accessories"),
		Remarks:     c.FormValue("remarks"),
		WorkCond:    3,
	})
	//TODO error handling
	return c.JSON(fiber.Map{"error": 0, "message": "Device created successfully!"})
}
