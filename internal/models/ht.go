package models

import "time"

type HTNews struct {
	ID        string    `json:"id"`
	Category  string    `json:"category"`
	Title     string    `json:"title"`
	Preview   string    `json:"preview"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
}
