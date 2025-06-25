package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ProblemRoute(router fiber.Router) {
	router.Post("/add", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(400).SendString("Failed to read file")
		}

		destination := fmt.Sprintf("/upload/%s", file.Filename)
		if err = c.SaveFile(file, destination); err != nil {
			return c.Status(400).SendString("Failed to save file")
		}

		return c.SendStatus(200)
	})
}
