package routes

import "github.com/gofiber/fiber/v2"

func registerProductRoutes(v1 fiber.Router, deps *Dependencies) {
	product := v1.Group("/products")

	product.Get("/", deps.ProductController.ListProduct)
	product.Get("/:id", deps.ProductController.GetProduct)
	product.Post("/", deps.ProductController.CreateProduct)
	product.Put("/:id", deps.ProductController.UpdateProduct)
	product.Delete("/:id", deps.ProductController.DeleteProduct)
}
