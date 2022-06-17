package coin

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
	app.Get("/getSpots", h.GetAllSpotsHandler)
	app.Get("/getCoins", h.GetAllCoinHandler)
}

func (h *Handler) GetAllSpotsHandler(c *fiber.Ctx) error {

	spots, err := h.Service.GetAllSpots()
	switch err {
	case nil:
		c.Status(fiber.StatusOK)
		c.JSON(spots)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}

func (h *Handler) GetAllCoinHandler(c *fiber.Ctx) error {

	coins, err := h.Service.GetAllCoins()
	switch err {
	case nil:
		c.Status(fiber.StatusOK)
		c.JSON(coins)
	default:
		fmt.Print(err)
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}
