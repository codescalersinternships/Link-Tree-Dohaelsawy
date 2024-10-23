package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AuthRouters(db repository.DbInstance, router *gin.Engine) {

	validate := validator.New(validator.WithRequiredStructEnabled())
	DBController := controllers.NewDBControllerImpl(db, *validate)

	routeGroup := router.Group("/auth")

	routeGroup.POST("/register", DBController.Register)
	routeGroup.POST("/login", DBController.Login)
	routeGroup.GET("/logout", DBController.Logout)
}
