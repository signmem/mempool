package db

import (
	"database/sql"
	"github.com/signmem/mempool/g"
	"log"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err := sql.Open("mysql", g.Config().Database)
	if err != nil {
		log.Printf("[ERROR] db open error: %s\n", err)
	}

	DB.SetMaxIdleConns(g.Config().MaxIdle)

	err = DB.Ping()
	if err != nil {
		log.Printf("[ERROR] db ping error %s\n", err)
	}
}