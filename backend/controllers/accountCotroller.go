package controllers

import (
	"context"
	"errors"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"slices"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
	ErrCouldNotOpenImage     = errors.New("couldn't open file")
	ErrCouldNotUploadImage   = errors.New("couldn't upload file")
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
func (c *Controller) DeleteAccount(ctx *gin.Context) {

	var account model.User

	user_id, err := utils.ExtractTokenID(ctx, *c.Config)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = c.store.GetUserID(&account, user_id)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = c.store.DeleteUser(&account, user_id)
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
func (c *Controller) EditAccount(ctx *gin.Context) {

	var reqBody AccountReq

	if err := ctx.BindJSON(&reqBody); err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := c.Validate.Struct(&reqBody); err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	user_id, err := utils.ExtractTokenID(ctx, *c.Config)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	var account model.User

	err = c.store.GetUserID(&account, user_id)
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

	err = c.store.PutOneUser(&account, account.ID)
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
func (c *Controller) GetAccount(ctx *gin.Context) {

	var account model.User

	user_id, err := utils.ExtractTokenID(ctx, *c.Config)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = c.store.GetUserID(&account, user_id)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}
	account.Password = ""

	SuccessRespondJSON(ctx, http.StatusOK, gin.H{"user": account})
}


//	@Summary		Get Account by username
//	@Description	Get Account data
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Param			username	path	string	true	"username"
//	@Security		basic
//	@Success		200	{object}	SuccessResponse
//	@Failure		400	{object}	ErrResponse
//	@Failure		401	{object}	ErrResponse
//	@Failure		404	{object}	ErrResponse
//	@Failure		500	{object}	ErrResponse
//	@Router			/account/{username} [get]
func (c *Controller) GetAccountByUsername(ctx *gin.Context) {

	var account model.User
	username := ctx.Params.ByName("username")

	err := c.store.GetUserUsername(&account, username)
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
func (c *Controller) UploadUserImage(ctx *gin.Context) {

	user_id, err := utils.ExtractTokenID(ctx, *c.Config)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	var account model.User
	err = c.store.GetUserID(&account, user_id)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	extension := filepath.Ext(file.Filename)
	imageType := []string{".png", ".jpg", ".jpeg"}

	ok := slices.Contains(imageType, extension)
	if !ok {
		ErrRespondJSON(ctx, http.StatusBadRequest, ErrImageTypeNotSupported)
		return
	}

	newFileName := account.Username + extension

	openFile, err := file.Open()
	if err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	client, err := prepareAwsClient(c)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("link-tree"),
		Key:    aws.String(newFileName),
		Body:   openFile,
		ACL:    "public-read",
	})
	if err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	account.Image = result.Location

	err = c.store.PutOneUser(&account, account.ID)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}
	account.Password = ""

	SuccessRespondJSON(ctx, http.StatusOK, gin.H{"user": account})
}

//	@Summary		returns all usernames
//	@Description	returns all usernames
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Security		basic
//	@Success		200	{object}	SuccessResponse
//	@Failure		400	{object}	ErrResponse
//	@Failure		401	{object}	ErrResponse
//	@Failure		404	{object}	ErrResponse
//	@Failure		500	{object}	ErrResponse
//	@Router			/account/search [get]
func (c *Controller) UsernameSearch(ctx *gin.Context) {
	var users []model.User
	var usernames []string

	err := c.store.GetAllUsers(&users)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	for _, user := range users{
		usernames = append(usernames, user.Username)
	}

	SuccessRespondJSON(ctx, http.StatusOK, gin.H{"usernames": usernames})
}

func prepareAwsClient(c *Controller) (*s3.Client, error) {
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithRegion(c.Config.AwsRegion))
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	return client, nil
}
