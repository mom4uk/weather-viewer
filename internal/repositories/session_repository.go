package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"weather-viewer/internal/domain"
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (s *SessionRepository) GetSession(id string) (domain.Session, error) {
	query := `SELECT id, user_id, expires_at FROM sessions WHERE id = $1`

	var session domain.Session

	err := s.db.QueryRow(query, id).Scan(&session.ID, &session.UserID, &session.ExpiresAt)
	if err != nil {
		fmt.Print(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Session{}, domain.ErrSessionNotFound
		}
		return domain.Session{}, err
	}

	return session, nil
}

func (s *SessionRepository) CreateSession(session domain.Session) error {
	query := `INSERT INTO sessions (id, user_id, expires_at) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(
		query,
		session.ID,
		session.UserID,
		session.ExpiresAt,
	)
	return err
}

func (s *SessionRepository) DeleteSession(sessionID string) error {
	query := `DELETE FROM sessions WHERE id = $1`
	_, err := s.db.Exec(query, sessionID)
	if err != nil {
		return err
	}
	return nil
}
