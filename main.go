package main

import (
	"belajar-gofiber-gorm/database"
	"belajar-gofiber-gorm/database/migration"
	"belajar-gofiber-gorm/routes"

	"github.com/gofiber/fiber/v2"

)

func main() {
	// INITIAL DATABASE
	database.DatabaseInit()

	//  RUN MIGRATION
	migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	routes.RouteInit(app)

	app.Listen(":4000")
}
