package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/database/repository"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/middleware"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/models"
	"github.com/gin-gonic/gin"
)

func AccountRouters(db repository.DbInstance, config model.Config, router *gin.Engine) {

	Controller := controllers.NewController(&db, config)

	routeGroup := router.Group("/account")
	routeGroup.GET("/search",Controller.UsernameSearch)
	routeGroup.GET("/:username",Controller.GetAccountByUsername)

	routeGroup.Use(middleware.AuthMiddleware(config))
	routeGroup.PUT("/edit_account", Controller.EditAccount)
	routeGroup.DELETE("/delete_account", Controller.DeleteAccount)
	routeGroup.GET("/get_account", Controller.GetAccount)
	routeGroup.POST("/add_photo", Controller.UploadUserImage)
	

}
