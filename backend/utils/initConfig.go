package utils

import (
	"fmt"
	"os"
	"strings"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/models"
)

func NewConfigController() (model.Config, error) {
	// pwd, err := os.
	// fmt.Println(pwd)
	// fmt.Println(err)
	// if err != nil {
	// 	return model.Config{}, err
	// }

	fmt.Println(os.ReadDir("./"))

	err := readFile(".env")
	if err != nil {
		return model.Config{}, err
	}

	return model.Config{
		DbHost:             os.Getenv("DB_HOST"),
		DbUser:             os.Getenv("DB_USER"),
		DbPassword:         os.Getenv("DB_PASSWORD"),
		DbName:             os.Getenv("DB_NAME"),
		DbPort:             os.Getenv("DB_PORT"),
		Port:               os.Getenv("PORT"),
		JwtSecret:          os.Getenv("JWT_SECRET"),
		TokenHourLifeTime:  os.Getenv("TOKEN_HOUR_LIFESPAN"),
		BaseUrl:            os.Getenv("BASE_URL"),
		UserImagePath:      os.Getenv("USERS_IMAGE_SAVE_PATH"),
		LinkTreePath:       os.Getenv("LINK_TREE_URL"),
		StaticImagesPath:   os.Getenv("STATIC_IMAGES"),
		AwsRegion:          os.Getenv("AWS_REGION"),
		AwsAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}, nil
}

func readFile(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return parse(string(file))
}

func parse(file string) error {

	lines := strings.Split(file, "\n")

	for _, line := range lines {

		iniLine := strings.TrimSpace(line)

		key, value, ok := strings.Cut(iniLine, "=")
		if ok {
			os.Setenv(key, value)
		}
	}
	return nil
}
