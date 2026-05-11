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

func (s *SessionRepository) GetUserID(ctx context.Context, sessionToken string) (int, error) {
	val, err := s.client.Get(ctx, sessionToken).Result()
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
	key := fmt.Sprintf("session:%v", session.ID)
	return s.client.Set(ctx, key, session.UserID, session.Duration).Err()
}

func (s *SessionRepository) DeleteSession(ctx context.Context, sessionID string) error {
	key := fmt.Sprintf("session:%v", sessionID)
	return s.client.Del(ctx, key).Err()
}
