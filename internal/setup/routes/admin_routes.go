package routes

import (
	"github.com/gofiber/fiber/v2"
	adminConstructor "github.com/phn00dev/golang-news-app/internal/domain/admin/constructor"
	categoryConstructor "github.com/phn00dev/golang-news-app/internal/domain/category/constructor"
	postConstructor "github.com/phn00dev/golang-news-app/internal/domain/post/constructor"
)

func AdminRoutes(app *fiber.App) {
	adminApiV1 := app.Group("v1/api/admin")

	// admin routes
	adminRoute := adminApiV1.Group("/admins")
	adminRoute.Get("/", adminConstructor.AdminHandler.GetAll)
	adminRoute.Get("/:adminID", adminConstructor.AdminHandler.GetOne)
	adminRoute.Post("/", adminConstructor.AdminHandler.Create)
	adminRoute.Put("/:adminID", adminConstructor.AdminHandler.Update)
	adminRoute.Delete("/:adminID", adminConstructor.AdminHandler.Delete)

	// category routes
	categoryRoute := adminApiV1.Group("/categories")
	categoryRoute.Get("/", categoryConstructor.CategoryHandler.GetAll)
	categoryRoute.Get("/:categoryID", categoryConstructor.CategoryHandler.GetOne)
	categoryRoute.Post("/", categoryConstructor.CategoryHandler.Create)
	categoryRoute.Put("/:categoryID", categoryConstructor.CategoryHandler.Update)
	categoryRoute.Delete("/:categoryID", categoryConstructor.CategoryHandler.Delete)

	// post routes
	postRoute := adminApiV1.Group("/posts")
	postRoute.Get("/", postConstructor.PostHandler.GetAll)
	postRoute.Get("/:postID", postConstructor.PostHandler.GetOne)
	postRoute.Post("/", postConstructor.PostHandler.Create)
	postRoute.Put("/:postID", postConstructor.PostHandler.Update)
	postRoute.Delete("/:postID", postConstructor.PostHandler.Delete)

}
