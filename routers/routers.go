// router.go
package routers

import (
	"smaperpus-service/config"
	"smaperpus-service/controllers"
	"smaperpus-service/controllers/comicb"
	"smaperpus-service/controllers/comicd"
	"smaperpus-service/controllers/users"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes digunakan untuk menambahkan rute ke aplikasi
func SetupRoutes(app *fiber.App) {

	app.Static("/assets", "./assets")

	app.Get("/", users.V_Login)
	app.Post("/login", users.Login)
	app.Get("/logout", users.Logout)
	app.Get("/dashboard", config.AuthMiddlewareWithConfig(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}), controllers.Dashboard)

	method := app.Group("/method", config.AuthMiddlewareWithConfig(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	})) // /api

	// comicbook
	comicbook := method.Group("/comicbook")
	comicbook.Post("/create", comicb.Createc)
	comicbook.Post("/update/:id", comicb.Updatec)
	comicbook.Post("/delete/:id", comicb.Deletec)

	// comicbook view
	app.Get("/comic", config.AuthMiddlewareWithConfig(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}), comicb.V_Comic_Home)
	app.Get("/comic/create", config.AuthMiddlewareWithConfig(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}), comicb.V_Comic_Create)
	app.Get("/comic/update/:id", config.AuthMiddlewareWithConfig(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}), comicb.V_Comic_Update)

	// comicdetail
	comicdetail := method.Group("/comicdetail")
	comicdetail.Post("/create/:id", comicd.Createc)
	comicdetail.Post("/update/:id", comicd.Updatec)

	// comicdetail view
	app.Get("/comic/detail/:id", config.AuthMiddlewareWithConfig(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}), comicd.V_Comic_Detail)

	// comicchapter
	comicchapter := method.Group("/comicchapter")
	comicchapter.Post("/create/:id", comicd.CreateChap)
	comicchapter.Post("/update/:id", comicd.Updatechap)
	comicchapter.Post("/delete/:id", comicd.Deletechap)

	// chapterimage
	chapterimage := method.Group("/chapterimage")
	chapterimage.Post("/upload/:id", comicd.Uploadimagebtns)
	chapterimage.Post("/update/:id", comicd.Updateimagebtns)
	chapterimage.Post("/delete/:id", comicd.Deleteimagebtns)

	// chapterimage view
	app.Get("/comic/detail/images/:id", config.AuthMiddlewareWithConfig(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}), comicd.V_Comic_Image)

	// users
	app.Post("/createusers", users.Create_user)
	app.Post("/updateusers/:id", users.Update_user)
	app.Post("/deleteusers/:id", users.Delete_user)

	// users view
	app.Get("/users", config.AuthMiddlewareWithConfig(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}), users.V_Users)

	api := app.Group("/api", config.AuthMiddlewareWithConfig(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	app.Post("/verify", users.FreeUser)
	comicbookapi := api.Group("/comicbook", )
	comicbookapi.Get("/", comicb.GetComics)
	comicbookapi.Get("/detail/:id", comicd.GetComicDetails)
	comicbookapi.Get("/subbtns/:id", comicd.GetSubbtns)
	comicbookapi.Get("/subbtns/images/:id", comicd.GetImage)
	comicbookapi.Get("/new", comicb.GetNewComics)
	comicbookapi.Get("/populer", comicb.GetPopComics)

	// app.Use(config.AuthMiddleware)
	// app.Use(config.AuthMiddlewareWithConfig(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	// }))
	app.Get("/restricted", config.AuthMiddlewareWithConfig(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}), users.Restricted)

}
