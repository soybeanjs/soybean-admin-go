package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/honghuangdc/soybean-admin-go/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadCconfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
