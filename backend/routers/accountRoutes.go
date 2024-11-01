package routers

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/middleware"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/gin-gonic/gin"
)

func AccountRouters(db repository.DbInstance, config model.Config, router *gin.Engine, client *s3.Client) {

	DBController := controllers.NewController(&db,config, client)

	routeGroup := router.Group("/account")

	routeGroup.Use(middleware.AuthMiddleware(config))
	routeGroup.PUT("/edit_account", DBController.EditAccount)
	routeGroup.DELETE("/delete_account", DBController.DeleteAccount)
	routeGroup.GET("/get_account", DBController.GetAccount)
	routeGroup.POST("/add_photo",DBController.UploadUserImage)

}
