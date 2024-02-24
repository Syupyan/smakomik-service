// template.go
package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
)

// SetupTemplateConfig digunakan untuk mengkonfigurasi mesin template
func SetupTemplateConfig() *fiber.App {
	// Membuat mesin template
	// Create a new engine
	engine := mustache.New("./views", ".mustache")

	// Or from an embedded system
	//   Note that with an embedded system the partials included from template files must be
	//   specified relative to the filesystem's root, not the current working directory
	// engine := mustache.NewFileSystem(http.Dir("./views", ".mustache"), ".mustache")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	return app
}
