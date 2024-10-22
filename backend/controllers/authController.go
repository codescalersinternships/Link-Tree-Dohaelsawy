package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
)

type LoginRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type RegisterRequest struct {
	FirstName string `validate:"required" json:"first_name"`
	LastName  string `validate:"required" json:"last_name"`
	Username  string `validate:"required" json:"username"`
	Email     string `validate:"required" json:"email"`
	Password  string `validate:"required" json:"password"`
}

type AuthController struct {
	Db       *repository.DbInstance
	Validate *validator.Validate
}

func NewAuthControllerImpl(Db repository.DbInstance, validate *validator.Validate) *AuthController {
	return &AuthController{Db: &Db, Validate: validate}
}

func (ac AuthController) Login(ctx *gin.Context) {
	var reqBody LoginRequest

	if err := ctx.BindJSON(&reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ac.Validate.Struct(reqBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Errorf("validation failed for field: %s", validationErrors[0].Field())
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, errorMessage)
		return
	}

	var existingUser model.User

	err := ac.Db.GetUserEmail(&existingUser, reqBody.Email)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusNotFound, err)
		return
	}

	valid := utils.ComparePassword(reqBody.Password, existingUser.Password)

	if !valid {
		utils.ErrRespondJSON(ctx, http.StatusUnauthorized, err)
		return
	}

	token, err := utils.CreateToken(uint(existingUser.ID))

	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	existingUser.Token = token

	if err := ac.Db.PutOneUser(&existingUser); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}
	err = godotenv.Load(".env")

	if err != nil {
		log.Printf("Error loading .env file, %s\n", err)
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token, 3600*token_lifespan, "", "", false, true)

	utils.SuccessRespondJSON(ctx, http.StatusOK, gin.H{"access_token": token})

}

func (ac AuthController) Register(ctx *gin.Context) {
	var reqBody RegisterRequest
	if err := ctx.BindJSON(&reqBody); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ac.Validate.Struct(reqBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Errorf("validation failed for field: %s", validationErrors[0].Field())
		utils.ErrRespondJSON(ctx, http.StatusBadRequest, errorMessage)
		return
	}

	var existingUser model.User

	err := ac.Db.GetUserEmail(&existingUser, reqBody.Email)

	if err == nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
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

	if err := ac.Db.AddNewUser(&newUser); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (ac AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie("Authorization", "", 0, "", "", false, true)
}
