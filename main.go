package main

import (
	"log"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/routers"
)

func main() {
	db, err := repository.DbConnect()
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}

	err = db.Migrate()
	if err != nil {
		log.Printf("Error %s\n", err)
		return
	}

	routes := routers.AuthRouter(db)

	routes.Run()
}
