package main

import (
	"fmt"
	"go-fiber/config"
	"go-fiber/internal/home"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	dbUrl := config.GetEnv("DATABASE_URL", "www.1.com")
	fmt.Println(dbUrl)

	app := fiber.New()
	app.Use(recover.New())

	home.NewHandler(app)

	app.Listen(":3000")
}
