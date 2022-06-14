package transactions

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
	app.Post("/transactions", h.CreateTransactionHandler)
	app.Get("/users/:id/transactions", h.GetTransactionHistoryHandler)
}

func (h *Handler) CreateTransactionHandler(c *fiber.Ctx) error {
	transactionDTO := TransactionDTO{}
	err := c.BodyParser(&transactionDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	transaction, err := h.Service.CreateTransaction(transactionDTO)

	switch err {
	case nil:
		c.Status(fiber.StatusCreated)
		c.JSON(transaction)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}

func (h *Handler) GetTransactionHistoryHandler(c *fiber.Ctx) error {
	userId := c.Params("id")

	transactions, err := h.Service.GetTransactionHistory(userId)
	switch err {
	case nil:
		c.Status(fiber.StatusOK)
		c.JSON(transactions)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return err
}
