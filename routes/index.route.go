package routes

import (
	"belajar-gofiber-gorm/config"
	"belajar-gofiber-gorm/controllers"
	"belajar-gofiber-gorm/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/assets")
	r.Get("/user", middleware.Auth, controllers.UserControllerGetAll)
	r.Get("/user/:id", controllers.UserControllerGetById)
	r.Post("/user", controllers.UserControllerCreate)
	r.Put("/user/:id", controllers.UserControllerUpdate)
	r.Put("/user/:id/update-email", controllers.UserControllerUpdateEmail)
	r.Delete("/user/:id", controllers.UserControllerDelete)
}
