package pages

import (
	"context"
	"log/slog"
	"news/internal/repository"
	"news/pkg/tadaptor"
	"news/views"
	"news/views/components"
	"news/views/widgets"
	"strings"

	"github.com/a-h/templ"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	router   fiber.Router
	userRepo *repository.UserRepository
	logger   *slog.Logger
}

var validate = validator.New()

func New(router fiber.Router, userRepo *repository.UserRepository, logger *slog.Logger) {
	h := &Handler{
		router:   router,
		userRepo: userRepo,
		logger:   logger,
	}
	h.router.Get("/", h.home)
	h.router.Get("/register", h.register)
	h.router.Post("/api/register", h.RegisterApi)
}

func (h *Handler) home(c *fiber.Ctx) error {
	blogProps := getBlogs()
	topicProps := getTopics()

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

	user, err := h.userRepo.Create(context.Background(), req.Name, req.Email, req.Password)
	if err != nil {
		// Проверяем, это ошибка дубликата email?
		if strings.Contains(err.Error(), "unique") || strings.Contains(err.Error(), "duplicate") {
			h.logger.Warn("Email already exists", "email", req.Email)

			// Возвращаем форму с ошибкой
			errors := map[string]string{
				"email": "Этот email уже зарегистрирован",
			}
			inputs := views.GetRegistrationInputForms(errors)
			component := widgets.RegisterForm(inputs)
			return tadaptor.Render(c, component)
		}

		// Неизвестная ошибка БД
		h.logger.Error("Failed to create user", "error", err, "email", req.Email)
		return c.Status(500).SendString("Ошибка сервера. Попробуйте позже.")
	}

	// 4. Успех!
	h.logger.Info("User registered successfully",
		"user_id", user.ID,
		"email", user.Email,
		"username", user.Username,
	)

	successMsg := templ.Raw("<div style='color: green; text-align: center; padding: 20px;'><i class='fa-solid fa-circle-check'></i>Регистрация прошла успешно!</div>")
	return tadaptor.Render(c, successMsg)
}

func getBlogs() []components.BlogCardProps {
	return []components.BlogCardProps{
		{Author: "Михаил Аршинов", AuthorImg: "static/images/blog/michail.jpg", ArticleHeader: "Открытие сезона байдарок", Article: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Est maiores molestiae, vitae dicta nihil porroet.", Date: "Август 18 , 2025", BlogImg: "static/images/blog/boat.jpg"},
		{Author: "Вася Программист", AuthorImg: "static/images/blog/vasya.jpg", ArticleHeader: "Выбери правильный ноутбук для задач", Article: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Est maiores molestiae, vitae dicta nihil porroet.", Date: "Июль 25 , 2025", BlogImg: "static/images/blog/comp.jpg"},
		{Author: "Мария", AuthorImg: "static/images/blog/mariya.jpg", ArticleHeader: "Создание автомобилей с автопилотом", Article: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Est maiores molestiae, vitae dicta nihil porroet.", Date: "Июль 14 , 2025", BlogImg: "static/images/blog/car.jpg"},
		{Author: "Ли Сюн", AuthorImg: "static/images/blog/li.jpg", ArticleHeader: "Как быстро приготовить вкусный обед", Article: "Lorem ipsum dolor sit amet consectetur adipisicing elit.", Date: "Май 10 , 2025", BlogImg: "static/images/blog/food.jpg"},
	}
}

func getTopics() []components.TopicCardProps {
	return []components.TopicCardProps{
		{Title: "Как безопасно водить", Text: "Длинный текст про то, как можно безопасно водить автомобиль.", Img: "static/images/topic/car.jpg"},
		{Title: "Создавай музыку!", Text: "Сегодня мы рассмотрим технику быстрого создания музыки за счёт использования...", Img: "static/images/topic/music.jpg"},
	}
}
