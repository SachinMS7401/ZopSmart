package movie

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-training/movie-management-system-2023/internels/filters"
	"github.com/go-training/movie-management-system-2023/models"
)

const tableName = "sample5"

type store struct {
	db *sql.DB
}

func New(db *sql.DB) *store {
	return &store{db: db}
}

func (s *store) GetAll(ctx context.Context) ([]models.Movie, error) {

	//query := "select id,name,genre,rating,releasedDate,updatedAt,createdAt,plot,released,deletedAt from " + tableName
	query := "select * from " + tableName + " where deletedAt is NULL"
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var movies []models.Movie

	for rows.Next() {
		var m models.Movie

		err := rows.Scan(&m.Id, &m.Name, &m.Genre, &m.Rating, &m.ReleasedDate, &m.UpdatedAt, &m.CreatedAt, &m.Plot, &m.Released, &m.DeletedAt)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		movies = append(movies, m)
	}
	fmt.Println("Data in movies is ", movies)
	return movies, nil
}

func (s *store) Get(ctx context.Context, f filters.Movie) ([]models.Movie, bool, error) {
	var details []interface{}

	if f.Id > 0 {
		details = append(details, f.Id)
	}
	query := "select id,name,genre,rating,releasedDate,updatedAt,createdAt,plot,released,deletedAt from " + tableName + " where (deletedAt is NULL AND id = ?)"
	rows, err := s.db.Query(query, details...)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()
	var movies []models.Movie
	notDel := true
	for rows.Next() {
		var m models.Movie
		err := rows.Scan(&m.Id, &m.Name, &m.Genre, &m.Rating, &m.ReleasedDate, &m.UpdatedAt, &m.CreatedAt, &m.Plot, &m.Released, &m.DeletedAt)
		if err != nil {
			fmt.Println(err)
			return nil, false, err
		}
		movies = append(movies, m)
		notDel = false
	}
	fmt.Println("Data in movies is ", movies)
	return movies, notDel, nil
}

func (s *store) Post(ctx context.Context, m filters.Details) (int64, error) {
	var details []interface{}

	query := "INSERT INTO " + tableName + "(name,genre,rating,releasedDate,updatedAt,createdAt,plot,released) VALUES(?,?,?,?,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP,?,?)"
	details = append(details, m.Name, m.Genre, m.Rating, m.ReleasedDate, m.Plot, m.Released)
	row, err := s.db.Exec(query, details...)
	fmt.Println(row)
	if err != nil {
		fmt.Println("Error while inserting into database")
		return 0, err
	}
	id, err := row.LastInsertId()

	if err != nil {
		fmt.Println("Error while getting the last inserted id")
		return 0, err
	}
	n, _ := row.RowsAffected()

	if n > 0 {
		fmt.Println(n, " number of rows affected in the database")
	}
	var mov filters.Details

	mov.Name = m.Name
	mov.Genre = m.Genre
	mov.Rating = m.Rating
	mov.ReleasedDate = m.ReleasedDate
	mov.Plot = m.Plot
	mov.Released = m.Released

	fmt.Println("Updated values are: ", mov)
	fmt.Println(id)
	return id, nil

}

func (s *store) Update(ctx context.Context, movieId int, m filters.UpdateMovie) error {
	var patch []string
	var details []interface{}
	if m.Rating > 0 {
		patch = append(patch, "rating=?")
		details = append(details, m.Rating)
	}
	if m.Plot != "" {
		patch = append(patch, "plot=?")
		details = append(details, m.Plot)
	}
	if m.ReleasedDate != "" {
		patch = append(patch, "releasedDate=?")
		details = append(details, m.ReleasedDate)
	}
	query := "UPDATE sample5 SET updatedAt=CURRENT_TIMESTAMP,"
	if len(patch) > 0 {
		query += strings.Join(patch, ",") + " WHERE id=?"
	}
	details = append(details, movieId)
	fmt.Println(query)
	row, err := s.db.ExecContext(ctx, query, details...)
	if err != nil {
		fmt.Println("Error while updating in the database")
		return err
	}

	if err != nil {
		fmt.Println("Error while getting the last inserted id")
		return err
	}
	n, _ := row.RowsAffected()

	if n > 0 {
		fmt.Println(n, " number of rows affected in the database")
	}
	var mov filters.UpdateMovie
	mov.Rating = m.Rating
	mov.ReleasedDate = m.ReleasedDate
	mov.Plot = m.Plot
	fmt.Println(mov)
	return nil
}

func (s *store) Delete(ctx context.Context, movieId int) error {
	query := "UPDATE sample5 SET deletedAt=CURRENT_TIMESTAMP where id=?"
	row, err := s.db.ExecContext(ctx, query, movieId)
	if err != nil {
		fmt.Println("Error while updating in the database")
		return err
	}
	n, _ := row.RowsAffected()

	if n > 0 {
		fmt.Println(n, " number of rows affected in the database")
	}
	return nil
}
