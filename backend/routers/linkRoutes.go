package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/middleware"

	"github.com/gin-gonic/gin"
)

func LinkRouters(db repository.DbInstance, router *gin.Engine) {

	DBController := controllers.NewDBService(&db)

	routeGroup := router.Group("/link")

	routeGroup.Use(middleware.AuthMiddleware())
	routeGroup.POST("/create_link", DBController.CreateLink)
	routeGroup.GET("/get_links", DBController.GetLinks)
	routeGroup.DELETE("/delete_link/:link_id", DBController.DeleteLink)
	routeGroup.PUT("/update_link/:link_id", DBController.UpdateLink)
}
