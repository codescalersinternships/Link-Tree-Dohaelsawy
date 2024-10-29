package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterRequest struct {
	FirstName string `validate:"required" json:"first_name"`
	LastName  string `validate:"required" json:"last_name"`
	Username  string `validate:"required" json:"username"`
	Email     string `validate:"required,email" json:"email"`
	Password  string `validate:"required,min=8" json:"password"`
}

var (
	ErrEmailExist    = errors.New("this email exist! ")
	ErrWrongPassword = errors.New("wrong password!! ")
)

//	@Summary		Login
//	@Description	returns users token after login
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			LoginRequest	body		LoginRequest	true	"email and password"
//	@Success		200				{object}	SuccessResponse
//	@Failure		400				{object}	ErrResponse
//	@Failure		401				{object}	ErrResponse
//	@Failure		404				{object}	ErrResponse
//	@Failure		500				{object}	ErrResponse
//	@Router			/auth/login [post]
func (ds *DBController) Login(ctx *gin.Context) {

	config := ds.Config

	secretToken := config.JwtSecret

	tokenLifeTime, err := strconv.Atoi(config.TokenHourLifeTime)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	var reqBody LoginRequest

	if err := ctx.BindJSON(&reqBody); err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ds.Validate.Struct(&reqBody); err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	var existingUser model.User

	err = ds.store.GetUserEmail(&existingUser, reqBody.Email)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusNotFound, err)
		return
	}

	valid := utils.ComparePassword(reqBody.Password, existingUser.Password)

	if !valid {
		ErrRespondJSON(ctx, http.StatusUnauthorized, ErrWrongPassword)
		return
	}

	token, err := utils.CreateToken(uint(existingUser.ID), tokenLifeTime, secretToken)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	existingUser.Token = token

	if err := ds.store.PutOneUser(&existingUser, existingUser.ID); err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)

	ctx.SetCookie("Authorization", token, 3600*tokenLifeTime, "", "", false, true)

	ctx.Header("Authorization", token)

	existingUser.Password = ""

	SuccessRespondJSON(ctx, http.StatusOK, gin.H{"access_token": token, "user": existingUser})
}

//	@Summary		Register
//	@Description	creates user account
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			RegisterRequest	body		RegisterRequest	true	"first name, last name, email, username and password"
//	@Success		200				{object}	SuccessResponse
//	@Failure		400				{object}	ErrResponse
//	@Failure		401				{object}	ErrResponse
//	@Failure		404				{object}	ErrResponse
//	@Failure		500				{object}	ErrResponse
//	@Router			/auth/register [post]
func (ds *DBController) Register(ctx *gin.Context) {

	var reqBody RegisterRequest
	if err := ctx.BindJSON(&reqBody); err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ds.Validate.Struct(reqBody); err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	var existingUser model.User

	err := ds.store.GetUserEmail(&existingUser, reqBody.Email)

	if err == nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, ErrEmailExist)
		return
	}

	password, err := utils.EncryptPassword(reqBody.Password)

	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	newUser := model.User{
		FirstName:   reqBody.FirstName,
		LastName:    reqBody.LastName,
		Username:    reqBody.Username,
		Email:       reqBody.Email,
		Password:    password,
		LinkTreeURL: utils.GenerateLinkTreeUrl(ds.Config.BaseUrl,ds.Config.LinkTreePath, reqBody.Username ),
	}
	
	if err := ds.store.AddNewUser(&newUser); err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	SuccessRespondJSON(ctx, http.StatusOK, gin.H{"message": "User registered successfully"})
}

//	@Summary		Logout
//	@Description	removes user token from cookies
//	@Tags			auth
//	@Accept			json
//	@Security		basic
//	@Success		200	{object}	SuccessResponse
//	@Router			/auth/logout [get]
func (ac *DBController) Logout(ctx *gin.Context) {
	ctx.SetCookie("Authorization", "", 0, "", "", false, true)
}
