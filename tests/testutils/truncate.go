package testutils

import "database/sql"

func TruncateAll(db *sql.DB) error {
	_, err := db.Exec(`
		TRUNCATE TABLE locations, users RESTART IDENTITY CASCADE
	`)
	return err
}
