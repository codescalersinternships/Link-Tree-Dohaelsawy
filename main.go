package main

import (
	"log"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	route "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repository.DbConnect()
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}

	router := gin.Default()
	
	route.AccountRouters(db, router)
	route.LinkRouters(db, router)
	route.AuthRouters(db, router)
	

	router.Run()
}
