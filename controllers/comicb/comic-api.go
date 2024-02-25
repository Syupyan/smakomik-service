package comicb

import (
	"smaperpus-service/config"
	"smaperpus-service/helpers"
	comicbooks "smaperpus-service/models/comic-books"

	"github.com/gofiber/fiber/v2"
)

func GetComics(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "free" {
		return c.Status(401).SendString("Unauthorized")
	}
	var comics []comicbooks.Comic_books
	if err := config.ConnDb().Order("id_comic DESC").Find(&comics).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(comics)
}

func GetNewComics(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "free" {
		return c.Status(401).SendString("Unauthorized")
	}
	var newcomics []comicbooks.Comic_books
	if err := config.ConnDb().Order("id_comic DESC").Limit(4).Find(&newcomics).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(newcomics)
}

func GetPopComics(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "free" {
		return c.Status(401).SendString("Unauthorized")
	}
	var newcomics []comicbooks.Comic_books
	if err := config.ConnDb().Order("access_u DESC").Limit(8).Find(&newcomics).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(newcomics)
}

// func GetSubbtns(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var subbtns []comicbooks.Comic_subbtns
// 	if err := config.ConnDb().Where("comic_id = ?", id).Find(&subbtns).Error; err != nil {
// 		return err
// 	}

// 	return c.Status(fiber.StatusOK).JSON(subbtns)
// }

