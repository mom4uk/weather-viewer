package testutils

import "database/sql"

func SeedUsers(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO users (login, password)
		VALUES ('test', 'test')
	`)
	return err
}

func SeedLocations(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO locations (name, user_id, latitude, longitude)
		VALUES
		('Москва', 1, 0, 0),
		('Санкт-Петербург', 1, 1, 1)
	`)
	return err
}
