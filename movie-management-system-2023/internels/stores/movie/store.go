package movie

import (
	"strings"

	gofrErrors "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"github.com/go-training/movie-management-system-2023/internels/models"
)

const tableName = "sample5"

type Store struct {
}

func New() *Store {
	return &Store{}
}
func (s *Store) GetAll(ctx *gofr.Context) ([]models.Movie, error) {
	query := "select * from " + tableName + " where deletedAt is NULL;"
	rows, err := ctx.DB().QueryContext(ctx, query)

	if err != nil {
		return nil, gofrErrors.DB{Err: err}
	}

	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Rating, &movie.ReleasedDate, &movie.UpdatedAt,
			&movie.CreatedAt, &movie.Plot, &movie.Released, &movie.DeletedAt)

		if err != nil {
			return nil, gofrErrors.DB{Err: err}
		}

		movie.ReleasedDate = movie.ReleasedDate[:10]
		movies = append(movies, movie)
	}

	return movies, nil
}

func (s *Store) Get(ctx *gofr.Context, movID int) (models.Movie, bool, error) {
	var details []interface{}

	if movID > 0 {
		details = append(details, movID)
	}

	query := "select id,name,genre,rating,releasedDate,updatedAt,createdAt,plot,released," +
		"deletedAt from " + tableName + " where (deletedAt is NULL AND id = ?);"
	rows, err := ctx.DB().QueryContext(ctx, query, details...)

	if err != nil {
		return models.Movie{}, true, gofrErrors.DB{Err: err}
	}

	defer rows.Close()

	deleted := true

	var movie models.Movie

	for rows.Next() {
		err = rows.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Rating, &movie.ReleasedDate,
			&movie.UpdatedAt, &movie.CreatedAt, &movie.Plot, &movie.Released, &movie.DeletedAt)

		if err != nil {
			return models.Movie{}, true, gofrErrors.DB{Err: err}
		}

		movie.ReleasedDate = movie.ReleasedDate[:10]
		deleted = false
	}

	if deleted {
		return models.Movie{}, true, gofrErrors.DB{Err: err}
	}

	return movie, deleted, nil
}

func (s *Store) Post(ctx *gofr.Context, mov *models.Movie) (models.Movie, bool, error) {
	var details []interface{}

	query := "INSERT INTO " + tableName + "(name,genre,rating,releasedDate,updatedAt,createdAt,plot,released) " +
		"VALUES(?,?,?,?,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP,?,?)"

	details = append(details, mov.Name, mov.Genre, mov.Rating, mov.ReleasedDate, mov.Plot, mov.Released)

	row, err := ctx.DB().ExecContext(ctx, query, details...)

	if err != nil {
		return models.Movie{}, true, gofrErrors.DB{Err: err}
	}

	id, err := row.LastInsertId()

	if err != nil {
		return models.Movie{}, true, gofrErrors.DB{Err: err}
	}

	movieID := int(id)
	movie, deleted, err := s.Get(ctx, movieID)

	if err != nil {
		return models.Movie{}, true, gofrErrors.DB{Err: err}
	}

	return movie, deleted, nil
}

func (s *Store) Update(ctx *gofr.Context, movieID int, movDetails models.UpdateMovie) (models.Movie, bool, error) {
	var patch []string

	var details []interface{}

	if movDetails.Rating > 0 {
		patch = append(patch, "rating=?")
		details = append(details, movDetails.Rating)
	}

	if movDetails.Plot != "" {
		patch = append(patch, "plot=?")
		details = append(details, movDetails.Plot)
	}

	if movDetails.ReleasedDate != "" {
		patch = append(patch, "releasedDate=?")
		details = append(details, movDetails.ReleasedDate)
	}

	query := "UPDATE " + tableName + " SET updatedAt=CURRENT_TIMESTAMP,"

	if len(patch) > 0 {
		query += strings.Join(patch, ",") + " WHERE id=?;"
	}

	details = append(details, movieID)
	_, err := ctx.DB().ExecContext(ctx, query, details...)

	if err != nil {
		return models.Movie{}, true, gofrErrors.DB{Err: err}
	}

	movie, deleted, err := s.Get(ctx, movieID)

	if err != nil {
		return models.Movie{}, true, gofrErrors.DB{Err: err}
	}

	return movie, deleted, nil
}

func (s *Store) Delete(ctx *gofr.Context, movieID int) error {
	query := "UPDATE " + tableName + " SET deletedAt=CURRENT_TIMESTAMP where (id=? AND deletedAt is NULL);"
	row, err := ctx.DB().ExecContext(ctx, query, movieID)

	if err != nil {
		return gofrErrors.DB{Err: err}
	}

	n, err := row.RowsAffected()

	if err != nil {
		return gofrErrors.DB{Err: err}
	}

	if n == 0 {
		return gofrErrors.Error("movie of that id is not present")
	}

	return nil
}
