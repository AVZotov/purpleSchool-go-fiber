package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerConfig
	LogConfig
	DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type LogConfig struct {
	Level slog.Level
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (c DatabaseConfig) GetURL() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
		c.SSLMode,
	)
}

func Init(file string) {
	if err := godotenv.Load("config/" + file); err != nil {
		log.Println("Error loading .env file")
	}
}

func NewConfig() *Config {
	return &Config{
		ServerConfig: ServerConfig{
			Port: getEnv("SERVER_PORT", "3000"),
		},
		LogConfig: LogConfig{
			Level: getLogLevel(),
		},
		DatabaseConfig: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "news_db"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
	}
}

func getLogLevel() slog.Level {
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
	return logLevel
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
