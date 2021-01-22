package handler

import (
	"log"
	"strconv"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/xyn/golang_inventory/database"
	"golang.org/x/crypto/bcrypt"
)

func getHash(pwd string) string {
	pass := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	pwd := []byte(plainPwd)
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Login(c *fiber.Ctx) error {
	emailForm := c.FormValue("email")
	passwordForm := c.FormValue("password")
	var user database.User
	database.DB.First(&user, "email = ?", emailForm)

	if user.Email == "" {
		return c.JSON(fiber.Map{
			"error":   1,
			"message": "No such email found!",
		})
	} else if !comparePasswords(user.Password, passwordForm) {
		return c.JSON(fiber.Map{
			"error":   1,
			"message": "Password does not match!",
		})
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["firstName"] = user.FirstName
	claims["lastName"] = user.LastName
	claims["email"] = user.Email
	claims["accountType"] = user.AccountType
	//have token expire after 1 week
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix()
	log.Println(claims["exp"])

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func Register(c *fiber.Ctx) error {
	accountType, _ := strconv.Atoi(c.FormValue("accountType"))

	var user database.User
	database.DB.First(&user, "email = ?", c.FormValue("email"))

	if user.Email != "" {
		return c.JSON(fiber.Map{
			"error":   1,
			"message": "This email is already registered!",
		})
	}
	result := database.DB.Create(&database.User{
		FirstName:   c.FormValue("firstName"),
		LastName:    c.FormValue("lastName"),
		Email:       c.FormValue("email"),
		AccountType: accountType,
		Password:    getHash(c.FormValue("password")),
	})
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"error":   1,
			"message": "There was an error creating your account, please try again!",
		})
	}
	return c.JSON(fiber.Map{
		"error":   0,
		"message": "Account created successfully!",
	})
}
