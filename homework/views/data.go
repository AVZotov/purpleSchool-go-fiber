package views

import "news/views/components"

var HomePageProps = components.LayoutProps{
	Title:           "Homework",
	MetaDescription: "TEMPL Based WebSite",
}

var NavItemProps = []components.NavItemProps{
	{Name: "#Еда", Link: "#", Image: "static/images/nav/food.jpg"},
	{Name: "#Животные", Link: "#", Image: "static/images/nav/animal.jpg"},
	{Name: "#Машины", Link: "#", Image: "static/images/nav/car.jpg"},
	{Name: "#Спорт", Link: "#", Image: "static/images/nav/sport.jpg"},
	{Name: "#Музыка", Link: "#", Image: "static/images/nav/music.jpg"},
	{Name: "#Технологии", Link: "#", Image: "static/images/nav/tech.jpg"},
	{Name: "#Прочее", Link: "#", Image: "static/images/nav/other.jpg"},
}

var LinkIconOnlyLeftRight = []components.ButtonProps{
	{Arrow: components.ArrowLeft, Variant: components.ButtonIconOnly, Link: "#"},
	{Arrow: components.ArrowRight, Variant: components.ButtonIconOnly, Link: "#"},
}
