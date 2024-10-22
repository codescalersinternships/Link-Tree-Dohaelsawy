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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.Validate.Struct(reqBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	var existingUser model.User

	err := ac.Db.GetUserEmail(&existingUser, reqBody.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}

	valid := utils.ComparePassword(reqBody.Password, existingUser.Password)

	if !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password invalid"})
		return
	}

	token, err := utils.CreateToken(uint(existingUser.ID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("JWT Error %s", err.Error())})
		return
	}

	existingUser.Token = token

	if err := ac.Db.PutOneUser(&existingUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("JWT Error %s", err.Error())})
		return
	}
	err = godotenv.Load(".env")

	if err != nil {
		log.Printf("Error loading .env file, %s\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error loading .env file, %s\n", err)})
		return
	}

	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error convert token life time to number, %s\n", err)})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token, 3600 * token_lifespan, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{"access_token": token})
}

func (ac AuthController) Register(ctx *gin.Context) {
	var reqBody RegisterRequest
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.Validate.Struct(reqBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	var existingUser model.User

	err := ac.Db.GetUserEmail(&existingUser, reqBody.Email)

	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	password, err := utils.EncryptPassword(reqBody.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
