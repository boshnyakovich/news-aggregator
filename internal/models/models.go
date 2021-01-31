package models

import "time"

type HabrNews struct {
	ID              uint64    `json:"stats"`
	Author          uint64    `json:"author"`
	Title           string    `json:"title"`
	Preview         string    `json:"preview"`
	Tags            string    `json:"tags"`
	PublicationDate time.Time `json:"publicationDate"`
	Link            string    `json:"link"`
}

type FourPDANews struct {
	ID              uint64    `json:"stats"`
	Title           string    `json:"title"`
	Preview         string    `json:"preview"`
	PublicationDate time.Time `json:"publicationDate"`
	Link            string    `json:"link"`
}
