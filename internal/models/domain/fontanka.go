package domain

import "time"

type FontankaNews struct {
	Title           string    `json:"title"`
	PublicationDate string    `json:"publication_date"`
	Link            string    `json:"link"`
	CreatedAt       time.Time `json:"created_at"`
}
