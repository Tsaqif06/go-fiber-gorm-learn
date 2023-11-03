package controller

import (
	"belajar-gofiber-gorm/database"
	"belajar-gofiber-gorm/model/entity"
	"belajar-gofiber-gorm/model/request"
	"belajar-gofiber-gorm/utils"
	"log"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

	// VALIDATION PASSWORD
	if isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password); !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong password!",
		})
	}

	// GENERATE JWT
	claims := jwt.MapClaims{}
	claims["sub"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential!",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
