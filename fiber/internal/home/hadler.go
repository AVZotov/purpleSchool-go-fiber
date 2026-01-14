package home

import "github.com/gofiber/fiber/v2"

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
	panic("implement me")
	return c.SendString("Home")
}
func (h *Handler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
