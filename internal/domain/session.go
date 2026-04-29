package domain

import "time"

type Session struct {
	ID        int
	UserID    int
	ExpiresAt time.Time
}
