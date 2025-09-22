package config

import (
	"fmt"
	"os"

	"github.com/Dionisius951/Golang-Blog-BE/internal/models"
	"github.com/joho/godotenv"
)

func LoadConfig() *models.Config{
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, using system environment variables")
	}

	config := &models.Config{
		APP_PORT: os.Getenv("APP_PORT"),
		DB_URL: os.Getenv("DB_URL"),
	}

	return config
}
