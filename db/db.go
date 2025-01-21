package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDBWithConnection(conn *sql.DB) {
	DB = conn
}

func InitDB(addr string) {
	var err error
	DB, err = sql.Open("postgres", addr)
	log.Println(addr, err)
	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
