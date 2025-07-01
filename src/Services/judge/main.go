package main

import (
	// "context"
	"fmt"
	// "judge/routes"
	// "judge/storage"
	"judge/utils"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		ServerHeader: "HCMUT-OJ",
	})
	app.Use(logger.New(logger.Config{
		Output: os.Stdout,
	}))

	if os.Getenv("POLYGON_API_KEY") == "" {
		panic("POLYGON_API_KEY not set")
	}

	err = utils.DownloadPackage(map[string]string{
		"problemId": "442306",
		"packageId": "1145478",
		"type":      "standard",
		"apiKey":    os.Getenv("POLYGON_API_KEY"),
		"time":      fmt.Sprintf("%d", time.Now().Unix()),
	})
	if err != nil {
		panic(err.Error())
	}

	// client, err := storage.GetMongoDbClient()
	// if err != nil {
	// 	log.Fatal("Error: ", err)
	// }
	// defer client.Disconnect(context.Background())

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("I'm a GET request!")
	// })
	//
	// routes.ProblemRoute(app.Group("/problem"))
	//
	// for _, route := range app.Stack() {
	// 	for _, r := range route {
	// 		fmt.Printf("%-6s %s\n", r.Method, r.Path)
	// 	}
	// }
	//
	// app.Listen(":3000")
}
