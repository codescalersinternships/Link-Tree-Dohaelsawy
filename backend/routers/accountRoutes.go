package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/middleware"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/gin-gonic/gin"
)

func AccountRouters(db repository.DbInstance, config model.Config, router *gin.Engine) {

	DBController := controllers.NewDBService(&db,config)

	routeGroup := router.Group("/account")

	routeGroup.Use(middleware.AuthMiddleware(*DBController.Config))
	routeGroup.POST("/edit_account", DBController.EditAccount)
	routeGroup.DELETE("/delete_account", DBController.DeleteAccount)
	routeGroup.GET("/get_account", DBController.GetAccount)
	routeGroup.GET("/create_link_tree_url", DBController.CreateLinkTreeUrl)

}
