package main

import (
	"fmt"
	"judge/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "HCMUT-OJ",
	})
	app.Use(logger.New(logger.Config{
		Output: os.Stdout,
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
	routes.ProblemRoute(app.Group("/problem"))

	for _, route := range app.Stack() {
		for _, r := range route {
			fmt.Printf("%-6s %s\n", r.Method, r.Path)
		}
	}

	app.Listen(":3000")
}
