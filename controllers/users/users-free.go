package users

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func FreeUser(c *fiber.Ctx) error {
	// Menerima input dalam bentuk JSON
	type RequestBody struct {
		User string `json:"user"`
		Pass string `json:"pass"`
	}

	var reqBody RequestBody
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	// Throws Unauthorized error
	if reqBody.User != "kodeman" || reqBody.Pass != "mankode" {
		return c.SendString("error")
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name": "Free User",
		"role": "free",
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Set token as cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		Secure:   true,
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	})

	// Mengembalikan token sebagai respon JSON
	return c.JSON(fiber.Map{
		"token": t,
	})
}
