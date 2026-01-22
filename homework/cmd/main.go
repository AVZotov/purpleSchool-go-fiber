package main

import (
	"log/slog"
	"news/config"
	"news/internal/pages"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

const Configs = "config.env"

func main() {
	config.Init(Configs)
	cfg := config.NewConfig()

	handlerOpts := &slog.HandlerOptions{Level: cfg.LogLevel}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, handlerOpts))

	slogFiberConfig := slogfiber.Config{
		DefaultLevel:     cfg.LogLevel,
		ClientErrorLevel: cfg.LogLevel,
		ServerErrorLevel: cfg.LogLevel,
	}

	app := fiber.New()

	app.Use(slogfiber.NewWithConfig(logger, slogFiberConfig))
	app.Use(recover.New())
	app.Static("/static", "./static")

	pages.New(app)

	err := app.Listen(":" + cfg.ServerPort)
	if err != nil {
		logger.Error(err.Error())
		return
	}
}
