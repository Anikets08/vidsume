package controllers

import "github.com/gofiber/fiber/v2"

func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Healthy",
		"data":    nil,
		"code":    200,
	})
}
