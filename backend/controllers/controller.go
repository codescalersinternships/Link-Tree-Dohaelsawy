package controllers

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	store    repository.Store
	Validate *validator.Validate
	Config   *model.Config
	Client   *s3.Client
}

func NewController(db *repository.DbInstance, config model.Config, client *s3.Client) *Controller {
	return &Controller{store: db, Validate: validator.New(), Config: &config, Client: client}
}
