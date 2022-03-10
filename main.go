package main

import (
	"database/sql"
	"log"

	"github.com/honghuangdc/soybean-admin-go/api"
	db "github.com/honghuangdc/soybean-admin-go/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/soybean_admin?sslmode=disable"
	serverAddress = "0.0.0.0:9999"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("connot start server:", err)
	}
}
