package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load("pkg/.env"); err != nil {
		log.Println("Error loading .env file")
		return
	}
	log.Println(".env file loaded")
}

type DBConfig struct {
	Url string
}

func GetEnv[T any](key string, defaultValue T) T {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	if result, ok := tryConvert[T](value); ok {
		return result
	}
	return defaultValue
}

func tryConvert[T any](value string) (T, bool) {
	var result T

	switch v := any(&result).(type) {
	case *string:
		*v = value
		return result, true
	case *int:
		if intVal, err := strconv.Atoi(value); err == nil {
			*v = intVal
			return result, true
		}
	case *bool:
		if boolVal, err := strconv.ParseBool(value); err == nil {
			*v = boolVal
			return result, true
		}
	case *time.Duration:
		if durVal, err := time.ParseDuration(value); err == nil {
			*v = durVal
			return result, true
		}
	}

	return result, false
}
