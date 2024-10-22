package repository

import (
	"fmt"
	"log"
	"os"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	DB *gorm.DB
}

func DbConnect() (DbInstance, error) {
	conString := prepareDbConnectionString()
	print(conString)
	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		log.Printf("error: %s", err)
		return DbInstance{}, err
	}
	log.Println("Connected to Database!")
	return DbInstance{DB: db}, nil
}

func (db DbInstance) Migrate() error {

	err := db.DB.AutoMigrate(&model.User{}, &model.Link{})
	if err != nil {
		return err
	}

	log.Println("Database Migration Completed!")
	return nil
}

func prepareDbConnectionString() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Error loading .env file, %s\n", err)
		return err.Error()
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", host, user, password, dbName, port)
}
