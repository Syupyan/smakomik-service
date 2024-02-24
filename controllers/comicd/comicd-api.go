package comicd

import (
	"smaperpus-service/config"
	comicbooks "smaperpus-service/models/comic-books"

	"github.com/gofiber/fiber/v2"
)

func GetComicDetails(c *fiber.Ctx) error {
	id := c.Params("id")

	// Ambil detail komik dari database berdasarkan ID
	var comicdetail comicbooks.Comic_details
	if err := config.ConnDb().Where("comic_id = ?", id).Find(&comicdetail).Error; err != nil {
		return err
	}

	// Ambil data komik dari database berdasarkan ID
	var comicbooks comicbooks.Comic_books
	if err := config.ConnDb().First(&comicbooks, id).Error; err != nil {
		return err
	}

	// Tambahkan 1 pada AccessU
	comicbooks.AccessU++

	// Simpan perubahan kembali ke database
	if err := config.ConnDb().Save(&comicbooks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal memperbarui AccessU")
	}

	// Kembalikan detail komik sebagai respons JSON
	return c.Status(fiber.StatusOK).JSON(comicdetail)
}

func GetSubbtns(c *fiber.Ctx) error {
	id := c.Params("id")
	var subbtns []comicbooks.Comic_subbtns
	if err := config.ConnDb().Where("comic_id = ?", id).Find(&subbtns).Error; err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(subbtns)
}

func GetImage(c *fiber.Ctx) error {
	id := c.Params("id")
	var imagesbbtns []comicbooks.Subbtns_images
	if err := config.ConnDb().Where("sub_id = ?", id).Find(&imagesbbtns).Error; err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(imagesbbtns)
}
