package repository

import (
	"fmt"
	"log"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	DB *gorm.DB
}

func DbConnect() (DbInstance, error) {

	conString, err := prepareDbConnectionString()

	if err != nil {
		log.Printf("error: %s", err)
		return DbInstance{}, err
	}

	fmt.Println(conString)

	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})
	if err != nil {
		log.Printf("error: %s", err)
		return DbInstance{}, err
	}

	err = db.AutoMigrate(&model.User{}, &model.Link{})
	if err != nil {
		return DbInstance{}, err
	}

	log.Println("Database Migration Completed!")

	return DbInstance{DB: db}, nil
}

func prepareDbConnectionString() (string, error) {

	config, err := utils.NewConfigController()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort),
		nil
}
