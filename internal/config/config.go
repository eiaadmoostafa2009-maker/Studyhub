package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	SecretKey     string
	DatabasePath  string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	secretKey := os.Getenv("SECRET_JWT")
	if secretKey == "" {
		return nil, fmt.Errorf("SECRET_JWT is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	databasePath := os.Getenv("DATABASE_PATH")
    if databasePath == "" {
       databasePath = "studyhub.db"
    }

    return &Config{
      Port:         port,
      SecretKey:    secretKey,
      DatabasePath: databasePath,
    }, nil
}