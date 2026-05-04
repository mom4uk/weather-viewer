package services

import (
	"time"
	"weather-viewer/internal/domain"
	"weather-viewer/internal/repositories"

	"github.com/google/uuid"
)

type SessionService struct {
	repo *repositories.SessionRepository
}

func NewSessionService(repo *repositories.SessionRepository) *SessionService {
	return &SessionService{
		repo: repo,
	}
}

func (s *SessionService) GetSession(id string) (domain.Session, error) {
	return s.repo.GetSession(id)
}

func (s *SessionService) CreateSession(userId int) (domain.Session, error) {
	session := domain.Session{
		ID:        uuid.New().String(),
		UserID:    userId,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}
	err := s.repo.CreateSession(session)
	if err != nil {
		return domain.Session{}, err
	}

	return session, nil
}
