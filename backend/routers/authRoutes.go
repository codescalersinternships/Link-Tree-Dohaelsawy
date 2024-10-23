package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/gin-gonic/gin"
)

func AuthRouters(db repository.DbInstance, router *gin.Engine) {

	DBController := controllers.NewDBService(&db)

	routeGroup := router.Group("/auth")

	routeGroup.POST("/register", DBController.Register)
	routeGroup.POST("/login", DBController.Login)
	routeGroup.GET("/logout", DBController.Logout)
}
