package controller

import (
	"belajar-gofiber-gorm/database"
	"belajar-gofiber-gorm/model/entity"
	"belajar-gofiber-gorm/model/request"
	"fmt"
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func BookControllerGetAll(c *fiber.Ctx) error {
	var books []entity.Book
	if result := database.DB.Find(&books); result.Error != nil {
		log.Println(result.Error)
	}

	return c.JSON(books)
}

func BookControllerCreate(c *fiber.Ctx) error {
	book := new(request.BookCreateRequest)

	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	validate := validator.New()
	if errValidate := validate.Struct(book); errValidate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed, data not valid!",
			"error":   errValidate,
		})
	}

	// Validation Required Image
	filename := c.Locals("filename")
	if filename == nil {
		return c.Status(422).JSON(fiber.Map{
			"message": "image cover is required",
		})
	}

	filenameString := fmt.Sprintf("%v", filename)

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filenameString,
	}

	if errCreatebook := database.DB.Create(&newBook).Error; errCreatebook != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to store data :(",
		})
	}

	return c.JSON(fiber.Map{
		"message": "store data succesfully! :)",
		"data":    newBook,
	})
}

func BookControllerGetById(c *fiber.Ctx) error {
	bookId := c.Params("id")

	var book entity.Book
	if err := database.DB.First(&book, "id = ?", bookId).Error; err != nil {
		// SELECT * FROM books WHERE id = bookId;
		return c.Status(404).JSON(fiber.Map{
			"message": "no data with id " + bookId,
		})
	}

	return c.JSON(fiber.Map{
		"message": "success fetching data",
		"data":    book,
	})
}

func BookControllerUpdate(c *fiber.Ctx) error {
	bookReq := new(request.BookUpdateRequest)

	if err := c.BodyParser(bookReq); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var book entity.Book
	bookId := c.Params("id")

	if err := database.DB.First(&book, "id = ?", bookId).Error; err != nil {
		// SELECT * FROM books WHERE id = bookId;
		return c.Status(404).JSON(fiber.Map{
			"message": "no data with id " + bookId,
		})
	}

	// UPDATE book DATA
	if bookReq.Title != "" {
		book.Title = bookReq.Title
	}
	book.Author = bookReq.Author
	book.Cover = bookReq.Cover
	if errUpdate := database.DB.Save(&book).Error; errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success update data",
		"data":    book,
	})
}

func BookControllerDelete(c *fiber.Ctx) error {
	bookId := c.Params("id")
	var book entity.Book

	if err := database.DB.First(&book, "id = ?", bookId).Error; err != nil {
		// SELECT * FROM books WHERE id = bookId;
		return c.Status(404).JSON(fiber.Map{
			"message": "no data with id " + bookId,
		})
	}

	if errDelete := database.DB.Delete(&book).Error; errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "book with id " + bookId + " has been deleted",
	})
}
