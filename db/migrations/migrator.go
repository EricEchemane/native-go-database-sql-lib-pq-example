package migrations

import (
	"database/sql"
)

func Migrate(db *sql.DB) {
	createProductTable(db)
}
