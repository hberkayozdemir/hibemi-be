package banner

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

	app.Get("/Banners", h.BannersListHandler)
	app.Post("/Banners/CreateBanner", h.CreateBanner)
	app.Post("/Banners/DeleteBanner:id", h.DeleteBannerHandler)
	app.Get("/Banners/getBanner:id", h.GetBannerHandler)
}

func (h *Handler) BannersListHandler(c *fiber.Ctx) error {
	banners, err := h.Service.BannerList()
	switch err {
	case nil:
		c.Status(fiber.StatusOK)
		c.JSON(banners)

	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return err
}

func (h *Handler) CreateBanner(c *fiber.Ctx) error {
	bannerDTO := BannerDTO{}
	err := c.BodyParser(&bannerDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}
	banner, err := h.Service.CreateBanner(bannerDTO)
	switch err {
	case nil:
		c.Status(fiber.StatusCreated)
		c.JSON(banner)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}
func (h *Handler) DeleteBannerHandler(c *fiber.Ctx) error {
	bannerID := c.Params("id")
	if len(bannerID) == 0 {
		c.Status(fiber.StatusBadRequest)
		return nil
	}
	err := h.Service.DeleteNews(bannerID)

	switch err {
	case nil:
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (h *Handler) GetBannerHandler(c *fiber.Ctx) error {
	bannerID := c.Params("id")
	if len(bannerID) == 0 {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	banner, err := h.Service.GetBanner(bannerID)

	switch err {
	case nil:
		c.Status(fiber.StatusOK)
		c.JSON(banner)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil

}
