package main

import (
	"log"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/middleware"
	_ "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	route "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/routers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
	_ "github.com/codescalersinternships/Link-Tree-Dohaelsawy/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			LinkTree
//	@version		1.0
//	@description	platform that allows users to share social media profiles, or other important links defined in a single place

//	@host	localhost:8010
//	@BasePath

//	@securityDefinitions.basic	BasicAuth

func main() {
	db, err := repository.DbConnect()
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}

	router := gin.Default()
	router.Use(middleware.CorsMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	config, err := utils.NewConfigController()
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}

	dbInstance := repository.NewDbInstance(db)

	route.AccountRouters(dbInstance, config, router)
	route.LinkRouters(dbInstance, config, router)
	route.AuthRouters(dbInstance, config, router)
	route.AnalyticsRouters(dbInstance, config, router)

	router.Run()
}
