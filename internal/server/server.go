package server

import (
	"go-learning/internal/helpers"
	"go-learning/internal/routes"
	"time"

	"go-learning/internal/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func NewServer() *fiber.App {
	app := fiber.New()
	response := helpers.NewResponse()

	// rate limit
	app.Use(limiter.New(limiter.Config{
		Max:        60,              // max 60 request
		Expiration: 1 * time.Minute, // per 1 menit per key
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
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

	// welcome route
	app.Get("/", WelcomeHandler)

	// init dependencies
	db := database.GetDB()
	deps := routes.NewDependencies(db)

	// setup routes
	routes.SetupRoutes(app, deps)

	// not found handling
	app.Use(NotFoundHandler)

	return app
}
