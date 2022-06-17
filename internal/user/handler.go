package user

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
	app.Post("/register", h.RegisterUserHandler)
	app.Post("/registerEditor", h.RegisterEditorHandler)
	app.Post("/login", h.LoginUserHandler)
	app.Post("/user/users/:userID", h.DeleteUserHandler)
	app.Post("/users/:userID/activate", h.ActivateUser)
	app.Get("/admin/stats", h.GetAllStatsHandler)
	app.Post("/users/activate", h.ActivateUser)
}

func (h *Handler) RegisterUserHandler(c *fiber.Ctx) error {
	userDTO := UserDTO{}
	err := c.BodyParser(&userDTO)
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

func (h *Handler) RegisterEditorHandler(c *fiber.Ctx) error {
	editorDTO := EditorDTO{}
	err := c.BodyParser(&editorDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	editor, err := h.Service.RegisterEditor(editorDTO)

	switch err {
	case nil:
		c.Status(fiber.StatusCreated)
		c.JSON(editor)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}

func (h *Handler) LoginUserHandler(c *fiber.Ctx) error {
	userCredentialsDTO := UserCredentialsDTO{}
	err := c.BodyParser(&userCredentialsDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	token, cookie, err := h.Service.LoginUser(userCredentialsDTO)

	switch err {
	case nil:
		c.JSON(token)
		c.Cookie(cookie)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (h *Handler) DeleteUserHandler(c *fiber.Ctx) error {
	userID := c.Params("userID")
	if len(userID) == 0 {
		c.Status(fiber.StatusBadRequest)
		return nil
	}
	err := h.Service.DeleteUser(userID)

	switch err {
	case nil:
		c.Status(fiber.StatusNoContent)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}

func (h *Handler) ActivateUser(c *fiber.Ctx) error {
	activationCodeDTO := ActivationCodeDTO{}
	err := c.BodyParser(&activationCodeDTO)

	if len(activationCodeDTO.Code) == 0 {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	token, _, err := h.Service.ActivateUser(activationCodeDTO.Email, activationCodeDTO.Code)

	switch err {
	case nil:
		c.JSON(token)
		c.Status(fiber.StatusOK)
	case UserNotFound:
		c.Status(fiber.StatusNotFound)
	case UserAlreadyActivated:
		c.Status(fiber.StatusBadRequest)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}

func (h *Handler) GetAllStatsHandler(c *fiber.Ctx) error {
	stat, err := h.Service.GetAllStats()

	switch err {
	case nil:
		c.Status(fiber.StatusOK)
		c.JSON(stat)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}
