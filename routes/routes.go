package routes

import (
	"go-backend/handlers"
	"go-backend/utils"

	"github.com/gofiber/fiber/v2"
)

func Routes(handlers *handlers.Handlers) *fiber.App {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: utils.CustomErrMsg,
	})

	api := app.Group("/api")

	test := api.Group("/test")
	test.Get("/:id", handlers.GetTestById)
	test.Post("/", handlers.CreateTest)

	return app
}
