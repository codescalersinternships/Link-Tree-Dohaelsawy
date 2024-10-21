package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

func AuthRouter(db repository.DbInstance) *gin.Engine {

	validate := validator.New()
	authController := controllers.NewAuthControllerImpl(db, validate)
	service := gin.Default()

	router := service.Group("/auth")

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	return service
}


