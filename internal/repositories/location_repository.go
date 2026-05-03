package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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
			return domain.Location{}, domain.ErrLocationNotFound
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

	return location, nil
}

func (r *LocationRepository) GetLocations(userID int) ([]domain.Location, error) {
	queryForLocations := `SELECT id, name, user_id, latitude, longitude FROM locations WHERE user_id = $1`

	var result []domain.Location

	rows, err := r.db.Query(queryForLocations, userID)
	if err != nil {
		return result, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("rows close error: %v", err)
		}
	}()

	for rows.Next() {
		var location domain.Location
		err := rows.Scan(
			&location.ID,
			&location.Name,
			&location.UserID,
			&location.Latitude,
			&location.Longitude,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, location)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	fmt.Print(result)
	return result, nil
}

func (r *LocationRepository) RemoveLocation(id int) error {
	query := `DELETE FROM locations WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
