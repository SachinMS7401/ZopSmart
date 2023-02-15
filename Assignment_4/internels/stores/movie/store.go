package movie

import (
	"Assignment_4/internels/filters"
	"Assignment_4/internels/stores"
	"Assignment_4/models"
	"context"
	"database/sql"
	"fmt"
)

const tableName = "Movies_details"

type store struct {
	store stores.MovieStorer
}

func New(db *sql.DB) *store {
	return &store{db: db}
}

func (s *store) Get(ctx context.Context, f filters.Movie) ([]models.Movie, error) {
	var details []interface{}

	if f.Id > 0 {
		details = append(details, f.Id)
	}
	query := "select * from " + tableName + " where Id = ?"

	rows, err := s.db.QueryContext(ctx, query, details...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var m models.Movie

		err := rows.Scan(&m.Id, &m.Name, &m.Genre, &m.Rating, &m.Released)
		if err != nil {
			return nil, err
		}

		movies = append(movies, m)
	}

	fmt.Println("Data in movies is ", movies)
	return movies, nil
}

func (s *store) Post(ctx context.Context, m models.Movie) (*models.Movie, error) {

	var details []interface{}

	query := "INSERT INTO Movie_details VALUES(?,?,?,?,?,?,?,?,?)"
	details = append(details, m.Id, m.Name, m.Genre, m.Rating, m.Released)

	row, err := s.db.ExecContext(ctx, query, details...)

	if err != nil {
		fmt.Println("Error while inserting into database")
		return nil, err
	}

	n, _ := row.RowsAffected()

	if n > 0 {
		fmt.Println(n, " number of rows affected in the database")
	}
	var mov models.Movie

	mov.Id = m.Id
	mov.Name = m.Name
	mov.Genre = m.Genre
	mov.Rating = m.Rating

	fmt.Println(mov)

	return &mov, nil

}
