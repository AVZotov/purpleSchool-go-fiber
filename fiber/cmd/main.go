package main

import (
	"fmt"
	"go-fiber/config"
	"go-fiber/internal/home"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/rs/zerolog"
)

func main() {
	config.Init()

	logConfig := config.LogConfig{
		Level: config.GetEnv("LOG_LEVEL", 0),
	}
	fmt.Printf("Level: %d, type: %T\n", logConfig.Level, logConfig.Level)

	engine := html.New("./html", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.Level(logConfig.Level))
	app.Use(fiberzerolog.New())
	app.Use(recover.New())

	home.NewHandler(app)

	app.Listen(":3000")
}
