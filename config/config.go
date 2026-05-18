package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type AppConfig struct {
	Name string
	Env  string
	Url  string
}

type JwtConfig struct {
	Secret string
}

type Config struct {
	Database DatabaseConfig
	App      AppConfig
	Jwt      JwtConfig
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env not found")
	}

	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "test"),
			Port:     getEnv("DB_PORT", "test"),
			User:     getEnv("DB_USER", "test"),
			Password: getEnv("DB_PASSWORD", "test"),
			Name:     getEnv("DB_NAME", "test"),
		},
		App: AppConfig{
			Name: getEnv("APP_NAME", "test"),
			Env:  getEnv("APP_ENV", "test"),
			Url:  getEnv("APP_URL", "test"),
		},
		Jwt: JwtConfig{
			Secret: getEnv("JWT_SECRET", "test"),
		},
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
