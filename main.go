package main

import (
	"database/sql"
	"log"

	"github.com/honghuangdc/soybean-admin-go/api"
	db "github.com/honghuangdc/soybean-admin-go/db/sqlc"
	"github.com/honghuangdc/soybean-admin-go/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadCconfig(".")
	if err != nil {
		log.Fatal("connect load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("connot start server:", err)
	}
}
