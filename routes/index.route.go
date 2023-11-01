package routes

import (
	"belajar-gofiber-gorm/config"
	"belajar-gofiber-gorm/controllers"

	"github.com/gofiber/fiber/v2"

)

func middleware(c *fiber.Ctx) error {
	if token := c.Get("x-token"); token != "secret" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	return c.Next()
}

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/assets")
	r.Get("/user", middleware, controllers.UserControllerGetAll)
	r.Get("/user/:id", controllers.UserControllerGetById)
	r.Post("/user", controllers.UserControllerCreate)
	r.Put("/user/:id", controllers.UserControllerUpdate)
	r.Put("/user/:id/update-email", controllers.UserControllerUpdateEmail)
	r.Delete("/user/:id", controllers.UserControllerDelete)
}
