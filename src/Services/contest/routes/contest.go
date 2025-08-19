package routes

import (
	repository "contest/domain/repository/contest/impl"
	"contest/internal/infrastructure/database"
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ContestRoutes(app *fiber.App) {
	route := app.Group("/contest")

	route.Post("/create", func(c *fiber.Ctx) error {
		cr := repository.NewContestRepository(database.Db)

		var cancel context.CancelFunc
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var authorId uint64
		var err error
		if authorId, err = strconv.ParseUint(c.Query("author-id", "0"), 10, 64); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		if authorId == 0 {
			return c.Status(400).SendString("author-id is required")
		}

		contestId, err := cr.Create(ctx, authorId)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(map[string]interface{}{
			"contest-id": contestId,
		})
	})

	editRoute := route.Group("/edit")

	editRoute.Post("/add-people", func(c *fiber.Ctx) error {
		var contestId string
		var peopleType string
		var userId uint64

		if contestId = c.Query("contest-id", ""); contestId == "" {
			return c.Status(400).SendString("contest-id required")
		}

		if peopleType = c.Query("people-type", ""); peopleType == "" {
			return c.Status(400).SendString("people-type required")
		}

		userId, err := strconv.ParseUint(c.Query("user-id", "0"), 10, 64)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		if userId == 0 {
			return c.Status(400).SendString("user-id must be non-zero")
		}

		cr := repository.NewContestRepository(database.Db)

		if err := cr.AddPeople(contestId, peopleType, userId); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		return c.SendStatus(200)
	})
}
