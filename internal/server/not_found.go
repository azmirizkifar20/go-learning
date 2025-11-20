package server

import "github.com/gofiber/fiber/v2"

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"status":  "failed",
		"message": "Route not found",
		"data":    nil,
		"error":   "404 Not Found",
	})
}
