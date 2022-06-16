package coin

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
	app.Get("/getSpots", h.GetAllSpotsHandler)
	app.Get("/getCoins", h.GetAllCoinHandler)
}

func (h *Handler) GetAllSpotsHandler(c *fiber.Ctx) error {
	pageStr := c.Query("page")
	page := 0
	if len(pageStr) != 0 {
		var err error
		page, err = strconv.Atoi(pageStr)
		if page < 0 || err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}
	}
	sizeStr := c.Query("size")
	size := 0
	if len(sizeStr) != 0 {
		var err error
		size, err = strconv.Atoi(sizeStr)
		if size <= 0 || err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}
	}
	spots, err := h.Service.GetAllSpots(page, size)
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
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}
