package middleware

import "github.com/gofiber/fiber/v2"

func Auth(c *fiber.Ctx) error {
	if token := c.Get("x-token"); token != "secret" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	return c.Next()
}

func PermissionCreate(c *fiber.Ctx) error {
	return c.Next()
}