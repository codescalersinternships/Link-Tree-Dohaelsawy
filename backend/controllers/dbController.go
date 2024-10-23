package controllers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/go-playground/validator/v10"
)

type DBController struct {
	Db       *repository.DbInstance
	Validate *validator.Validate
}

func NewDBControllerImpl(Db repository.DbInstance, validate validator.Validate) *DBController {
	return &DBController{Db: &Db, Validate: &validate}
}
