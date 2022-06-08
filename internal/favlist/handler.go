package favlist

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) SetupApp(app *fiber.App) {
	app.Post("/addAssetToFavList", h.GetFavListHandler)
	app.Get("/getFavList", h.AddAssetToFavListHandler)
	app.Delete("/new/deleteAssetFromFavlist/:id", h.DeleteAssetFromFavListHandler)
}
func (h *Handler) GetFavListHandler(c *fiber.Ctx) error {

	return nil
}
func (h *Handler) AddAssetToFavListHandler(c *fiber.Ctx) error {

	return nil
}
func (h *Handler) DeleteAssetFromFavListHandler(c *fiber.Ctx) error {

	return nil
}