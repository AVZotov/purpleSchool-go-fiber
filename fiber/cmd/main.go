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

	app := fiber.New()
	app.Use(recover.New())

	home.NewHandler(app)

	fmt.Println("Start")

	app.Listen(":3000")
}
