package routes

import (
	"belajar-gofiber-gorm/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", controllers.UserControllerGetAll)
	r.Get("/user/:id", controllers.UserControllerGetById)
	r.Post("/user", controllers.UserControllerCreate)
	r.Put("/user/:id", controllers.UserControllerUpdate)
}
