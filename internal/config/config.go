package config

import (
	"github.com/joho/godotenv"
	//"github.com/go-sql-driver/mysql"
	"fmt"
	"os"
)

//initiaalize config struct
type Config struct{
	Port string
	SecretKey string
}
//pulling func
func LoadConfig() (*Config, error){
	err := godotenv.Load()
	if err !=nil {
		return nil,fmt.Errorf("failed to load .env file")
	}

	fmt.Println("config loaded")

	return &Config{
	Port: os.Getenv("PORT"),
	SecretKey: os.Getenv("SECRET_JWT"),
	}, nil
}