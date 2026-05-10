package domain

import "time"

type Session struct {
	ID        string
	UserID    int
	ExpiresAt time.Time
}

func (s Session) IsExpired(now time.Time) bool {
	return s.ExpiresAt.Before(now)
}
