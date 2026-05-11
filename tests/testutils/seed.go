package testutils

import (
	"database/sql"
	"weather-viewer/internal/utilities"
)

func SeedUsers(db *sql.DB) error {
	hash, err := utilities.HashPassword("qwerty1234")
	if err != nil {
		return err
	}
	_, err = db.Exec(`
		INSERT INTO users (login, password)
		VALUES ('test1234', $1)
	`, hash)
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

func SeedSession(db *sql.DB, sessionID string) error {
	_, err := db.Exec(`
		INSERT INTO sessions (id, user_id, expires_at)
		VALUES ($1, $2, NOW() + interval '1 hour')
	`, sessionID, 1)

	return err
}
