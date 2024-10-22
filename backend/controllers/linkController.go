package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	ErrNotFound = errors.New("not found")
)

type LinkController struct {
	Db       *repository.DbInstance
	Validate *validator.Validate
}

type CreateLinkReq struct {
	Name       string `validate:"required" json:"name"`
	Url        string `validate:"required" json:"url"`
}

func NewLinksController(Db *repository.DbInstance, validate *validator.Validate) *LinkController {
	return &LinkController{Db: Db, Validate: validate}
}

func (l *LinkController) CreateLink(ctx *gin.Context) {

	var reqBody CreateLinkReq

	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := l.Validate.Struct(reqBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}


	user_id, err := utils.ExtractTokenID(ctx)

	if err != nil {
		errorMessage := fmt.Sprintf("can't find your token %s",err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	link := model.Link{
		Name: reqBody.Name,
		Url:  reqBody.Url,
		UserID: user_id,
	}

	err = l.Db.AddNewLink(&link)

	if err != nil {
		utils.ErrRespondJSON(ctx, 404, err)
		return
	}

	utils.SuccessRespondJSON(ctx, 200, link)
}

func (l *LinkController) DeleteLink(ctx *gin.Context) {
	var link model.Link
	id := ctx.Params.ByName("id")
	err := l.Db.DeleteLink(&link, id)

	if err != nil {
		utils.ErrRespondJSON(ctx, 404, err)
		return
	}

	utils.SuccessRespondJSON(ctx, 200, link)
}

func (l *LinkController) UpdateLink(ctx *gin.Context) {

	var link model.Link
	id := ctx.Params.ByName("update_id")
	err := l.Db.GetOneLink(&link, id)

	if err != nil {
		utils.ErrRespondJSON(ctx, 404, err)
		return
	}

	ctx.BindJSON(&link)
	err = l.Db.PutOneLink(&link, id)

	if err != nil {
		utils.ErrRespondJSON(ctx, 404, err)
		return
	}

	utils.SuccessRespondJSON(ctx, 200, link)
}

func (l *LinkController) GetLinks(ctx *gin.Context) {

	var links []model.Link
	user_id := 8
	err := l.Db.GetAllLinksForUser(&links, string(user_id))

	if err != nil {
		utils.ErrRespondJSON(ctx, 404, err)
		return
	}

	utils.SuccessRespondJSON(ctx, 200, links)
}
