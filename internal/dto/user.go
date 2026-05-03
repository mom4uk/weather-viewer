package dto

import (
	"regexp"
	"unicode/utf8"
	"weather-viewer/internal/domain"
)

type UserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserResponse struct {
	Login string `json:"login"`
}

func ValidateCredentials(login string, password string) error {
	if login == "" || password == "" {
		return domain.ErrAbsenceOfLoginPass
	}

	var loginRule = regexp.MustCompile(`^[A-Za-z0-9]+$`)
	if utf8.RuneCountInString(login) < 6 || utf8.RuneCountInString(login) > 20 {
		return domain.ErrLoginInvalidLength
	}
	if !loginRule.MatchString(login) {
		return domain.ErrInvalidLogin
	}

	if utf8.RuneCountInString(password) < 6 || utf8.RuneCountInString(password) > 20 {
		return domain.ErrPasswordInvalidLength
	}

	return nil
}
