package pages

import "github.com/gofiber/fiber/v2"

type Handler struct {
	router fiber.Router
}

type Category struct {
	Name string
}

type PageData struct {
	Categories []Category
}

func New(router fiber.Router) {
	h := &Handler{
		router: router,
	}
	h.router.Get("/", h.home)
}

func (h *Handler) home(c *fiber.Ctx) error {
	data := PageData{
		Categories: []Category{
			{Name: "Еда"},
			{Name: "Животные"},
			{Name: "Машины"},
			{Name: "Спорт"},
			{Name: "Музыка"},
			{Name: "Технологии"},
			{Name: "Прочее"},
		},
	}
	return c.Render("home", data)
}
