package main

import (
	"log"
	"net/http"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repository.DbConnect()

	if err != nil {
		log.Printf("error: %s\n", err)
	}

	db.Migrate()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	if err != nil {
		log.Printf("Error loading .env file, %s\n", err)
	}
	
	router.Run()
}
