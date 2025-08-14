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

	route.Post("add_author", func(c *fiber.Ctx) error {
		var authorId int
		var contestId string
		var err error

		authorId, err = (strconv.Atoi(c.Query("author_id", "")))
		if err != nil {
			return err
		}
		if authorId == 0 {
			return c.Status(400).SendString("missing required parameter: author_id")
		}

		contestId = c.Query("contest_id", "")
		if contestId == "" {
			return c.Status(400).SendString("missing required parameter: contest_id")
		}

		cr := repository.NewContestRepository(database.Db)
		if err := cr.AddAuthor(contestId, uint64(authorId)); err != nil {
			return err
		}

		return c.SendStatus(200)
	})
}
