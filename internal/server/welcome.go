package server

import "github.com/gofiber/fiber/v2"

func WelcomeHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Welcome to learning API 🚀",
	})
}
