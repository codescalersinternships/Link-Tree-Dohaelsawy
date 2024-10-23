package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

type LinkReq struct {
	Name string `validate:"required" json:"name"`
	Url  string `validate:"required" json:"url"`
}

func NewLinksController(Db repository.DbInstance, validate validator.Validate) *LinkController {
	return &LinkController{Db: &Db, Validate: &validate}
}

func (l *LinkController) CreateLink(ctx *gin.Context) {

	var reqBody LinkReq

	if err := ctx.BindJSON(&reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := l.Validate.Struct(reqBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Errorf("validation failed for field: %s", validationErrors[0].Field())
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, errorMessage)
		return
	}

	user_id, err := utils.ExtractTokenID(ctx)

	if err != nil {
		errorMessage := fmt.Errorf("can't find your token %s", err)
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, errorMessage)
		return
	}

	link := model.Link{
		Name:   reqBody.Name,
		Url:    reqBody.Url,
		UserID: user_id,
	}


	err = l.Db.AddNewLink(&link)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, link)
}

func (l *LinkController) DeleteLink(ctx *gin.Context) {

	var link model.Link

	idString := ctx.Params.ByName("link_id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = l.Db.DeleteLink(&link, id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, "deleted")
}

func (l *LinkController) UpdateLink(ctx *gin.Context) {

	var reqBody LinkReq

	if err := ctx.BindJSON(&reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := l.Validate.Struct(reqBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Errorf("validation failed for field: %s", validationErrors[0].Field())
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, errorMessage)
		return
	}

	idString := ctx.Params.ByName("link_id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	var link model.Link

	err = l.Db.GetOneLink(&link, id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	link.Url = reqBody.Url
	link.Name = reqBody.Name

	err = l.Db.PutOneLink(&link, link.ID)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, link)
}

func (l *LinkController) GetLinks(ctx *gin.Context) {

	var links []model.Link

	user_id, err := utils.ExtractTokenID(ctx)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = l.Db.GetAllLinksForUser(&links, user_id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, links)
}


