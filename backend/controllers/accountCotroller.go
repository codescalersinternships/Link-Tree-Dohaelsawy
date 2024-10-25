package controllers

import (
	"mime/multipart"
	"net/http"
	"path/filepath"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
	"github.com/gin-gonic/gin"
)

type AccountReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone" validate:"min=11,max=11"`
	Bio       string `json:"bio"`
}

type UserImageReq struct {
	Image *multipart.FileHeader `form:"image"`
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

	user_id, err := utils.ExtractTokenID(ctx, *config)

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

	account.LinkTreeURL = utils.GenerateLinkTreeUrl(account.Username, config.BaseUrl)

	err = ds.store.PutOneUser(&account, account.ID)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, account)
}

func (ds *DBService) UploadUserImage(ctx *gin.Context) {

	file, err := ctx.FormFile("image")
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	config := ds.Config
	user_id, err := utils.ExtractTokenID(ctx, *config)
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

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := ds.Config.UserImagePath + account.Username + extension

	// The file is received, so let's save it
	if err := ctx.SaveUploadedFile(file, newFileName); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	account.Image = newFileName

	err = ds.store.PutOneUser(&account, account.ID)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, gin.H{
		"user_image_path": newFileName,
	})
}
