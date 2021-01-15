package main

import (
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/xyn/golang_inventory/database"
)

func getHash(pwd string) string {
	pass := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd string) bool { // Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	pwd := []byte(plainPwd)
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func main() {
	app := fiber.New()
	database.Connect()

	app.Post("/login", login)
	app.Post("/register", register)
	app.Get("/", accessible)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	app.Get("/dashboard", dashboard)

	app.Listen(":3000")
}

func login(c *fiber.Ctx) error {
	emailForm := c.FormValue("email")
	passwordForm := c.FormValue("password")
	var user database.User
	database.DB.First(&user, "email = ?", emailForm)

	if user.Email == "" {
		c.JSON(fiber.Map{"error": 1, "message": "No such email found!"})
		return c.SendStatus(fiber.StatusUnauthorized)
	} else if !comparePasswords(user.Password, passwordForm) {
		c.JSON(fiber.Map{"error": 1, "message": "Password does not match!"})
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func register(c *fiber.Ctx) error {
	firstNameForm := c.FormValue("firstName")
	lastNameForm := c.FormValue("lastName")
	emailForm := c.FormValue("email")
	passwordForm := c.FormValue("password")
	accountType, _ := strconv.Atoi(c.FormValue("accountType"))

	var user database.User
	database.DB.First(&user, "email = ?", emailForm)

	if user.Email != "" {
		c.JSON(fiber.Map{"error": 1, "message": "This email is already registered!"})
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	database.DB.Create(&database.User{
		FirstName:   firstNameForm,
		LastName:    lastNameForm,
		Email:       emailForm,
		AccountType: accountType,
		Password:    getHash(passwordForm),
	})
	return c.JSON(fiber.Map{"error": 0, "message": "Account created successfully!"})
}

func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func dashboard(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["email"].(string)
	return c.SendString("Welcome " + name)
}
