package home

import (
	//"bytes"
	//"html/template"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &Handler{
		router: router,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *Handler) home(c *fiber.Ctx) error {
	data := struct {
		Count   int
		IsAdmin bool
	}{Count: 5, IsAdmin: true}

	return c.Render("page", data)
}
func (h *Handler) error(c *fiber.Ctx) error {

	return c.SendString("Error")
}
