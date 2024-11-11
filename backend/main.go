package main

import (
	"fmt"
	"log"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/database/repository"
	_ "github.com/codescalersinternships/Link-Tree-Dohaelsawy/docs"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/middleware"
	_ "github.com/codescalersinternships/Link-Tree-Dohaelsawy/models"
	route "github.com/codescalersinternships/Link-Tree-Dohaelsawy/routers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			LinkTree
//	@version		1.0
//	@description	platform that allows users to share social media profiles, or other important links defined in a single place

//	@host	185.206.122.17:31010
//	@BasePath

//	@securityDefinitions.basic	BasicAuth

func main() {
	db, err := repository.DbConnect()
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}
	config, err := utils.NewConfigController()
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}
	fmt.Println(config.BaseUrl)

	router := gin.Default()
	router.Use(middleware.CorsMiddleware(config.Origin))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	

	dbInstance := repository.NewDbInstance(db)

	route.AccountRouters(dbInstance, config, router)
	route.LinkRouters(dbInstance, config, router)
	route.AuthRouters(dbInstance, config, router)
	route.AnalyticsRouters(dbInstance, config, router)

	err = router.Run()
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}
}
