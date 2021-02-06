package models

import "time"

type HabrCriteria struct {
	Articles bool   `json:"articles"`
	All      bool   `json:"all"`
	Rating   uint64 `json:"rating,omitempty"`
	Best     bool   `json:"best"`
	Period   string `json:"period,omitempty"`
}

type HabrNews struct {
	ID              string    `json:"id" db:"id"`
	Author          string    `json:"author" db:"author"`
	AuthorLink      string    `json:"author_link" db:"author_link"`
	Title           string    `json:"title" db:"title"`
	Preview         string    `json:"preview,omitempty" db:"preview"`
	Views           string    `json:"views" db:"views"`
	PublicationDate string    `json:"publication_date" db:"publication_date"`
	Link            string    `json:"link" db:"link"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

func (h *HabrNews) InsertColumns() []string {
	return []string{
		"id",
		"author",
		"author_link",
		"title",
		"preview",
		"views",
		"publication_date",
		"link",
		"created_at",
	}
}

func (h *HabrNews) Columns() []string {
	return []string{
		"id",
		"author",
		"author_link",
		"title",
		"preview",
		"views",
		"publication_date",
		"link",
		"created_at",
	}
}

func (h *HabrNews) Values() []interface{} {
	return []interface{}{
		h.ID,
		h.Author,
		h.AuthorLink,
		h.Title,
		h.Preview,
		h.Views,
		h.PublicationDate,
		h.Link,
		h.CreatedAt,
	}
}

func (h *HabrNews) ScanValues() []interface{} {
	return []interface{}{
		&h.ID,
		&h.Author,
		&h.AuthorLink,
		&h.Title,
		&h.Preview,
		&h.Views,
		&h.PublicationDate,
		&h.Link,
		&h.CreatedAt,
	}
}
