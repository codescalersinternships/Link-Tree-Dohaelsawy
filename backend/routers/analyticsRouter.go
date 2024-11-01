package routers

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/middleware"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/gin-gonic/gin"
)

func AnalyticsRouters(db repository.DbInstance, config model.Config, router *gin.Engine, client *s3.Client) {

	DBController := controllers.NewController(&db, config, client)

	routeGroup := router.Group("/analytics")

	routeGroup.Use(middleware.AuthMiddleware(config))
	routeGroup.GET("/get_analytics/:user_id", DBController.GetAnalytics)
}
