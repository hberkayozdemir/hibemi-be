package favlist

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
	app.Post("/Users/Favlist/AddCoin", h.CreateFavCoinHandler)
}

func (h *Handler) CreateFavCoinHandler(c *fiber.Ctx) error {
	favCoinDTO := FavCoinDTO{}
	err := c.BodyParser(&favCoinDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	createdFavCoin, err := h.Service.CreateFavCoin(favCoinDTO)
	fmt.Println(createdFavCoin)
	switch err {
	case nil:
		c.JSON(createdFavCoin)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}
