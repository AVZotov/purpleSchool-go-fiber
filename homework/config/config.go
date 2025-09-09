package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init(file string) {
	if err := godotenv.Load("config/" + file); err != nil {
		log.Println("Error loading .env file")
	}
}

type Config struct {
	ServerPort string `env:"SERVER_PORT" envDefault:"8080"`
}

func NewConfig() *Config {
	return &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
