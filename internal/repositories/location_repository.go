package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"weather-viewer/internal/domain"
)

type LocationRepository struct {
	db *sql.DB
}

func NewLocationRepository(db *sql.DB) *LocationRepository {
	return &LocationRepository{
		db: db,
	}
}
func (r *LocationRepository) GetLocation(id int) (domain.Location, error) {
	query := `SELECT id, name, user_id, latitude, longitude FROM locations WHERE id = $1`

	var location domain.Location
	err := r.db.QueryRow(query, id).Scan(
		&location.ID,
		&location.Name,
		&location.UserID,
		&location.Latitude,
		&location.Longitude,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Location{}, domain.ErrIncorrectNotFound
		}
		return domain.Location{}, err
	}

	return location, nil
}

func (r *LocationRepository) AddLocation(location domain.Location) (domain.Location, error) {
	query := `INSERT INTO locations (name, user_id, latitude, longitude) VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.db.QueryRow(
		query,
		location.Name,
		location.UserID,
		location.Latitude,
		location.Longitude,
	).Scan(&location.ID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return domain.Location{}, domain.ErrLocationAlreadyExists
		}
		return domain.Location{}, err
	}
	fmt.Print("test")
	return location, nil
}
