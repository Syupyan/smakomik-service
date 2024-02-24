package users

import (
	"smaperpus-service/config"
	"smaperpus-service/helpers"
	"smaperpus-service/models/users"

	"github.com/gofiber/fiber/v2"
)

func V_Login(c *fiber.Ctx) error {

	cookie := c.Cookies("token")
	if cookie == "" {
		return c.Render("pages/users/v-login", fiber.Map{})
	}
	return c.Redirect("/dashboard")

}

func V_Users(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	var users []users.Users

	if err := config.ConnDb().Find(&users).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Render("pages/users/v-users", fiber.Map{
		"users": users,
	})
	// return c.JSON(users)
}

func C_Users(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	return c.Render("pages/users/c-users", fiber.Map{})
}

func U_Users(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	return c.Render("pages/users/u-users", fiber.Map{})
}
