package users

import (
	"smaperpus-service/config"
	"smaperpus-service/helpers"
	users "smaperpus-service/models/users"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	username := c.FormValue("user")
	password := c.FormValue("pass")

	// Find user by username
	// Find user by username
	var user users.Users
	result := config.ConnDb().Where("username = ?", username).First(&user)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Check if user is found
	if user.IdUsers == 0 { // Assuming IdUsers is the primary key
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Compare password with hashed password from database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name": user.Username,
		"role": "management",
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
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
		Name:  "token",
		Value: t,
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Strict",
		Expires:  time.Now().Add(time.Hour * 24),
	})

	return c.Redirect("/dashboard")
}

func Restricted(c *fiber.Ctx) error {
	// Ambil token dari variabel lokal
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	if err != nil {
		return err
	}

	// Kirim pesan selamat datang
	return c.SendString("Welcome " + " to the restricted area!")
}

func Logout(c *fiber.Ctx) error {
	// Hapus cookie yang berisi token
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour), // Atur waktu kedaluwarsa ke belakang untuk menghapus cookie
	})

	// Kirim respons logout berhasil
	return c.Redirect("/")
}

func Create_user(c *fiber.Ctx) error {
	// Parse request body
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	var newUser users.Users
	if err := c.BodyParser(&newUser); err != nil {
		return err
	}

	// Generate hashed password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// Create the user in the database
	user := users.Users{Username: newUser.Username, Password: string(hashedPassword)}
	result := config.ConnDb().Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": result.Error.Error()})
	}

	referer := c.Get("Referer")
	// Mengarahkan pengguna kembali ke URL referer
	return c.Redirect(referer)
}

func Update_user(c *fiber.Ctx) error {
	// Parse request body
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	var updatedUser users.Users
	if err := c.BodyParser(&updatedUser); err != nil {
		return err
	}

	// Get user ID from URL params
	userID := c.Params("id")

	// Find the user in the database
	existingUser := users.Users{}
	if err := config.ConnDb().Where("id_users = ?", userID).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Update user fields
	existingUser.Username = updatedUser.Username

	// Check if the password is updated
	if updatedUser.Password != "" {
		// Generate hashed password using bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
		}
		existingUser.Password = string(hashedPassword)
	}

	// Save changes to the database
	result := config.ConnDb().Save(&existingUser)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	referer := c.Get("Referer")
	// Mengarahkan pengguna kembali ke URL referer
	return c.Redirect(referer)
}

func Delete_user(c *fiber.Ctx) error {
	// Get user ID from URL params
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	userID := c.Params("id")

	// Find the user in the database
	existingUser := users.Users{}
	if err := config.ConnDb().Where("id_users = ?", userID).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Delete the user from the database
	result := config.ConnDb().Delete(&existingUser)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	referer := c.Get("Referer")
	// Mengarahkan pengguna kembali ke URL referer
	return c.Redirect(referer)
}
