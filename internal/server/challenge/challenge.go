package challenge

import "time"

type Challenge struct {
	ID         string    `json:"id"`
	Difficulty int       `json:"difficulty"`
	ExpiresAt  time.Time `json:"expires_at"`
}
