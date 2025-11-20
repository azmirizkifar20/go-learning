package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"go-learning/internal/config"
	"go-learning/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	app := fiber.New()

	server.SetupRoutes(app)

	log.Printf("starting server on %s", cfg.AppPort)
	if err := app.Listen(cfg.AppPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
