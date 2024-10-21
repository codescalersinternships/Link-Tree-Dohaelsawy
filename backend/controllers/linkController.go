package controllers

import (
	"net/http"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type LinkController struct {
	Db       *gorm.DB
	Validate *validator.Validate
}

func NewLinksController(Db repository.DbInstance, validate *validator.Validate) *LinkController {
	return &LinkController{Db: Db.DB, Validate: validate}
}

func (l LinkController) CreateLink(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "worked",
	})

}

func (l LinkController) DeleteLink(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "worked",
	})
}

func (l LinkController) UpdateLink(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "worked",
	})
}

func (l LinkController) GetLinks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "worked",
	})
}
