package main

import (
	"fmt"
	"news/config"
	"news/internal/pages"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const DevConfig = "dev.env"

func main() {
	fmt.Println("Hello World")
	config.Init(DevConfig)
	cfg := config.NewConfig()

	app := fiber.New()
	pages.New(app)
	app.Use(recover.New())
	err := app.Listen(":" + cfg.ServerPort)
	if err != nil {
		fmt.Println(err)
		return
	}
}
