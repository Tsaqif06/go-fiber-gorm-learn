package controllers

import (
	"belajar-gofiber-gorm/database"
	"belajar-gofiber-gorm/model/entity"
	"belajar-gofiber-gorm/model/request"
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func UserControllerGetAll(c *fiber.Ctx) error {
	// //userInfo := ctx.Locals("userInfo").(jwt.MapClaims)
	// //log.Println("email :: ", userInfo["email"])

	var users []entity.User
	if result := database.DB.Find(&users); result.Error != nil {
		log.Println(result.Error)
	}

	// err := database.DB.Find(&users).Error
	// if err != nil {
	// 	log.Println(err)
	// }

	return c.JSON(users)
}

func UserControllerCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	if errValidate := validate.Struct(user); errValidate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed, data not valid!",
			"error":   errValidate,
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	if errCreateUser := database.DB.Create(&newUser).Error; errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to store data :(",
		})
	}

	return c.JSON(fiber.Map{
		"message": "store data succesfully! :)",
		"data":    newUser,
	})
}

func UserControllerGetById(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user entity.User
	if err := database.DB.First(&user, "id = ?", userId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "no data with id " + userId,
		})
	}

	return c.JSON(fiber.Map{
		"message": "success fetching data",
		"data":    user,
	})
}
