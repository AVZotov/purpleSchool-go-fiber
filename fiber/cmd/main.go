package main

import (
	"go-fiber/config"
	"go-fiber/internal/home"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
)

func main() {
	config.Init()

	logConfig := config.LogConfig{
		Level: config.GetEnv("LOG_LEVEL", 0),
	}
	app := fiber.New()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.Level(logConfig.Level))
	app.Use(fiberzerolog.New())
	app.Use(recover.New())
	app.Static("/static", "./static")

	home.NewHandler(app)

	app.Listen(":3000")
}
