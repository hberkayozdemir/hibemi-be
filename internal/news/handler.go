package news

import "github.com/gofiber/fiber/v2"

type Handler struct {
	Service Service
}

func NewHandler(service Service) Handler {
	return Handler{
		Service: service,
	}
}

func (h *Handler) SetupApp(app *fiber.App) {
	app.Post("/addNews", h.RegisterUserHandler)
	app.Get("/getNews", h.LoginUserHandler)
	app.Delete("/new/news/:id", h.DeleteUserHandler)
	app.Get("/news/:hashtags", h.ActivateUser)
}

func (h *Handler) AddNewsHandler(c *fiber.Ctx) error {
	newsDTO := NewsDTO{}
	err := c.BodyParser(&newsDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	user, err := h.Service.RegisterUser(userDTO)

	switch err {
	case nil:
		c.Status(fiber.StatusCreated)
		c.JSON(user)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}
