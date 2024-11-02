package main

import (
	"context"
	"log"

	awsConfig"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	_ "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	route "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/routers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
	_ "github.com/codescalersinternships/Link-Tree-Dohaelsawy/docs"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
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

	cfg, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	client := s3.NewFromConfig(cfg)

	router := gin.Default()
	router.Use(cors.Default())
	// router.Static("/images", "./database/users_photo/")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	config, err := utils.NewConfigController()
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}

	dbInstance := repository.NewDbInstance(db)

	route.AccountRouters(dbInstance, config, router, client)
	route.LinkRouters(dbInstance, config, router,client)
	route.AuthRouters(dbInstance, config, router,client)
	route.AnalyticsRouters(dbInstance, config, router,client)

	router.Run()
}
