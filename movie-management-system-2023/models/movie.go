package models

import (
	"database/sql"
	"time"
)

type Movie struct {
	Id           int            `json:"id"`
	Name         string         `json:"name"`
	Genre        string         `json:"genre"`
	Rating       float64        `json:"rating"`
	ReleasedDate string         `json:"releaseDate"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	CreatedAt    time.Time      `json:"createdAt"`
	Plot         string         `json:"plot"`
	Released     bool           `json:"released"`
	DeletedAt    sql.NullString `json:"-"`
}
