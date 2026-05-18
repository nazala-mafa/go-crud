package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type databaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type appConfig struct {
	Name string
	Env  string
	Url  string
}

type jwt struct {
	Secret string
}

var DatabaseConfig databaseConfig
var App appConfig
var Jwt jwt

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}

	DatabaseConfig = databaseConfig{
		Host:     getEnv("DB_HOST", "test"),
		Port:     getEnv("DB_PORT", "test"),
		User:     getEnv("DB_USER", "test"),
		Password: getEnv("DB_PASSWORD", "test"),
		Name:     getEnv("DB_NAME", "test"),
	}

	App = appConfig{
		Name: getEnv("APP_NAME", "test"),
		Env:  getEnv("APP_ENV", "test"),
		Url:  getEnv("APP_URL", "test"),
	}

	Jwt = jwt{
		Secret: getEnv("JWT_SECRET", "test"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
