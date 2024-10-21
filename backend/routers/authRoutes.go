package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

func AuthRouters(db repository.DbInstance, router *gin.Engine) {

	validate := validator.New()
	authController := controllers.NewAuthControllerImpl(db, validate)
	
	routeGroup :=router.Group("/auth")

	routeGroup.POST("/register", authController.Register)
	routeGroup.POST("/login", authController.Login)
}


