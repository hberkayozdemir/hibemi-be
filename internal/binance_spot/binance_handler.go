package binance_spot

import (
	"fmt"
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
	app.Get("/fetchSpots", h.fetchSpotsHandler)

}

func (h *Handler) fetchSpotsHandler(c *fiber.Ctx) error {

	symbols, err := h.Service.getSpots()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
	}
	switch err {
	case nil:
		c.Status(fiber.StatusCreated)
		fmt.Print(symbols)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}
