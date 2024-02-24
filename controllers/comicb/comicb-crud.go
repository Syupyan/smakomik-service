package comicb

import (
	"path/filepath"
	"smaperpus-service/config"
	"smaperpus-service/helpers"
	comicbooks "smaperpus-service/models/comic-books"

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

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gagal mengunggah file",
		})
	}

	// Membuat nama file acak
	fileExt := filepath.Ext(file.Filename)
	randomFilename := uuid.New().String() + fileExt

	// Simpan file dengan nama acak
	if err := c.SaveFile(file, "./assets/comic-image/comic-cover/"+randomFilename); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan file",
		})
	}

	// Menggunakan nama file acak untuk menyimpan di database
	var comic comicbooks.Comic_books
	comic.Image = randomFilename // Setel nama file gambar yang baru di dalam struktur data Comic_books

	if err := c.BodyParser(&comic); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	if err := config.ConnDb().Create(&comic).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Redirect ke halaman lain setelah pembuatan berhasil
	return c.Redirect("/comic")
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
	var comic comicbooks.Comic_books
	if err := config.ConnDb().First(&comic, id).Error; err != nil {
		return err
	}

	file, err := c.FormFile("file")
	if err == nil {
		// Membuat nama file acak
		fileExt := filepath.Ext(file.Filename)
		randomFilename := uuid.New().String() + fileExt

		// Simpan file gambar yang baru
		if err := c.SaveFile(file, "./assets/comic-image/comic-cover/"+randomFilename); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Gagal menyimpan file gambar baru",
			})
		}

		// Hapus file gambar lama jika ada
		if comic.Image != "" {
			if err := helpers.DeleteFileWithRetry("./assets/comic-image/comic-cover/" + comic.Image); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
		}

		// Update nama file gambar di dalam data komik
		comic.Image = randomFilename
	}

	c.BodyParser(&comic)
	config.ConnDb().Model(&comic).Updates(comic)
	return c.Redirect("/comic", fiber.StatusFound)

}

func Deletec(c *fiber.Ctx) error {
	role, err := helpers.VerifyToken(c)
	if err != nil {
		return err
	}

	if role != "management" {
		return c.Status(401).SendString("Unauthorized")
	}

	id := c.Params("id")
	var comic comicbooks.Comic_books
	if err := config.ConnDb().First(&comic, id).Error; err != nil {
		return err // Mengembalikan kesalahan jika gagal mengambil data dari database
	}
	// Hapus file gambar jika entri telah dihapus dari database
	if err := helpers.DeleteFileWithRetry("./assets/comic-image/comic-cover/" + comic.Image); err != nil {
		return err // Mengembalikan kesalahan jika gagal menghapus file
	}

	var comicChap comicbooks.Comic_subbtns
	config.ConnDb().Where("comic_id = ?", id).First(&comicChap)

	var subbtnsimg []comicbooks.Subbtns_images
	config.ConnDb().Where("sub_id = ?", comicChap.IdSub).Find(&subbtnsimg)

	// Hapus file gambar dari direktori assets
	for _, img := range subbtnsimg {
		if err := helpers.DeleteFileWithRetry("./assets/comic-image/comic-chapter/" + img.Images); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	}

	// Hapus entri dari database
	if err := config.ConnDb().Delete(&comic).Error; err != nil {
		return err // Mengembalikan kesalahan jika gagal menghapus dari database
	}

	return c.Redirect("/comic", fiber.StatusFound)
}
