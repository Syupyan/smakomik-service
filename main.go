package main

import (
	"smaperpus-service/config"
	"smaperpus-service/routers"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Mengkonfigurasi mesin template
	app := config.SetupTemplateConfig()
	app.Use(logger.New(logger.Config{
		Format: "${method} - ${path} - ${status}\n",
	}))

	// Menambahkan rute
	routers.SetupRoutes(app)
	// Menghubungkan ke database
	config.ConnDb()
	app.Listen(":8080")
}
