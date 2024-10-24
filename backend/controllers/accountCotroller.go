package controllers

import (
	"fmt"
	"net/http"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
	"github.com/gin-gonic/gin"
)

type AccountReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone" validate:"min=11,max=11"`
	Photo     string `json:"photo"`
	Bio       string `json:"bio"`
}

func (ds *DBService) DeleteAccount(ctx *gin.Context) {

	var account model.User

	user_id, err := utils.ExtractTokenID(ctx, *ds.Config)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = ds.store.DeleteUser(&account, user_id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, "deleted")
}

func (ds *DBService) EditAccount(ctx *gin.Context) {

	var reqBody AccountReq

	if err := ctx.BindJSON(&reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ds.Validate.Struct(&reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	user_id, err := utils.ExtractTokenID(ctx, *ds.Config)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	var account model.User

	err = ds.store.GetUserID(&account, user_id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	account.FirstName = reqBody.FirstName
	account.LastName = reqBody.LastName
	account.Phone = reqBody.Phone
	account.Photo = reqBody.Photo // TODO: need to handle upload image
	account.Bio = reqBody.Bio

	err = ds.store.PutOneUser(&account, account.ID)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, account)
}

func (ds *DBService) GetAccount(ctx *gin.Context) {

	var account model.User

	user_id, err := utils.ExtractTokenID(ctx, *ds.Config)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = ds.store.GetUserID(&account, user_id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, account)
}

func (ds *DBService) CreateLinkTreeUrl(ctx *gin.Context) {

	config := ds.Config

	user_id, err := utils.ExtractTokenID(ctx, *ds.Config)

	if err != nil {
		errorMessage := fmt.Errorf("can't find your token %s", err)
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, errorMessage)
		return
	}

	var account model.User

	err = ds.store.GetUserID(&account, user_id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	account.LinkTreeURL = utils.GenerateLinkTreeUrl(account.Username, config.BaseUrl)

	err = ds.store.PutOneUser(&account, account.ID)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, account)
}
