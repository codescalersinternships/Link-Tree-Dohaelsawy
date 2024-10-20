package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func DbConnection() (*sql.DB, error) {
	driver, connStr := prepareDbConnectionString()
	db, err := sql.Open(driver, connStr)

	if err != nil {
		log.Printf("DB open failed, %s", err)
		return nil, err
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Printf("DB Ping Failed, %s", err)
		return nil, err
	}

	log.Println("DB Connection started successfully")
	return db, nil
}

func prepareDbConnectionString() (string, string) {
	host := os.Getenv("DB_HOST")
	driver := os.Getenv("DB_DRIVER")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	return driver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
}
