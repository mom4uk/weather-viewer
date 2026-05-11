package services

import (
	"context"
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

func (s *SessionService) GetUserID(ctx context.Context, sessionToken string) (int, error) {
	return s.repo.GetUserID(ctx, sessionToken)
}

func (s *SessionService) CreateSession(ctx context.Context, userId int) (domain.Session, error) {
	session := domain.Session{
		ID:       uuid.New().String(),
		UserID:   userId,
		Duration: time.Hour,
	}
	err := s.repo.CreateSession(ctx, session)
	if err != nil {
		return domain.Session{}, err
	}

	return session, nil
}

func (s *SessionService) DeleteSession(ctx context.Context, sessionID string) error {
	return s.repo.DeleteSession(ctx, sessionID)
}
