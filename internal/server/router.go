package server

import (
	"go-learning/internal/domain/category"
	"go-learning/internal/domain/product"
	"go-learning/internal/helpers"
	"go-learning/internal/storage"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// init infra yang dibutuhkan handler
	response := helpers.NewResponse()
	minioClient := storage.NewMinioClient()

	// category
	categoryRepo := category.NewRepository()
	categoryService := category.NewService(categoryRepo, minioClient)
	categoryHandler := category.NewHandler(categoryService, response)

	// product
	productRepo := product.NewRepository()
	productService := product.NewService(productRepo)
	productHandler := product.NewHandler(productService, response)

	// welcome route
	app.Get("/", WelcomeHandler)

	// api group ke /api/v1/
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Category routes
	categoryGroup := v1.Group("/categories")
	categoryGroup.Post("/", categoryHandler.CreateCategory)

	// Product routes
	productGroup := v1.Group("/products")
	productGroup.Get("/", productHandler.ListProduct)
	productGroup.Get("/:id", productHandler.GetProduct)
	productGroup.Post("/", productHandler.CreateProduct)
	productGroup.Put("/:id", productHandler.UpdateProduct)
	productGroup.Delete("/:id", productHandler.DeleteProduct)
}
