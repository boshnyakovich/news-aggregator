package dao

import "time"

type HabrNews struct {
	ID              uint64    `db:"id"`
	Author          string    `db:"author"`
	AuthorLink      string    `db:"author_link"`
	Title           string    `db:"title"`
	Preview         string    `db:"preview"`
	Views           string    `db:"views"`
	PublicationDate string    `db:"publication_date"`
	Link            string    `db:"link"`
	CreatedAt       time.Time `db:"created_at"`
}

func (h *HabrNews) InsertColumns() []string {
	return []string{
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
