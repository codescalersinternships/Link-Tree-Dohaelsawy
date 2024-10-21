package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/middleware"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func LinkRouters(db repository.DbInstance, router *gin.Engine) {

	validate := validator.New()
	linkController := controllers.NewLinksController(db, validate)

	routeGroup := router.Group("/link")

	routeGroup.Use(middleware.AuthMiddleware())
	routeGroup.POST("/create_link", linkController.CreateLink)
	routeGroup.GET("/get_links", linkController.GetLinks)
	routeGroup.DELETE("/delete_link/:link_id", linkController.DeleteLink)
	routeGroup.PUT("/update_link/:link_id", linkController.UpdateLink)
}
