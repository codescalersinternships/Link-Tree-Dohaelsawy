package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound = errors.New("not found")
)

type LinkReq struct {
	Name string `validate:"required" json:"name"`
	Url  string `validate:"required" json:"url"`
}

func (ds *DBService) CreateLink(ctx *gin.Context) {

	var reqBody LinkReq

	if err := ctx.BindJSON(&reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ds.Validate.Struct(reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	user_id, err := utils.ExtractTokenID(ctx, *ds.Config)

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

	err = ds.store.AddNewLink(&link)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, link)
}

func (ds *DBService) DeleteLink(ctx *gin.Context) {

	var link model.Link

	idString := ctx.Params.ByName("link_id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = ds.store.DeleteLink(&link, id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, "deleted")
}

func (ds *DBService) UpdateLink(ctx *gin.Context) {

	var reqBody LinkReq

	if err := ctx.BindJSON(&reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ds.Validate.Struct(reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	idString := ctx.Params.ByName("link_id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	var link model.Link

	err = ds.store.GetOneLink(&link, id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	link.Url = reqBody.Url
	link.Name = reqBody.Name

	err = ds.store.PutOneLink(&link, link.ID)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, link)
}

func (ds *DBService) GetLinks(ctx *gin.Context) {

	var links []model.Link

	user_id, err := utils.ExtractTokenID(ctx, *ds.Config)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = ds.store.GetAllLinksForUser(&links, user_id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, links)
}
