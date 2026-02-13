package main

import (
	"context"
	"log/slog"
	"news/config"
	"news/internal/pages"
	"news/internal/repository"
	"news/pkg/database"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

const Configs = "config.env"

func main() {
	config.Init(Configs)
	cfg := config.NewConfig()

	handlerOpts := &slog.HandlerOptions{Level: cfg.Level}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, handlerOpts))

	slogFiberConfig := slogfiber.Config{
		DefaultLevel:     cfg.Level,
		ClientErrorLevel: cfg.Level,
		ServerErrorLevel: cfg.Level,
	}

	// Подключение к БД
	db, err := database.Connect(context.Background(), cfg.GetURL(), logger)
	if err != nil {
		logger.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// Создаём репозиторий
	userRepo := repository.NewUserRepository(db)

	app := fiber.New()

	app.Use(slogfiber.NewWithConfig(logger, slogFiberConfig))
	app.Use(recover.New())
	app.Static("/static", "./static")

	// Передаём репозиторий и логгер в pages
	pages.New(app, userRepo, logger)

	logger.Info("Starting server", "port", cfg.ServerConfig.Port)
	if err := app.Listen(":" + cfg.ServerConfig.Port); err != nil {
		logger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
