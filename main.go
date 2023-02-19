package main

import (
	"database/sql"
	"log"

	// "os"

	"github.com/AlbertoParente/go-finance-app/api"
	db "github.com/AlbertoParente/go-finance-app/db/sqlc"
	_ "github.com/lib/pq"
	// "github.com/joho/godotenv"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// 	dbDriver := os.Getenv("DB_DRIVER")
	// 	dbSource := os.Getenv("DB_SOURCE")
	// 	serverAddress := os.Getenv("SERVER_ADDRESS")

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start api: ", err)
	}
}
