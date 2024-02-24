package controllers

import (
	"smaperpus-service/helpers"

	"github.com/gofiber/fiber/v2"
)

func Dashboard(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}


	return c.Render("pages/dashboard", fiber.Map{})
}
