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
	app.Post("/login", h.LoginUserHandler)
	app.Delete("/user/users/:userID", h.DeleteUserHandler)
	app.Patch("/users/:userID/activate", h.ActivateUser)
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

	user, err := h.Service.ActivateUser(activationCodeDTO.Email, activationCodeDTO.Code)

	switch err {
	case nil:
		c.JSON(user)
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
