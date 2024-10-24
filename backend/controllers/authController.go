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
	Email     string `validate:"required" json:"email"`
	Password  string `validate:"required,min=8" json:"password"`
}

var (
	ErrEmailExist    = errors.New("this email exist! ")
	ErrWrongPassword = errors.New("wrong password!! ")
)

func (ds *DBService) Login(ctx *gin.Context) {

	config := ds.Config

	secretToken := config.JwtSecret

	tokenLifeTime, err := strconv.Atoi(config.TokenHourLifeTime)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	var reqBody LoginRequest

	if err := ctx.BindJSON(&reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ds.Validate.Struct(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	var existingUser model.User

	err = ds.store.GetUserEmail(&existingUser, reqBody.Email)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusNotFound, err)
		return
	}

	valid := utils.ComparePassword(reqBody.Password, existingUser.Password)

	if !valid {
		utils.ErrRespondJSON(ctx, http.StatusUnauthorized, ErrWrongPassword)
		return
	}

	token, err := utils.CreateToken(uint(existingUser.ID), tokenLifeTime, secretToken)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	existingUser.Token = token

	if err := ds.store.PutOneUser(&existingUser, existingUser.ID); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token, 3600*tokenLifeTime, "", "", false, true)

	utils.SuccessRespondJSON(ctx, http.StatusOK, gin.H{"access_token": token})
}

func (ds *DBService) Register(ctx *gin.Context) {

	var reqBody RegisterRequest
	if err := ctx.BindJSON(&reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ds.Validate.Struct(reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	var existingUser model.User

	err := ds.store.GetUserEmail(&existingUser, reqBody.Email)

	if err == nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, ErrEmailExist)
		return
	}

	password, err := utils.EncryptPassword(reqBody.Password)

	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	newUser := model.User{
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
		Username:  reqBody.Username,
		Email:     reqBody.Email,
		Password:  password,
	}

	if err := ds.store.AddNewUser(&newUser); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (ac *DBService) Logout(ctx *gin.Context) {
	ctx.SetCookie("Authorization", "", 0, "", "", false, true)
}
