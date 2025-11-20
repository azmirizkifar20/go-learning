package server

import (
	"go-learning/internal/controllers"
	"go-learning/internal/helpers"
	"go-learning/internal/repositories"
	"go-learning/internal/services"
	"go-learning/internal/storage"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// init infra yang dibutuhkan handler
	response := helpers.NewResponse()
	minioClient := storage.NewMinioClient()

	// category
	categoryRepo := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepo, minioClient)
	categoryHandler := controllers.NewCategoryController(categoryService, response)

	// product
	productRepo := repositories.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productHandler := controllers.NewProductController(productService, response)

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

	// not found handling
	app.Use(NotFoundHandler)
}
