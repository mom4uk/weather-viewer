package domain

import "errors"

type ErrorResponse struct {
	Message string `json:"message"`
}

type WeatherAPIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e WeatherAPIError) Error() string {
	return e.Message
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
	ErrIncorrectCredentials  = errors.New("incorrect login or password")
	ErrLoginInvalidLength    = errors.New("login length is invalid")
	ErrPasswordInvalidLength = errors.New("password length is invalid")
	ErrInvalidLogin          = errors.New("invalid login")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrAbsenceOfLoginPass    = errors.New("absence of login / password")
)
