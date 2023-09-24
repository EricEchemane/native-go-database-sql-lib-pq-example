package db

import (
	"database/sql"
	"go_pg_anap/config"
	"log"

	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	config := config.Init()

	db, err := sql.Open("postgres", config.PGconnStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
