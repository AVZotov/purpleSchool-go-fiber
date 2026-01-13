package main

import (
	"fmt"
	"go-fiber/internal/home"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Start")
	app := fiber.New()

	home.NewHandler(app)

	app.Listen(":3000")
}
