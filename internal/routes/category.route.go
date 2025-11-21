package routes

import "github.com/gofiber/fiber/v2"

func registerCategoryRoutes(v1 fiber.Router, deps *Dependencies) {
	category := v1.Group("/categories")

	category.Post("/", deps.CategoryController.CreateCategory)
}
