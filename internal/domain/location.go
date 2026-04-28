package domain

type LocationResponse struct {
	Name      string  `json:"name"`
	UserID    int64   `json:"user_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
