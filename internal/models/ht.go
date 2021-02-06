package models

import "time"

type HTCriteria struct {
	Category string `json:"category,omitempty"`
	Page     uint64 `json:"page"`
}

type HTNews struct {
	ID        string    `json:"id" db:"id"`
	Category  string    `json:"category" db:"category"`
	Title     string    `json:"title" db:"title"`
	Preview   string    `json:"preview" db:"preview"`
	Link      string    `json:"link" db:"link"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
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
