package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(500).JSON(fiber.Map{
			"message": "Hello World",
		})
		//return c.SendString("Hello World!")
	})

	app.Listen(":4000")
}
