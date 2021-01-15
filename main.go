package main

import (
	jwt "github.com/form3tech-oss/jwt-go"
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
	app.Get("/", root)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	app.Get("/dashboard", dashboard)

	app.Listen(":3000")
}

func root(c *fiber.Ctx) error {
	if c.Get("Authorization") != "" {
		c.Redirect("/dashboard")
	}
	return c.SendString("*should do stuff yet but it doesnt lol*")
}

func dashboard(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["email"].(string)
	return c.SendString("Welcome " + claims["firstName"].(string) + name)
}
