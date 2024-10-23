package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GenerateLinkTreeUrl(username string) (string, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	
	baseURL := os.Getenv("BASE_URL")

	return fmt.Sprintf("%s/%s", baseURL,username), nil
}
