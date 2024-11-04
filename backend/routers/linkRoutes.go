package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/middleware"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"

	"github.com/gin-gonic/gin"
)

func LinkRouters(db repository.DbInstance, config model.Config, router *gin.Engine) {

	Controller := controllers.NewController(&db, config)

	router.GET("/link_tree/:username", Controller.GetLinks)

	routeGroup := router.Group("/link")

	routeGroup.Use(middleware.AuthMiddleware(config))
	routeGroup.POST("/create_link", Controller.CreateLink)
	routeGroup.DELETE("/delete_link/:link_id", Controller.DeleteLink)
	routeGroup.PUT("/update_link/:link_id", Controller.UpdateLink)
}
