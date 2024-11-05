package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/database/repository"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/middleware"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/models"
	"github.com/gin-gonic/gin"
)

func AnalyticsRouters(db repository.DbInstance, config model.Config, router *gin.Engine) {

	Controller := controllers.NewController(&db, config)

	routeGroup := router.Group("/analytics")

	routeGroup.Use(middleware.AuthMiddleware(config))
	routeGroup.GET("/get_analytics/:user_id", Controller.GetAnalytics)
}
