package domain

import "errors"

type ErrorResponse struct {
	Message string `json:"message"`
}

var (
	ErrInvalidId             = errors.New("incorrect id for location")
	ErrLocationNotFound      = errors.New("location not found")
	ErrInvalidName           = errors.New("invalid name for location")
	ErrInvalidLatitude       = errors.New("invalid latitude for location")
	ErrInvalidLongitude      = errors.New("invalid longitude for location")
	ErrLocationAlreadyExists = errors.New("location already exists")
	ErrSessionNotFound       = errors.New("session not found")
	ErrUserNotFound          = errors.New("user not found")
)
