package comicb

import (
	"smaperpus-service/config"
	"smaperpus-service/helpers"
	comicbooks "smaperpus-service/models/comic-books"

	"github.com/gofiber/fiber/v2"
)

func V_Comic_Home(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}


	var comics []comicbooks.Comic_books // Pastikan Comic_books adalah model Anda

	// Lakukan query ke database untuk mengambil semua data komik
	if err := config.ConnDb().Order("id_comic DESC").Find(&comics).Error; err != nil {
		return err
	}

	// Render template dengan data komik yang ditemukan
	return c.Render("pages/comicb/v-comic", fiber.Map{
		"comics": comics,
	})
	
	// return c.JSON(fiber.Map{
	// 	"comics": comics,
	// })
}

func V_Comic_Create(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	return c.Render("pages/comicb/c-comic", fiber.Map{})
}

func V_Comic_Update(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	id := c.Params("id")
	var comic comicbooks.Comic_books // Menggunakan variabel tunggal untuk menyimpan data satu komik saja

	// Cari komik berdasarkan ID
	if err := config.ConnDb().Where("id_comic = ?", id).First(&comic).Error; err != nil {
		return err
	}

	// Lakukan apa pun yang diperlukan untuk mempersiapkan data sebelum ditampilkan di halaman update

	// Kembalikan data komik dalam format HTML
	return c.Render("pages/comicb/u-comic", fiber.Map{
		"comic": comic,
		"message": "berhasil", // Menyertakan data komik yang akan diperbarui dalam rendering template
	})
}

