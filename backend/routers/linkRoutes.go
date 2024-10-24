package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/middleware"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"

	"github.com/gin-gonic/gin"
)

func LinkRouters(db repository.DbInstance, config model.Config, router *gin.Engine) {

	DBController := controllers.NewDBService(&db, config)

	routeGroup := router.Group("/link")

	routeGroup.Use(middleware.AuthMiddleware(*DBController.Config))
	routeGroup.POST("/create_link", DBController.CreateLink)
	routeGroup.GET("/get_links", DBController.GetLinks)
	routeGroup.DELETE("/delete_link/:link_id", DBController.DeleteLink)
	routeGroup.PUT("/update_link/:link_id", DBController.UpdateLink)
}
