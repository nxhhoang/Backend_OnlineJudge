package routes

import (
	"fmt"
	"judge/storage"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ProblemRoute(router fiber.Router) {
	router.Post("/add", func(c *fiber.Ctx) error {
		var problemId, packageId int
		var err error

		problemId, err = strconv.Atoi(c.Query("problemId", ""))
		if err != nil {
			return err
		}
		packageId, err = strconv.Atoi(c.Query("packageId", ""))
		if err != nil {
			return err
		}

		if err := storage.AddProblem(uint64(problemId), uint64(packageId)); err != nil {
			return c.SendString(fmt.Sprintf("Got error: %s", err.Error()))
		}

		return c.SendStatus(200)
	})
}
