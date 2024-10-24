package utils

import (
	"os"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/joho/godotenv"
)

func NewConfigController() (model.Config, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return model.Config{}, err
	}

	return model.Config{
		DbHost:            os.Getenv("DB_HOST"),
		DbUser:            os.Getenv("DB_USER"),
		DbPassword:        os.Getenv("DB_PASSWORD"),
		DbName:            os.Getenv("DB_NAME"),
		DbPort:            os.Getenv("DB_PORT"),
		Port:              os.Getenv("PORT"),
		JwtSecret:         os.Getenv("JWT_SECRET"),
		TokenHourLifeTime: os.Getenv("TOKEN_HOUR_LIFESPAN"),
		BaseUrl:           os.Getenv("BASE_URL"),
	}, nil
}