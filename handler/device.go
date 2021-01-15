package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/xyn/golang_inventory/database"
)

func GetDevices(c *fiber.Ctx) error {
	var device database.Device
	result := database.DB.Preload("Donator").Preload("Delivery").Find(&device)
	log.Println(result.RowsAffected)
	return c.JSON(fiber.Map{
		"DevID":       device.DevID,
		"DevModel":    device.DevModel,
		"Accessories": device.Accessories,
		"Condition":   device.Condition,
		"Donator":     device.Donator,
		"Specs":       device.Specs,
		"WorkCond":    device.WorkCond,
		"Hostname":    device.Hostname,
		"RemoteName":  device.RemoteName,
		"Delivery":    device.Delivery,
		"Remarks":     device.Remarks,
	})
}

func GetDevice(c *fiber.Ctx) error {
	var device database.Device
	database.DB.Preload("Donator").Preload("Delivery").First(&device, "dev_id = ?", c.Params("id"))
	return c.JSON(fiber.Map{
		"DevID":       device.DevID,
		"DevModel":    device.DevModel,
		"Accessories": device.Accessories,
		"Condition":   device.Condition,
		"Donator":     device.Donator,
		"Specs":       device.Specs,
		"WorkCond":    device.WorkCond,
		"Hostname":    device.Hostname,
		"RemoteName":  device.RemoteName,
		"Delivery":    device.Delivery,
		"Remarks":     device.Remarks,
	})
}

func SetDevice(c *fiber.Ctx) error {
	var device database.Device
	database.DB.First(&device, "dev_id = ?", c.Params("id"))
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
