package filters

import "database/sql"

type Movie struct {
	Id int
}

type Details struct {
	Name         string  `json:"name"`
	Genre        string  `json:"genre"`
	Rating       float64 `json:"rating"`
	ReleasedDate string  `json:"releasedDate"`
	Plot         string  `json:"plot"`
	Released     bool    `json:"released"`
}

type UpdateMovie struct {
	Rating       float64 `json:"rating,omitempty"`
	Plot         string  `json:"plot,omitempty"`
	ReleasedDate string  `json:"releasedDate,omitempty"`
}

type Deleted struct {
	DeletedAt sql.NullString `json:"-"`
}
