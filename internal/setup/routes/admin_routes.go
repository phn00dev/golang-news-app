package routes

import (
	"github.com/gofiber/fiber/v2"
	categoryConstructor "github.com/phn00dev/golang-news-app/internal/domain/category/constructor"
	"github.com/phn00dev/golang-news-app/internal/domain/post/constructor"
)

func AdminRoutes(app *fiber.App) {
	adminApiV1 := app.Group("v1/api/admin")

	// category routes
	categoryRoute := adminApiV1.Group("/categories")
	categoryRoute.Get("/", categoryConstructor.CategoryHandler.GetAll)
	categoryRoute.Get("/:categoryID", categoryConstructor.CategoryHandler.GetOne)
	categoryRoute.Post("/", categoryConstructor.CategoryHandler.Create)
	categoryRoute.Put("/:categoryID", categoryConstructor.CategoryHandler.Update)
	categoryRoute.Delete("/:categoryID", categoryConstructor.CategoryHandler.Delete)

	// post routes
	postRoute := adminApiV1.Group("/posts")
	postRoute.Get("/", constructor.PostHandler.GetAll)
	postRoute.Get("/:postID", constructor.PostHandler.GetOne)
	postRoute.Post("/", constructor.PostHandler.Create)
	postRoute.Put("/:postID", constructor.PostHandler.Update)
	postRoute.Delete("/:postID", constructor.PostHandler.Delete)

}
