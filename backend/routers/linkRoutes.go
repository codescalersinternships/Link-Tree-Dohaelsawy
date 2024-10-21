package routers

import (
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/middleware"
	"github.com/gin-gonic/gin"
)

func LinkRouters(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())
	router.POST("/create_link", controllers.CreateLink())
	router.GET("/get_links", controllers.GetLinks())
	router.DELETE("/link/:link_id", controllers.DeleteLink())
	router.PUT("/update_link", controllers.UpdateLink())
}
