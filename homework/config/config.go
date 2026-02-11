package config

import (
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string     `env:"SERVER_PORT" envDefault:"8080"`
	LogLevel   slog.Level `env:"LOG_LEVEL"`
}

func Init(file string) {
	if err := godotenv.Load("config/" + file); err != nil {
		log.Println("Error loading .env file")
	}
}

func NewConfig() *Config {
	ll := os.Getenv("LOG_LEVEL")
	var logLevel slog.Level

	switch strings.ToLower(ll) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}
	return &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		LogLevel:   logLevel,
	}
}
