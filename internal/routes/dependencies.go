package routes

import (
	"go-learning/internal/controllers"
	"go-learning/internal/helpers"
	"go-learning/internal/repositories"
	"go-learning/internal/services"
	"go-learning/internal/storage"

	"gorm.io/gorm"
)

type Dependencies struct {
	Response *helpers.Response

	CategoryController *controllers.CategoryController
	ProductController  *controllers.ProductController
}

func NewDependencies(db *gorm.DB) *Dependencies {
	response := helpers.NewResponse()
	minioClient := storage.NewMinioClient()

	// Category deps
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo, minioClient)
	categoryController := controllers.NewCategoryController(categoryService, response)

	// Product deps
	productRepo := repositories.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService, response)

	return &Dependencies{
		Response:           response,
		CategoryController: categoryController,
		ProductController:  productController,
	}
}
