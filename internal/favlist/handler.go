package favlist

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hberkayozdemir/hibemi-be/internal/auth"
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
	bearerToken := string(c.Request().Header.Peek("Authorization"))
	isTokenValid, _ := auth.VerifyToken(bearerToken)
	if !isTokenValid {
		c.Status(fiber.StatusUnauthorized)
		return nil
	}

	favCoinDTO := FavCoinDTO{}
	err := c.BodyParser(&favCoinDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	createdFavCoin, err := h.Service.CreateFavCoin(favCoinDTO)
	switch err {
	case nil:
		c.JSON(createdFavCoin)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}
