package domain

type LocationResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	UserID    int     `json:"user_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
