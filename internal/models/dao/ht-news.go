package dao

import (
	"time"
)

type HTNews struct {
	ID              string    `db:"id"`
	Title           string    `db:"title"`
	PublicationDate string    `db:"publication_date"`
	Link            string    `db:"link"`
	CreatedAt       time.Time `db:"created_at"`
}

func (ht *HTNews) InsertColumns() []string {
	return []string{
		"id",
		"title",
		"publication_date",
		"link",
		"created_at",
	}
}

func (ht *HTNews) Columns() []string {
	return []string{
		"id",
		"title",
		"publication_date",
		"link",
		"created_at",
	}
}

func (ht *HTNews) Values() []interface{} {
	return []interface{}{
		ht.ID,
		ht.Title,
		ht.PublicationDate,
		ht.Link,
		ht.CreatedAt,
	}
}

func (ht *HTNews) ScanValues() []interface{} {
	return []interface{}{
		&ht.ID,
		&ht.Title,
		&ht.PublicationDate,
		&ht.Link,
		&ht.CreatedAt,
	}
}
