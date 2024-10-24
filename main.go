package main

import (
	"log"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	route "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/routers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repository.DbConnect()
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}

	router := gin.Default()
	config, err := utils.NewConfigController()
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}
	dbInstance := repository.NewDbInstance(db)

	route.AccountRouters(dbInstance, config, router)
	route.LinkRouters(dbInstance, config, router)
	route.AuthRouters(dbInstance, config, router)

	router.Run()
}
