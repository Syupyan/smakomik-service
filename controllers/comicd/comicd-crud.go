package comicd

import (
	"path/filepath"
	"smaperpus-service/config"
	"smaperpus-service/helpers"
	comicbooks "smaperpus-service/models/comic-books"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Createc(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	// Mendapatkan daftar komik dari database
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("ID harus berupa angka")
	}

	var comicDetail comicbooks.Comic_details
	comicDetail.ComicId = id

	if err := c.BodyParser(&comicDetail); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	if err := config.ConnDb().Create(&comicDetail).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// return c.JSON(comicDetail)
	return c.Redirect("/comic/detail/" + idStr)
}

func Updatec(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	id := c.Params("id")
	var comicDetail comicbooks.Comic_details
	// Cari entri berdasarkan id
	if err := config.ConnDb().Where("comic_id = ?", id).First(&comicDetail).Error; err != nil {
		return err // Mengembalikan kesalahan jika gagal mencari entri
	}

	// Parsing input dari body request ke dalam struct comicDetail
	if err := c.BodyParser(&comicDetail); err != nil {
		return err // Mengembalikan kesalahan jika gagal membaca body request
	}

	// Lakukan pembaruan di database menggunakan model comicDetail
	if err := config.ConnDb().Model(&comicDetail).Updates(&comicDetail).Error; err != nil {
		return err // Mengembalikan kesalahan jika gagal melakukan pembaruan
	}

	return c.Redirect("/comic/detail/" + id)

}

func CreateChap(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	ids := c.Params("id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("ID harus berupa angka")
	}
	var comicChap comicbooks.Comic_subbtns
	comicChap.ComicID = id

	if err := c.BodyParser(&comicChap); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	if err := config.ConnDb().Create(&comicChap).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Redirect("/comic/detail/" + ids)

}

func Updatechap(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	id := c.Params("id")
	var comicChap comicbooks.Comic_subbtns
	// Cari entri berdasarkan id
	if err := config.ConnDb().First(&comicChap, id).Error; err != nil {
		return err
	}
	// Parsing input dari body request ke dalam struct comicChap
	if err := c.BodyParser(&comicChap); err != nil {
		return err // Mengembalikan kesalahan jika gagal membaca body request
	}

	config.ConnDb().Model(&comicChap).Updates(comicChap)

	return c.JSON(comicChap)

}

func Deletechap(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	id := c.Params("id")
	var comicChap comicbooks.Comic_subbtns
	if err := config.ConnDb().First(&comicChap, id).Error; err != nil {
		return err // Mengembalikan kesalahan jika gagal mengambil data dari database
	}
	// Ambil data dari database
	var subbtnsimg []comicbooks.Subbtns_images
	if err := config.ConnDb().Where("sub_id = ?", id).Find(&subbtnsimg).Error; err != nil {
		return err
	}

	// Hapus file gambar dari direktori assets
	for _, img := range subbtnsimg {
		if err := helpers.DeleteFileWithRetry("./assets/comic-image/comic-chapter/" + img.Images); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	}

	// Hapus entri dari database
	if err := config.ConnDb().Delete(&comicChap).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	referer := c.Get("Referer")
	// Mengarahkan pengguna kembali ke URL referer
	return c.Redirect(referer)
}

func Uploadimagebtns(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	ids := c.Params("id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("ID harus berupa angka")
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Parse each file from form
	files := form.File["file[]"]
	for _, file := range files {
		var subbtns comicbooks.Subbtns_images

		// Membuat nama file acak
		fileExt := filepath.Ext(file.Filename)
		randomFilename := uuid.New().String() + fileExt

		// Simpan file gambar yang baru
		if err := c.SaveFile(file, "./assets/comic-image/comic-chapter/"+randomFilename); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Gagal menyimpan file gambar baru",
			})
		}

		// Simpan informasi file ke dalam struct Subbtns_images
		subbtns.Images = randomFilename
		subbtns.SubId = id

		// Simpan informasi subbtns ke dalam database
		if err := config.ConnDb().Create(&subbtns).Error; err != nil {
			return err
		}
	}

	referer := c.Get("Referer")
	// Mengarahkan pengguna kembali ke URL referer
	return c.Redirect(referer)
}

func Updateimagebtns(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	id := c.Params("id")
	var subbtns comicbooks.Subbtns_images
	if err := config.ConnDb().First(&subbtns, id).Error; err != nil {
		return err
	}

	file, err := c.FormFile("file")
	if err == nil {
		// Membuat nama file acak
		fileExt := filepath.Ext(file.Filename)
		randomFilename := uuid.New().String() + fileExt

		// Simpan file gambar yang baru
		if err := c.SaveFile(file, "./assets/comic-image/comic-chapter/"+randomFilename); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Gagal menyimpan file gambar baru",
			})
		}

		// Hapus file gambar lama jika ada
		if err := helpers.DeleteFileWithRetry("./assets/comic-image/comic-chapter/" + subbtns.Images); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Update nama file gambar di dalam data komik
		subbtns.Images = randomFilename
	}

	c.BodyParser(&subbtns)
	config.ConnDb().Model(&subbtns).Updates(subbtns)
	referer := c.Get("Referer")
	// Mengarahkan pengguna kembali ke URL referer
	return c.Redirect(referer)

}

func Deleteimagebtns(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	id := c.Params("id")
	var subbtns comicbooks.Subbtns_images
	if err := config.ConnDb().First(&subbtns, id).Error; err != nil {
		return err // Mengembalikan kesalahan jika gagal mengambil data dari database
	}
	// Hapus entri dari database
	if err := helpers.DeleteFileWithRetry("./assets/comic-image/comic-chapter/" + subbtns.Images); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if err := config.ConnDb().Delete(&subbtns).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	referer := c.Get("Referer")
	// Mengarahkan pengguna kembali ke URL referer
	return c.Redirect(referer)
}
