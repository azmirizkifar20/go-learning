package server

import (
	"go-learning/internal/controllers"
	"go-learning/internal/helpers"
	"go-learning/internal/repositories"
	"go-learning/internal/services"
	"go-learning/internal/storage"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func SetupRoutes(app *fiber.App) {
	// init response
	response := helpers.NewResponse()

	// GLOBAL RATE LIMITER
	app.Use(limiter.New(limiter.Config{
		Max:        60,              // max 60 request
		Expiration: 1 * time.Minute, // per 1 menit per key
		// Key per user (di sini per IP)
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		// custom response saat kena limit
		LimitReached: func(c *fiber.Ctx) error {
			return response.Send(
				c,
				fiber.StatusTooManyRequests,
				nil,
				"Too many requests, please try again later",
				"rate_limit_exceeded",
			)
		},
	}))

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
