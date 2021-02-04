package domain

import "time"

type HTNews struct {
	Category        string    `json:"category"`
	Title           string    `json:"title"`
	Preview         string    `json:"preview"`
	Link            string    `json:"link"`
	CreatedAt       time.Time `json:"created_at"`
}
