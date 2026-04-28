package domain

import "errors"

type ErrorResponse struct {
	Message string `json:"message"`
}

var (
	ErrInvalidId         = errors.New("incorrect id for location")
	ErrIncorrectNotFound = errors.New("location not found")
)
