package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AccountRouters(db repository.DbInstance, router *gin.Engine) {

	validate := validator.New(validator.WithRequiredStructEnabled())
	DBController := controllers.NewDBControllerImpl(db, *validate)

	routeGroup := router.Group("/account")

	routeGroup.Use(middleware.AuthMiddleware())
	routeGroup.POST("/edit_account/:user_id", DBController.EditAccount)
	routeGroup.DELETE("/delete_account/:user_id", DBController.DeleteAccount)
	routeGroup.GET("/get_account/",DBController.GetAccount)
	routeGroup.GET("/create_link_tree_url",DBController.CreateLinkTreeUrl)

}
