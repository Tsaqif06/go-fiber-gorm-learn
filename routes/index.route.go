package routes

import (
	"belajar-gofiber-gorm/config"
	"belajar-gofiber-gorm/controller"
	"belajar-gofiber-gorm/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/assets")

	r.Post("/login", controller.LoginController)

	r.Get("/user", middleware.Auth, controller.UserControllerGetAll)
	r.Get("/user/:id", controller.UserControllerGetById)
	r.Post("/user", controller.UserControllerCreate)
	r.Put("/user/:id", controller.UserControllerUpdate)
	r.Put("/user/:id/update-email", controller.UserControllerUpdateEmail)
	r.Delete("/user/:id", controller.UserControllerDelete)
}
