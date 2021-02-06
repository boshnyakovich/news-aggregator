package repository

import (
	"time"
)

type HTNews struct {
	ID              string    `db:"id"`
	Category        string    `db:"category"`
	Title           string    `db:"title"`
	Preview         string    `db:"preview"`
	Link            string    `db:"link"`
	CreatedAt       time.Time `db:"created_at"`
}

func (ht *HTNews) InsertColumns() []string {
	return []string{
		"id",
		"category",
		"title",
		"preview",
		"link",
		"created_at",
	}
}

func (ht *HTNews) Columns() []string {
	return []string{
		"id",
		"category",
		"title",
		"preview",
		"link",
		"created_at",
	}
}

func (ht *HTNews) Values() []interface{} {
	return []interface{}{
		ht.ID,
		ht.Category,
		ht.Title,
		ht.Preview,
		ht.Link,
		ht.CreatedAt,
	}
}

func (ht *HTNews) ScanValues() []interface{} {
	return []interface{}{
		&ht.ID,
		&ht.Category,
		&ht.Title,
		&ht.Preview,
		&ht.Link,
		&ht.CreatedAt,
	}
}
