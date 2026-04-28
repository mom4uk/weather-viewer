package httputil

import (
	"strconv"
	"strings"
	"weather-viewer/internal/domain"
)

func GetIdFromUrl(url string) (int, error) {
	index := strings.LastIndex(url, "/")
	id, err := strconv.Atoi(url[index+1:])
	if err != nil {
		return 0, domain.ErrIncorrectId
	}
	return id, nil
}
