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

type AccountController struct {
	Db       *repository.DbInstance
	Validate *validator.Validate
}

type AccountReq struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone" validate:"min=11,max=11"`
	Photo       string `json:"photo"`
	Bio         string `json:"bio"`
}

func NewAccountControllerImpl(Db repository.DbInstance, validate validator.Validate) *AccountController {
	return &AccountController{Db: &Db, Validate: &validate}
}

func (a *AccountController) DeleteAccount(ctx *gin.Context) {

	var account model.User

	idString := ctx.Params.ByName("user_id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = a.Db.DeleteUser(&account, id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, "deleted")
}

func (a *AccountController) EditAccount(ctx *gin.Context) {

	var reqBody AccountReq


	if err := ctx.BindJSON(&reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, errors.New("here"))
		return
	}

	if err := a.Validate.Struct(&reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	idString := ctx.Params.ByName("user_id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	var account model.User

	err = a.Db.GetUserID(&account, id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	account.FirstName = reqBody.FirstName
	account.LastName = reqBody.LastName
	account.Phone = reqBody.Phone
	account.Photo = reqBody.Photo             // TODO: need to handle upload image
	account.Bio = reqBody.Bio

	err = a.Db.PutOneUser(&account, account.ID)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, account)
}

func (a *AccountController) GetAccount(ctx *gin.Context) {

	var account model.User

	user_id, err := utils.ExtractTokenID(ctx)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = a.Db.GetUserID(&account, user_id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, account)
}

func (a AccountController) CreateLinkTreeUrl(ctx *gin.Context) {

	user_id, err := utils.ExtractTokenID(ctx)

	if err != nil {
		errorMessage := fmt.Errorf("can't find your token %s", err)
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, errorMessage)
		return
	}

	var account model.User


	err = a.Db.GetUserID(&account, user_id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	account.LinkTreeURL, err = utils.GenerateLinkTreeUrl(account.Username)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = a.Db.PutOneUser(&account, account.ID)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, account)
}