package comicd

import (
	"smaperpus-service/config"
	"smaperpus-service/helpers"
	comicbooks "smaperpus-service/models/comic-books"

	"github.com/gofiber/fiber/v2"
)

func V_Comic_Detail(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}


	id := c.Params("id")

	var comic comicbooks.Comic_books
	if err := config.ConnDb().Where("id_comic = ?", id).First(&comic).Error; err != nil {
		return err
	}

	var comicChap []comicbooks.Comic_subbtns
	if err := config.ConnDb().Where("comic_id = ?", id).Find(&comicChap).Error; err != nil {
		// data := fiber.Map{
		//     "comic":       comic,
		// }

		// // Render halaman HTML dengan data yang telah disiapkan
		// return c.Render("pages/comicd/v-comicd", data)
		return c.JSON(comicChap)
	}

	var comicDetails comicbooks.Comic_details
	config.ConnDb().Where("comic_id = ?", id).First(&comicDetails)

	// Masukkan semua data ke dalam satu map
	data := fiber.Map{
		"comic":       comic,
		"comicDetail": comicDetails,
		"comicChap":   comicChap,
	}

	// Render halaman HTML dengan data yang telah disiapkan
	return c.Render("pages/comicd/v-comicd", data)
}

func V_Comic_Image(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}


	id := c.Params("id")
	var subbtns []comicbooks.Subbtns_images

	if err := config.ConnDb().Where("sub_id = ?", id).Find(&subbtns).Error; err != nil {
		return err
	}

	var comicsbbtns comicbooks.Comic_subbtns
	if err := config.ConnDb().First(&comicsbbtns, id).Error; err != nil {
		return err
	}

	return c.Render("pages/comicd/v-comicd-image", fiber.Map{
		"subbtns": subbtns,
		"id":      id,
		"comicid":    comicsbbtns.ComicID,
	})
}
