package domain

import "errors"

type ErrorResponse struct {
	Message string `json:"message"`
}

var (
	ErrIncorrectId = errors.New("incorrect id for location")
)
