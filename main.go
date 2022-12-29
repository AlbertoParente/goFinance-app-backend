package main

import (
	// "database/sql"
	// "log"

	_ "github.com/lib/pq"
	// db "github.com/AlbertoParente/go-finance-app/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	// conn, err := sql.Open(dbDriver, dbSource)
	// if err != nil {
	// 	log.Fatal("Cannot connect to db: ", err)
	// }

	// store := db.NewStore(conn)
	// server := api.newServer(store)

	// err = server.Start(serverAddress)
	// if err != nil {
	// 	log.Fatal("Cannot start api: ", err)
	// }
}
