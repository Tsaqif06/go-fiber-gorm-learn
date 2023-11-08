package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(c *fiber.Ctx) error {
	// HANDLE FILE
	file, errFile := c.FormFile("cover")
	if errFile != nil {
		log.Println("error file = ", errFile)
	}

	var filename string
	if file != nil {
		filename = file.Filename
		if errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/covers/%s", filename)); errSaveFile != nil {
			log.Println("fail to store file")
		}
	} else {
		log.Println("nothing files to be uploaded")
	}

	c.Locals("filename", filename)

	return c.Next()
}
