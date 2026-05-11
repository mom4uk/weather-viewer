package repositories

import (
	"context"
	"fmt"
	"strconv"
	"weather-viewer/internal/domain"

	"github.com/redis/go-redis/v9"
)

type SessionRepository struct {
	client *redis.Client
}

func NewSessionRepository(client *redis.Client) *SessionRepository {
	return &SessionRepository{
		client: client,
	}
}

func (s *SessionRepository) GetUserID(ctx context.Context, sessionID string) (int, error) {
	val, err := s.client.Get(ctx, sessionKey(sessionID)).Result()
	if err != nil {
		return 0, err
	}

	userID, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (s *SessionRepository) CreateSession(ctx context.Context, session domain.Session) error {
	return s.client.Set(ctx, sessionKey(session.ID), session.UserID, session.Duration).Err()
}

func (s *SessionRepository) DeleteSession(ctx context.Context, sessionID string) error {
	return s.client.Del(ctx, sessionKey(sessionID)).Err()
}

func sessionKey(sessionID string) string {
	return fmt.Sprintf("session:%v", sessionID)
}
