package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/gin-gonic/gin"
)

func AuthRouters(db repository.DbInstance, config model.Config, router *gin.Engine) {

	DBController := controllers.NewDBService(&db, config)

	routeGroup := router.Group("/auth")

	routeGroup.POST("/register", DBController.Register)
	routeGroup.POST("/login", DBController.Login)
	routeGroup.GET("/logout", DBController.Logout)
}
