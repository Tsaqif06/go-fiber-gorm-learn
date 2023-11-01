package controller

import (
	"belajar-gofiber-gorm/database"
	"belajar-gofiber-gorm/model/entity"
	"belajar-gofiber-gorm/model/request"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func LoginController(c *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	validate := validator.New()
	if errValidate := validate.Struct(loginRequest); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var user entity.User
	if err := database.DB.First(&user, "email = ?", loginRequest.Email).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "no data with email " + loginRequest.Email,
		})
	}

	return c.JSON(fiber.Map{
		"token": "chelsea",
	})
}
