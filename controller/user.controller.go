package controller

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
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	validate := validator.New()
	if errValidate := validate.Struct(user); errValidate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed, data not valid!",
			"error":   errValidate,
		})
	}

	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Address:  user.Address,
		Phone:    user.Phone,
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
		// SELECT * FROM users WHERE id = userId;
		return c.Status(404).JSON(fiber.Map{
			"message": "no data with id " + userId,
		})
	}

	return c.JSON(fiber.Map{
		"message": "success fetching data",
		"data":    user,
	})
}

func UserControllerUpdate(c *fiber.Ctx) error {
	userReq := new(request.UserUpdateRequest)

	if err := c.BodyParser(userReq); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User
	userId := c.Params("id")

	if err := database.DB.First(&user, "id = ?", userId).Error; err != nil {
		// SELECT * FROM users WHERE id = userId;
		return c.Status(404).JSON(fiber.Map{
			"message": "no data with id " + userId,
		})
	}

	// UPDATE USER DATA
	if userReq.Name != "" {
		user.Name = userReq.Name
	}
	user.Address = userReq.Address
	user.Phone = userReq.Phone
	if errUpdate := database.DB.Save(&user).Error; errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success update data",
		"data":    user,
	})
}

func UserControllerUpdateEmail(c *fiber.Ctx) error {
	userReq := new(request.UserEmailRequest)

	if err := c.BodyParser(userReq); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User
	var isEmailUserExist entity.User
	userId := c.Params("id")

	if err := database.DB.First(&user, "id = ?", userId).Error; err != nil {
		// SELECT * FROM users WHERE id = userId;
		return c.Status(404).JSON(fiber.Map{
			"message": "no data with id " + userId,
		})
	}

	if errCheckEmail := database.DB.First(&isEmailUserExist, "email = ?", userReq.Email).Error; errCheckEmail == nil {
		// SELECT * FROM users WHERE id = userId;
		return c.Status(404).JSON(fiber.Map{
			"message": userReq.Email + " already used",
		})
	}

	user.Email = userReq.Email
	if errUpdate := database.DB.Save(&user).Error; errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success update data",
		"data":    user,
	})
}

func UserControllerDelete(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user entity.User

	if err := database.DB.First(&user, "id = ?", userId).Error; err != nil {
		// SELECT * FROM users WHERE id = userId;
		return c.Status(404).JSON(fiber.Map{
			"message": "no data with id " + userId,
		})
	}

	if errDelete := database.DB.Delete(&user).Error; errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "user with id " + userId + " has been deleted",
	})
}
