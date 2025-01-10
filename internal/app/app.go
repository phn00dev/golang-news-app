package app

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/phn00dev/golang-news-app/internal/setup/routes"
)

func NewApp(dependencies *Dependencies) (httpServer *fiber.App) {
	httpServer = fiber.New(fiber.Config{
		ServerHeader: dependencies.Config.HttpServer.Header,
		BodyLimit:    35 * 1024 * 1024,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return ctx.Status(code).JSON(fiber.Map{
				"status":  code,
				"message": "Näsazlyk ýüze çykdy, Sonrak synanysyn!!!",
			})
		},
		AppName: dependencies.Config.HttpServer.AppName,
	})

	// get routes
	routes.AdminRoutes(httpServer)
	routes.UserRoutes(httpServer)
	routes.StaticRoutes(httpServer)

	return httpServer
}
