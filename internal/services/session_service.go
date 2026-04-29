package services

import (
	"weather-viewer/internal/domain"
	"weather-viewer/internal/repositories"
)

type SessionService struct {
	repo *repositories.SessionRepository
}

func NewSessionService(repo *repositories.SessionRepository) *SessionService {
	return &SessionService{
		repo: repo,
	}
}

func (s *SessionService) Authenticate(id string) (domain.Session, error) {
	return s.repo.GetSession(id)
}
