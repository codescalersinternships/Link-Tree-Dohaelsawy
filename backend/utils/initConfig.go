package utils

import (
	"fmt"
	"os"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/models"
)

func NewConfigController() (model.Config, error) {

	fmt.Println(os.ReadDir("./"))
	fmt.Println(os.Environ())
	fmt.Println(getEnv("DB_HOST", "postgres-service"))

	config := model.Config{
		DbHost:             getEnv("DB_HOST", "postgres-service"),
		Origin:             getEnv("ALLOW_ORIGIN", "http://localhost:8020"),
		DbUser:             os.Getenv("DB_USER"),
		DbPassword:         os.Getenv("DB_PASSWORD"),
		DbName:             os.Getenv("DB_NAME"),
		DbPort:             os.Getenv("DB_PORT"),
		DB_CACHE_ADDR:      os.Getenv("DB_CACHE_ADDR"),
		DB_CACHE_PASSWORD:  os.Getenv("DB_CACHE_PASSWORD"),
		Port:               os.Getenv("PORT"),
		JwtSecret:          os.Getenv("JWT_SECRET"),
		TokenHourLifeTime:  os.Getenv("TOKEN_HOUR_LIFESPAN"),
		BaseUrl:            os.Getenv("BASE_URL"),
		UserImagePath:      os.Getenv("USERS_IMAGE_SAVE_PATH"),
		LinkTreePath:       os.Getenv("LINK_TREE_URL"),
		AwsRegion:          os.Getenv("AWS_REGION"),
		AwsAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}

	return config, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
