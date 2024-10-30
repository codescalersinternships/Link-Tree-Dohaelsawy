package controllers

import (
	"errors"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"slices"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
	"github.com/gin-gonic/gin"
)

type AccountReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Bio       string `json:"bio"`
}

type UserImageReq struct {
	Image *multipart.FileHeader `json:"image" form:"image"`
}

var (
	ErrImageTypeNotSupported = errors.New("we don't support this image extension, only png, jpj, jpeg")
)

//	@Summary		Delete Account
//	@Description	delete account
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Security		basic
//	@Success		200	{object}	SuccessResponse
//	@Failure		400	{object}	ErrResponse
//	@Failure		401	{object}	ErrResponse
//	@Failure		404	{object}	ErrResponse
//	@Failure		500	{object}	ErrResponse
//	@Router			/account/delete_account [delete]
func (ds *DBController) DeleteAccount(ctx *gin.Context) {

	var account model.User

	user_id, err := utils.ExtractTokenID(ctx, *ds.Config)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = ds.store.GetUserID(&account, user_id)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = ds.store.DeleteUser(&account, user_id)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	SuccessRespondJSON(ctx, http.StatusOK, "deleted")
}

//	@Summary		Edit Account
//	@Description	Edit Account data
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Param			AccountReq	body	AccountReq	true	"first name, last name, phone, bio"
//	@Security		basic
//	@Success		200	{object}	SuccessResponse
//	@Failure		400	{object}	ErrResponse
//	@Failure		401	{object}	ErrResponse
//	@Failure		404	{object}	ErrResponse
//	@Failure		500	{object}	ErrResponse
//	@Router			/account/edit_account [put]
func (ds *DBController) EditAccount(ctx *gin.Context) {

	var reqBody AccountReq

	if err := ctx.BindJSON(&reqBody); err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ds.Validate.Struct(&reqBody); err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	user_id, err := utils.ExtractTokenID(ctx, *ds.Config)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	var account model.User

	err = ds.store.GetUserID(&account, user_id)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	if !checkEmpty(reqBody.FirstName) {
		account.FirstName = reqBody.FirstName
	}
	if !checkEmpty(reqBody.LastName) {
		account.LastName = reqBody.LastName
	}
	if !checkEmpty(reqBody.Phone) {
		account.Phone = reqBody.Phone
	}
	if !checkEmpty(reqBody.Bio) {
		account.Bio = reqBody.Bio
	}

	err = ds.store.PutOneUser(&account, account.ID)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}
	account.Password = ""

	SuccessRespondJSON(ctx, http.StatusOK, gin.H{"user": account})
}

//	@Summary		Get Account
//	@Description	Get Account data
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Security		basic
//	@Success		200	{object}	SuccessResponse
//	@Failure		400	{object}	ErrResponse
//	@Failure		401	{object}	ErrResponse
//	@Failure		404	{object}	ErrResponse
//	@Failure		500	{object}	ErrResponse
//	@Router			/account/get_account [get]
func (ds *DBController) GetAccount(ctx *gin.Context) {

	var account model.User

	user_id, err := utils.ExtractTokenID(ctx, *ds.Config)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = ds.store.GetUserID(&account, user_id)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}
	account.Password = ""

	SuccessRespondJSON(ctx, http.StatusOK, gin.H{"user": account})
}

//	@Summary		Upload User Image
//	@Description	Upload User Image
//	@Tags			account
//	@Accept			mpfd
//	@Produce		json
//	@Param			UserImageReq	body	UserImageReq	true	"image"
//	@Security		basic
//	@Success		200	{object}	SuccessResponse
//	@Failure		400	{object}	ErrResponse
//	@Failure		401	{object}	ErrResponse
//	@Failure		404	{object}	ErrResponse
//	@Failure		500	{object}	ErrResponse
//	@Router			/account/add_photo [post]
func (ds *DBController) UploadUserImage(ctx *gin.Context) {

	file, err := ctx.FormFile("image")
	if err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	config := ds.Config
	user_id, err := utils.ExtractTokenID(ctx, *config)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	var account model.User
	err = ds.store.GetUserID(&account, user_id)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	extension := filepath.Ext(file.Filename)
	imageType := []string{".png", ".jpg", "jpeg"}

	if ok := slices.Contains(imageType, extension); !ok {
		ErrRespondJSON(ctx, http.StatusBadRequest, ErrImageTypeNotSupported)
		return
	}

	newFileName := ds.Config.UserImagePath + account.Username + extension
	if err := ctx.SaveUploadedFile(file, newFileName); err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	account.Image = newFileName

	err = ds.store.PutOneUser(&account, account.ID)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}
	account.Password = ""

	SuccessRespondJSON(ctx, http.StatusOK, gin.H{"user": account})
}
