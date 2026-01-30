package pages

import (
	"news/pkg/tadaptor"
	"news/views"
	"news/views/components"

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
	blogProps := []components.BlogCardProp{
		{Author: "Михаил Аршинов", AuthorImg: "static/images/blog/michail.jpg", ArticleHeader: "Открытие сезона байдарок", Article: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Est maiores molestiae, vitae dicta nihil porroet. Exercitationem tempore iusto quidem?", Date: "Август 18 , 2025", BlogImg: "static/images/blog/boat.jpg"},
		{Author: "Вася Программист", AuthorImg: "static/images/blog/vasya.jpg", ArticleHeader: "Выбери правильный ноутбук для задач", Article: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Est maiores molestiae, vitae dicta nihil porroet. Exercitationem tempore iusto quidem?", Date: "Июль 25 , 2025", BlogImg: "static/images/blog/comp.jpg"},
		{Author: "Мария", AuthorImg: "static/images/blog/mariya.jpg", ArticleHeader: "Создание автомобилей с автопилотом", Article: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Est maiores molestiae, vitae dicta nihil porroet. Exercitationem tempore iusto quidem?", Date: "Июль 14 , 2025", BlogImg: "static/images/blog/car.jpg"},
		{Author: "Ли Сюн", AuthorImg: "static/images/blog/li.jpg", ArticleHeader: "Как быстро приготовить вкусный обед", Article: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Est maiores molestiae, vitae dicta nihil porroet. Exercitationem tempore iusto quidem?", Date: "Май 10 , 2025", BlogImg: "static/images/blog/food.jpg"},
	}
	topicProps := []components.TopicCardProp{
		{Title: "Как безопасно водить", Text: "Длинный текст про то, как можно безопасно водить автомобиль.", Img: "static/images/topic/car.jpg"},
		{Title: "Создавай музыку!", Text: "Сегодня мы рассмотрим технику быстрого создания музыки за счёт использования...", Img: "static/images/topic/music.jpg"},
	}

	component := views.Main(blogProps, topicProps)
	return tadaptor.Render(c, component)
}
