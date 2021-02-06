package domain

import "time"

type HabrNews struct {
	ID              string    `json:"id"`
	Author          string    `json:"author"`
	AuthorLink      string    `json:"author_link"`
	Title           string    `json:"title"`
	Preview         string    `json:"preview"`
	Views           string    `json:"views"`
	PublicationDate string    `json:"publication_date"`
	Link            string    `json:"link"`
	CreatedAt       time.Time `json:"created_at"`
}
