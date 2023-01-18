package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connerct to db: ", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}

func TestMain2(m *testing.M) {
	conn1, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connerct to db: ", err)
	}
	testQueries = New(conn1)
	os.Exit(m.Run())
}

func TestMain3(m *testing.M) {
	conn2, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connerct to db: ", err)
	}
	testQueries = New(conn2)
	os.Exit(m.Run())
}
