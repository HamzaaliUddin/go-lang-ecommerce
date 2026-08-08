package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort     string
	AppEnv      string
	DatabaseURL string
	JWTSecret   string

	SuperAdminFirstName string
	SuperAdminLastName  string
	SuperAdminEmail     string
	SuperAdminPassword  string
}

func Load() Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: .env file not found")
	}

	return Config{
		AppPort:     os.Getenv("APP_PORT"),
		AppEnv:      os.Getenv("APP_ENV"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		SuperAdminFirstName: os.Getenv("SUPER_ADMIN_FIRST_NAME"),
		SuperAdminLastName:  os.Getenv("SUPER_ADMIN_LAST_NAME"),
		SuperAdminEmail:     os.Getenv("SUPER_ADMIN_EMAIL"),
		SuperAdminPassword:  os.Getenv("SUPER_ADMIN_PASSWORD"),
	}
}