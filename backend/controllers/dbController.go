package controllers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/go-playground/validator/v10"
)

type DBService struct {
	store repository.Store
	Validate *validator.Validate
}

func NewDBService(db *repository.DbInstance) *DBService {
	return &DBService{store: db,Validate: validator.New()}
}
