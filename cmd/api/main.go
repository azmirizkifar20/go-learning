package main

import (
	"log"

	"go-learning/internal/config"
	"go-learning/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	app := server.NewServer()

	log.Printf("starting server on %s", cfg.AppPort)
	if err := app.Listen(cfg.AppPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
