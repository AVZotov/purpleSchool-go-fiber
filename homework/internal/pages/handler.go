package pages

import (
	"news/pkg/tadaptor"
	"news/views"

	"github.com/gofiber/fiber/v2"
)

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
	component := views.Main()
	return tadaptor.Render(c, component)
}
