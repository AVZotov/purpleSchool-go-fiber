package main

import (
	"context"
	"log/slog"
	"news/config"
	"news/internal/pages"
	"news/internal/repository"
	"news/pkg/database"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger" // ← встроенный logger
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
)

const Configs = "config.env"

func main() {
	config.Init(Configs)
	cfg := config.NewConfig()

	handlerOpts := &slog.HandlerOptions{Level: cfg.Level}
	slogLogger := slog.New(slog.NewJSONHandler(os.Stdout, handlerOpts))

	// Подключение к БД
	db, err := database.Connect(context.Background(), cfg.GetURL(), slogLogger)
	if err != nil {
		slogLogger.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// Создаём хранилище сессий
	storage := postgres.New(postgres.Config{
		DB:         db,
		Table:      "fiber_storage",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})

	store := session.New(session.Config{
		Expiration:     24 * time.Hour,
		CookieHTTPOnly: true,
		CookieSameSite: "Lax",
		Storage:        storage,
	})

	// Создаём репозиторий
	userRepo := repository.NewUserRepository(db)

	app := fiber.New()

	// Встроенный logger Fiber (вместо slog-fiber)
	app.Use(logger.New())
	app.Use(recover.New())
	app.Static("/static", "./static")

	// Передаём репозиторий, store и логгер в pages
	pages.New(app, userRepo, store, slogLogger)

	slogLogger.Info("Starting server", "port", cfg.ServerConfig.Port)
	if err := app.Listen(":" + cfg.ServerConfig.Port); err != nil {
		slogLogger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
