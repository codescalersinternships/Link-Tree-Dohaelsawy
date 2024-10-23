package repository

import (
	"fmt"
	"log"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
	"github.com/go-playground/validator/v10"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

type Store interface {
	AddNewLink(l *model.Link) (err error)
	AddNewUser(u *model.User) (err error)
	DeleteLink(l *model.Link, id int) (err error)
	DeleteUser(u *model.User, id int) (err error)
	GetAllLinksForUser(l *[]model.Link, user_id int) (err error)
	GetOneLink(l *model.Link, id int) (err error)
	GetUserEmail(u *model.User, email string) (err error)
	GetUserID(u *model.User, id int) (err error)
	PutOneLink(l *model.Link, id int) (err error)
	PutOneUser(u *model.User, id int) (err error)
}

func NewDbInstance(db *gorm.DB) DbInstance {
	return DbInstance{DB: db}
}

func DbConnect() (*gorm.DB, error) {

	conString, err := prepareDbConnectionString()

	if err != nil {
		log.Printf("error: %s", err)
		return nil, err
	}

	fmt.Println(conString)

	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})
	if err != nil {
		log.Printf("error: %s", err)
		return nil, err
	}

	err = db.AutoMigrate(&model.User{}, &model.Link{})
	if err != nil {
		return nil, err
	}

	log.Println("Database Migration Completed!")

	return db, nil
}

func prepareDbConnectionString() (string, error) {

	config, err := utils.NewConfigController()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort),
		nil
}
