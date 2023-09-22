package migration

import (
	"belajar-gofiber-gorm/database"
	"belajar-gofiber-gorm/model/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
