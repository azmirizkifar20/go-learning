package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"go-learning/internal/config"
	// "go-learning/internal/database"
	// "go-learning/internal/domain/category"
	// "go-learning/internal/domain/product"
	"go-learning/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	// db := database.GetDB()

	// migration
	// if err := db.AutoMigrate(
	// 	&category.Category{},
	// 	&product.Product{},
	// ); err != nil {
	// 	log.Fatalf("failed to migrate: %v", err)
	// }

	app := fiber.New()

	server.SetupRoutes(app)

	log.Printf("starting server on %s", cfg.AppPort)
	if err := app.Listen(cfg.AppPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
