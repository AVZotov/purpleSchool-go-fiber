package home

import (
	//"bytes"
	//"html/template"

	"go-fiber/pkg/tadaptor"
	"go-fiber/views"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &Handler{
		router: router,
	}
	h.router.Get("/", h.home)
}

func (h *Handler) home(c *fiber.Ctx) error {
	component := views.Main()
	return tadaptor.Render(c, component)
}
func (h *Handler) error(c *fiber.Ctx) error {

	return c.SendString("Error")
}
