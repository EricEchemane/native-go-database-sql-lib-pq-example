package main

import (
	"fmt"
	"go_pg_anap/db"
	"go_pg_anap/db/migrations"
	"log"
)

func main() {
	db := db.Init()
	defer db.Close()

	migrations.Migrate(db)

	createProductQuery := `INSERT INTO product(name,code,price) VALUES($1,$2,$3) RETURNING id,created_at`
	var pk int
	var created_at string
	err := db.QueryRow(createProductQuery, "magi magic sarap", "msg-mms3", 5).Scan(&pk, &created_at)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pk, created_at)
}
