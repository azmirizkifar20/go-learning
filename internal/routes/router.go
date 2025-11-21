package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App, deps *Dependencies) {
	// set prefix path
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// register API
	registerCategoryRoutes(v1, deps)
	registerProductRoutes(v1, deps)
}
