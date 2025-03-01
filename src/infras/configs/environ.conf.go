package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	APP_NAME string
	APP_PORT string
	APP_VER  string

	DB_HOST string
	DB_NAME string
	DB_PORT string
	DB_USER string
	DB_PASS string

	JWT_SECRET string
}

func LoadConfig() *Configs {
	// check environment by dynamic type
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	// load .env file
	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Configs{
		APP_NAME: os.Getenv("APP_NAME"),
		APP_PORT: os.Getenv("APP_PORT"),
		APP_VER:  os.Getenv("APP_VER"),

		DB_HOST: os.Getenv("DB_HOST"),
		DB_NAME: os.Getenv("DB_NAME"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASS"),

		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}
}
