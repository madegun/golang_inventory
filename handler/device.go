package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/xyn/golang_inventory/database"
	"gorm.io/gorm"
)

func GetDevices(c *fiber.Ctx) error {
	var device []database.Device
	result := database.DB.Preload("Donator").Preload("Delivery").Find(&device)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(fiber.Map{
			"error":   1,
			"message": "No devices registered!",
		})
	}
	return c.JSON(device)
}

func GetDevice(c *fiber.Ctx) error {
	var device database.Device
	result := database.DB.Preload("Donator").Preload("Delivery").First(&device, "dev_id = ?", c.Params("id"))
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(fiber.Map{
			"error":   1,
			"message": "This ID is not registered!",
		})
	}
	return c.JSON(device)
}

func SetDevice(c *fiber.Ctx) error {
	var device database.Device
	database.DB.First(&device, "dev_id = ?", c.Params("id"))
	if device.DevID != "" {
		return c.JSON(fiber.Map{
			"error":   1,
			"message": "This ID is already registered!",
		})
	}
	result := database.DB.Create(&database.Device{
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
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"error":   1,
			"message": "There was an error creating this device, please try again!",
		})
	}
	return c.JSON(fiber.Map{
		"error":   0,
		"message": "Device created successfully!",
	})
}

func ModifyDevice(c *fiber.Ctx) error {
	workCondition, _ := strconv.Atoi(c.FormValue("workcond"))
	deliveryStatus, _ := strconv.Atoi(c.FormValue("delivery_status"))
	//TODO error handling
	var device database.Device
	result := database.DB.Preload("Donator").Preload("Delivery").Model(&device).Where("dev_id = ?", c.Params("id")).Updates(database.Device{
		DevModel:    c.FormValue("devmodel"),
		Accessories: c.FormValue("accessories"),
		Condition:   c.FormValue("condition"),
		Donator: database.Donator{
			FirstName:   c.FormValue("donator_firstname"),
			LastName:    c.FormValue("donator_lastname"),
			PhoneNumber: c.FormValue("donator_phonenumber"),
		},
		Specs:      c.FormValue("specs"),
		WorkCond:   workCondition,
		Hostname:   c.FormValue("hostname"),
		RemoteName: c.FormValue("remotename"),
		Delivery: database.Delivery{
			Status:    deliveryStatus,
			FirstName: c.FormValue("delivery_firstname"),
			LastName:  c.FormValue("delivery_lastname"),
			City:      c.FormValue("delivery_city"),
		},
		Remarks: c.FormValue("remarks"),
	})
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"error":   1,
			"message": "There was an error modifying the device, please try again!",
		})
	}
	return c.JSON(fiber.Map{
		"error":   0,
		"message": "Device modified successfully!",
	})
}
