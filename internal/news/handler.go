package news

import (
	"github.com/gofiber/fiber/v2"
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
	app.Post("/news/:id", h.DeleteNewsHandler)
	app.Get("/news/:id", h.GetNewHandler)
}

func (h *Handler) GetNewsHandler(c *fiber.Ctx) error {

	news, err := h.Service.getNews()
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
	newID := c.Params("id")
	if len(newID) == 0 {
		c.Status(fiber.StatusBadRequest)
		return nil
	}
	err := h.Service.DeleteNews(newID)

	switch err {
	case nil:
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (h *Handler) GetNewHandler(c *fiber.Ctx) error {
	newID := c.Params("id")
	if len(newID) == 0 {
		c.Status(fiber.StatusOK)
		return nil
	}

	new, err := h.Service.GetNew(newID)

	switch err {
	case nil:
		c.Status(fiber.StatusOK)
		c.JSON(new)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}
