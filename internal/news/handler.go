package news

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) Handler {
	return Handler{
		Service: service,
	}
}

func (h *Handler) SetupApp(app *fiber.App) {
	app.Post("/addNews", h.AddNewsHandler)
	app.Get("/getNews", h.GetNewsHandler)
	app.Delete("/new/news/:id", h.DeleteNewsHandler)
}

func (h *Handler) GetNewsHandler(c *fiber.Ctx) error {
	pageStr := c.Query("page")
	page := 0
	if len(pageStr) != 0 {
		var err error
		page, err = strconv.Atoi(pageStr)
		if page <= 0 || err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}
	}
	sizeStr := c.Query("size")
	size := 0
	if len(sizeStr) != 0 {
		var err error
		page, err = strconv.Atoi(sizeStr)
		if size <= 0 || err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}
	}

	news, err := h.Service.getNews(page, size)
	switch err {
	case nil:
		c.Status(fiber.StatusOK)
		c.JSON(news)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil

}

func (h *Handler) AddNewsHandler(c *fiber.Ctx) error {
	newsDTO := NewsDTO{}
	err := c.BodyParser(&newsDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	user, err := h.Service.AddNews(newsDTO)

	switch err {
	case nil:
		c.Status(fiber.StatusCreated)
		c.JSON(user)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}

func (h *Handler) DeleteNewsHandler(c *fiber.Ctx) error {
	newsSearchDTO := NewSearchCredentialsDTO{}
	err := c.BodyParser(&newsSearchDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	err = h.Service.DeleteNews(newsSearchDTO.ID)

	switch err {
	case nil:
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}
