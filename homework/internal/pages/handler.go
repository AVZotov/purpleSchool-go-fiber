package pages

import (
	"news/pkg/tadaptor"
	"news/views"
	"news/views/components"
	"news/views/widgets"

	"github.com/a-h/templ"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	router fiber.Router
}

var validate = validator.New()

func New(router fiber.Router) {
	h := &Handler{
		router: router,
	}
	h.router.Get("/", h.home)
	h.router.Get("/api/register", h.register)
	h.router.Post("/api/register", h.RegisterApi)
}

func (h *Handler) home(c *fiber.Ctx) error {
	blogProps := []components.BlogCardProps{
		{Author: "Михаил Аршинов", AuthorImg: "static/images/blog/michail.jpg", ArticleHeader: "Открытие сезона байдарок", Article: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Est maiores molestiae, vitae dicta nihil porroet.", Date: "Август 18 , 2025", BlogImg: "static/images/blog/boat.jpg"},
		{Author: "Вася Программист", AuthorImg: "static/images/blog/vasya.jpg", ArticleHeader: "Выбери правильный ноутбук для задач", Article: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Est maiores molestiae, vitae dicta nihil porroet.", Date: "Июль 25 , 2025", BlogImg: "static/images/blog/comp.jpg"},
		{Author: "Мария", AuthorImg: "static/images/blog/mariya.jpg", ArticleHeader: "Создание автомобилей с автопилотом", Article: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Est maiores molestiae, vitae dicta nihil porroet.", Date: "Июль 14 , 2025", BlogImg: "static/images/blog/car.jpg"},
		{Author: "Ли Сюн", AuthorImg: "static/images/blog/li.jpg", ArticleHeader: "Как быстро приготовить вкусный обед", Article: "Lorem ipsum dolor sit amet consectetur adipisicing elit.", Date: "Май 10 , 2025", BlogImg: "static/images/blog/food.jpg"},
	}
	topicProps := []components.TopicCardProps{
		{Title: "Как безопасно водить", Text: "Длинный текст про то, как можно безопасно водить автомобиль.", Img: "static/images/topic/car.jpg"},
		{Title: "Создавай музыку!", Text: "Сегодня мы рассмотрим технику быстрого создания музыки за счёт использования...", Img: "static/images/topic/music.jpg"},
	}

	component := views.Main(blogProps, topicProps)
	return tadaptor.Render(c, component)
}

func (h *Handler) register(c *fiber.Ctx) error {
	component := views.Register()
	return tadaptor.Render(c, component)
}

func (h *Handler) RegisterApi(c *fiber.Ctx) error {
	var req RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).SendString("error parsing request")
	}

	if err := validate.Struct(req); err != nil {
		errors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			switch field {
			case "Name":
				errors["name"] = "Name should be more than 2 simbols"
			case "Email":
				errors["email"] = "Please enter proper email"
			case "Password":
				errors["password"] = "Password must contains at least 5 simbols"
			}
		}

		inputs := views.GetRegistrationInputForms(errors)
		component := widgets.RegisterForm(inputs)
		return tadaptor.Render(c, component)
	}

	successMsg := templ.Raw("<div style='color: green; text-align: center; padding: 20px;'><i class='fa-solid fa-circle-check'></i>Регистрация прошла успешно!</div>")
	return tadaptor.Render(c, successMsg)
}
