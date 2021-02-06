package models

type Period string

const (
	DAY   = Period("day")
	WEEK  = Period("week")
	MONTH = Period("month")
	YEAR  = Period("year")
)
