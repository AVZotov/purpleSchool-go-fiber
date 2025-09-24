package pages

import "github.com/gofiber/fiber/v2"

type Handler struct {
	router fiber.Router
}

func New(router fiber.Router) {
	h := &Handler{
		router: router,
	}

	h.router.Get("/", h.home)
}

func (h *Handler) home(c *fiber.Ctx) error {

	return c.SendString("Home Page")
}
