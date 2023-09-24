package migrations

import (
	"database/sql"
	"log"
	"strings"
)

func createProductTable(db *sql.DB) {
	createProductQuery := `CREATE TABLE IF NOT EXISTS product (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		code VARCHAR(100) NOT NULL,
		price NUMERIC(6,2) NOT NULL,
		created_at timestamp DEFAULT NOW(),
		modified_at timestamp DEFAULT NULL,
		CONSTRAINT product_code_unique UNIQUE (code)
	)`

	createProductNameIndexQuery := `CREATE INDEX product_code_idx ON product(code)`

	_, err := db.Exec(createProductQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(createProductNameIndexQuery)
	if err != nil && !strings.Contains(err.Error(), "relation \"product_code_idx\" already exists") {
		log.Fatal(err)
	}
}
