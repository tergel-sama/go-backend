package routes

import (
	"go-backend/handlers"
	"go-backend/utils"

	_ "go-backend/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"

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

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	return app
}
