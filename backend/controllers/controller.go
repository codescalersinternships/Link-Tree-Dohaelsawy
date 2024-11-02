package controllers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	store    repository.Store
	Validate *validator.Validate
	Config   *model.Config
}

func NewController(db *repository.DbInstance, config model.Config) *Controller {
	return &Controller{store: db, Validate: validator.New(), Config: &config}
}
