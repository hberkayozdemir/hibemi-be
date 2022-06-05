package foo

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/modanisatech/marketplace/shared/errors"
	"gitlab.com/modanisatech/marketplace/shared/log"
)

type Handler struct{}

func (h Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/foo", h.GetFoo)
}

func (h Handler) GetFoo(c *fiber.Ctx) error {
	log.AddFields(c, log.EventName("getFooEvent"))

	countQuery := c.Query("count")
	if countQuery == "" {
		return errors.BadRequest("count query parameter can not be empty")
	}

	count, err := strconv.Atoi(countQuery)
	if err != nil {
		return errors.BadRequest("count query parameter must be number").Err(err)
	}

	result, err := fooRepeater(c.Context(), count)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(map[string]string{"result": result})
}
