package models

import "time"

type HabrNews struct {
	ID              uint64    `json:"id"`
	Author          uint64    `json:"author"`
	Title           string    `json:"title"`
	Preview         string    `json:"preview"`
	Tags            string    `json:"tags"`
	PublicationDate time.Time `json:"publication_date"`
	Link            string    `json:"link"`
}

type FontankaNews struct {
	ID              uint64 `json:"id"`
	Title           string `json:"title"`
	PublicationDate string `json:"publication_date"`
	Link            string `json:"link"`
}
