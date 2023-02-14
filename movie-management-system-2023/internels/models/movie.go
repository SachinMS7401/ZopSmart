package models

import (
	"database/sql"
	"time"
)

type Movie struct {
	ID           int            `json:"id,omitempty"`
	Name         string         `json:"name"`
	Genre        string         `json:"genre"`
	Rating       float64        `json:"rating"`
	ReleasedDate string         `json:"releasedDate"`
	UpdatedAt    time.Time      `json:"updatedAt,omitempty"`
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	Plot         string         `json:"plot"`
	Released     bool           `json:"released"`
	DeletedAt    sql.NullString `json:"-"`
}

type UpdateMovie struct {
	Rating       float64 `json:"rating,omitempty"`
	Plot         string  `json:"plot,omitempty"`
	ReleasedDate string  `json:"releasedDate,omitempty"`
}
