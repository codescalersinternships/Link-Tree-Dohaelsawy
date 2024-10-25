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
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (ds *DBController) CreateLink(ctx *gin.Context) {

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

func (ds *DBController) DeleteLink(ctx *gin.Context) {

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

func (ds *DBController) UpdateLink(ctx *gin.Context) {

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

func (ds *DBController) GetLinks(ctx *gin.Context) {

	var user model.User
	username := ctx.Params.ByName("username")

	if err := ds.store.GetUserUsername(&user,username); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	var links []model.Link
	if err := ds.store.GetAllLinksForUser(&links, user.ID); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	getGuestUsername(ctx,ds,user.ID)

	utils.SuccessRespondJSON(ctx, http.StatusOK, links)
}


func getGuestUsername(ctx *gin.Context,ds *DBController, user_id int) {

	guest_id, err := utils.ExtractTokenID(ctx, *ds.Config)
	if err != nil{
		ds.CalculateAnalytics(ctx,"unknown",user_id)
		return
	}

	if guest_id == user_id {
		return
	}

	var guest model.User
	if err := ds.store.GetUserID(&guest,guest_id); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	ds.CalculateAnalytics(ctx,guest.Username,user_id)
}