package repository

import (
	"time"
)

type FontankaNews struct {
	ID              uint64    `db:"id"`
	Title           string    `db:"title"`
	PublicationDate string    `db:"publication_date"`
	Link            string    `db:"link"`
	CreatedAt       time.Time `db:"created_at"`
}

func (f *FontankaNews) InsertColumns() []string {
	return []string{
		"title",
		"publication_date",
		"link",
		"created_at",
	}
}

func (f *FontankaNews) Columns() []string {
	return []string{
		"id",
		"title",
		"publication_date",
		"link",
		"created_at",
	}
}

func (f *FontankaNews) Values() []interface{} {
	return []interface{}{
		f.ID,
		f.Title,
		f.PublicationDate,
		f.Link,
		f.CreatedAt,
	}
}

func (f *FontankaNews) ScanValues() []interface{} {
	return []interface{}{
		&f.ID,
		&f.Title,
		&f.PublicationDate,
		&f.Link,
		&f.CreatedAt,
	}
}
