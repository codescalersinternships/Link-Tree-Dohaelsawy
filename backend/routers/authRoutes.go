package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/database/repository"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/middleware"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/models"
	"github.com/gin-gonic/gin"
)

func AuthRouters(db repository.DbInstance, config model.Config, router *gin.Engine) {

	Controller := controllers.NewController(&db, config)

	routeGroup := router.Group("/auth")

	routeGroup.POST("/register", Controller.Register)
	routeGroup.POST("/login", Controller.Login)
	routeGroup.GET("/logout", middleware.AuthMiddleware(config), Controller.Logout)
}
