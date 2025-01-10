package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phn00dev/golang-news-app/internal/domain/category/constructor"
)

func AdminRoutes(app *fiber.App) {
	adminApiV1 := app.Group("v1/api/admin")

	// category routes
	categoryRoute := adminApiV1.Group("/categories")
	categoryRoute.Get("/", constructor.CategoryHandler.GetAll)
	categoryRoute.Get("/:categoryID", constructor.CategoryHandler.GetOne)
	categoryRoute.Post("/", constructor.CategoryHandler.Create)
	categoryRoute.Put("/:categoryID", constructor.CategoryHandler.Update)
	categoryRoute.Delete("/:categoryID", constructor.CategoryHandler.Delete)

}
