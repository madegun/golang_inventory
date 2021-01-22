package main

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/xyn/golang_inventory/database"
	"github.com/xyn/golang_inventory/handler"
)

func main() {
	app := fiber.New()
	database.Connect()

	app.Post("/api/login", handler.Login)
	app.Post("/api/register", handler.Register)

	app.Static("/assets", "./frontend/assets")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./frontend/authentication/index.html")
	})

	app.Get("/dashboard/device/:id", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./frontend/dashboard/device.html")
	})

	app.Static("/dashboard", "./frontend/dashboard")

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	app.Get("/api/device/:id", handler.GetDevice)
	app.Get("/api/devices", handler.GetDevices)
	app.Post("/api/device/:id", handler.SetDevice)
	app.Post("/api/device/modify/:id", handler.ModifyDevice)

	app.Listen(":3000")
}
