package movie

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"

	"github.com/go-training/movie-management-system-2023/internels/models"
	"github.com/go-training/movie-management-system-2023/internels/stores"
)

func initializeTest(t *testing.T) (sqlmock.Sqlmock, *gofr.Context, stores.MovieStorer) {
	db, mockDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("error in creating mockdb: %v", err)
	}

	app := gofr.New()
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	ctx.DataStore = datastore.DataStore{ORM: db}

	store := New()

	return mockDB, ctx, store
}

func TestStore_Get(t *testing.T) {
	currentTime := time.Now()
	mock, ctx, h := initializeTest(t)
	testCases := []struct {
		description string
		id          int
		expOut      *models.Movie
		expBool     bool
		expError    error
		mockCall    []interface{}
	}{
		{description: "Successfully get the movie details",
			id: 1,
			expOut: &models.Movie{ID: 1,
				Name:         "shivu",
				Genre:        "comedy",
				Rating:       5,
				ReleasedDate: "2019-09-05",
				UpdatedAt:    currentTime,
				CreatedAt:    currentTime,
				Plot:         "director haru based on war",
				Released:     true,
			},
			expBool:  false,
			expError: nil,
			mockCall: []interface{}{mock.ExpectQuery("select id,name,genre,rating,releasedDate," +
				"updatedAt,createdAt,plot,released," +
				"deletedAt from " + tableName + " where (deletedAt is NULL AND id = ?);").WithArgs(1).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "genre", "rating", "releasedDate", "updatedAt",
					"createdAt", "plot", "released", "deletedAt"}).AddRow(1, "shivu", "comedy", 5, "2019-09-05",
					currentTime, currentTime, "director haru based on war", true, sql.NullString{})),
			},
		},
		{
			description: "Error case 1",
			id:          2,
			expOut:      &models.Movie{},
			expBool:     true,
			expError:    sql.ErrNoRows,
			mockCall: []interface{}{mock.ExpectQuery("select id,name,genre,rating,releasedDate," +
				"updatedAt,createdAt,plot,released," +
				"deletedAt from " + tableName + " where (deletedAt is NULL AND id = ?);").WithArgs(2).
				WillReturnError(sql.ErrNoRows),
			},
		},
		{description: "Error case 2",
			id:       1,
			expOut:   &models.Movie{},
			expBool:  true,
			expError: errors.New("sql: expected 9 destination arguments in Scan, not 10"),
			mockCall: []interface{}{mock.ExpectQuery("select id,name,genre,rating,releasedDate," +
				"updatedAt,createdAt,plot,released," +
				"deletedAt from " + tableName + " where (deletedAt is NULL AND id = ?);").WithArgs(1).
				WillReturnRows(sqlmock.NewRows([]string{"id", "genre", "rating", "releasedDate", "updatedAt",
					"createdAt", "plot", "released", "deletedAt"}).AddRow(
					1, "comedy", 5, "2019-09-05", currentTime, currentTime, "director haru based on war",
					true, sql.NullString{})),
			},
		},
	}

	for _, tc := range testCases {
		result, deleted, err := h.Get(ctx, tc.id)
		if err != nil && err.Error() != tc.expError.Error() {
			t.Errorf("expected %v got %v", tc.expError, err)
		} else if !reflect.DeepEqual(deleted, tc.expBool) {
			t.Errorf("got this %v expected this %v", deleted, tc.expBool)
		} else if !reflect.DeepEqual(result, *tc.expOut) {
			t.Errorf("got this %v expected this %v", result, *tc.expOut)
		}
	}
}

func TestStore_Post(t *testing.T) {
	currentTime := time.Now()
	mock, ctx, h := initializeTest(t)
	testCases := []struct {
		description string
		movie       models.Movie
		expOut      *models.Movie
		expBool     bool
		expError    error
		mockCall    []interface{}
	}{
		{description: "Successfully post the movie details",
			movie: models.Movie{
				Name:         "Karthik",
				Genre:        "Action",
				Rating:       4.5,
				ReleasedDate: "2021-08-02",
				Plot:         "Directed by Bharathi ,its base upon the war ",
				Released:     true,
			},
			expOut: &models.Movie{ID: 5,
				Name:         "Karthik",
				Genre:        "Action",
				Rating:       4.5,
				ReleasedDate: "2021-08-02",
				UpdatedAt:    currentTime,
				CreatedAt:    currentTime,
				Plot:         "Directed by Bharathi ,its base upon the war ",
				Released:     true,
			},
			expBool:  false,
			expError: nil,
			mockCall: []interface{}{mock.ExpectExec("INSERT INTO "+tableName+"(name,genre,rating,"+
				"releasedDate,updatedAt,createdAt,plot,released) "+
				"VALUES(?,?,?,?,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP,?,?)").WithArgs(
				"Karthik", "Action", 4.5, "2021-08-02", "Directed by Bharathi ,its base upon the war ", true).
				WillReturnResult(sqlmock.NewResult(5, 1)), mock.ExpectQuery(
				"select id,name,genre,rating,releasedDate,updatedAt,createdAt,plot,released," +
					"deletedAt from " + tableName + " where (deletedAt is NULL AND id = ?);").WithArgs(5).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "genre", "rating", "releasedDate",
					"updatedAt", "createdAt", "plot", "released", "deletedAt"}).AddRow(
					5, "Karthik", "Action", 4.5, "2021-08-02", currentTime, currentTime,
					"Directed by Bharathi ,its base upon the war ", true, sql.NullString{})),
			},
		},
		{
			description: "Error case 1",
			movie: models.Movie{
				Name:         "Ajay",
				Genre:        "Action",
				Rating:       4.5,
				ReleasedDate: "2021-08-02",
				Plot:         "Directed by Bharathi ,its base upon the war ",
				Released:     true,
			},
			expOut:   &models.Movie{},
			expBool:  true,
			expError: sql.ErrConnDone,
			mockCall: []interface{}{mock.ExpectExec("INSERT INTO "+tableName+"(name,genre,rating,"+
				"releasedDate,updatedAt,createdAt,plot,released) "+
				"VALUES(?,?,?,?,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP,?,?)").WithArgs(
				"Ajay", "Action", 4.5, "2021-08-02", "Directed by Bharathi ,its base upon the war ", true).
				WillReturnError(sql.ErrConnDone),
			},
		},
		{
			description: "Error case 2",
			movie: models.Movie{
				Name:         "Ajay",
				Genre:        "Action",
				Rating:       4.5,
				ReleasedDate: "2021-08-02",
				Plot:         "Directed by Bharathi ,its base upon the war ",
				Released:     true,
			},
			expOut:   &models.Movie{},
			expBool:  true,
			expError: errors.New("error"),
			mockCall: []interface{}{mock.ExpectExec("INSERT INTO "+tableName+"(name,genre,"+
				"rating,releasedDate,updatedAt,createdAt,plot,released) "+
				"VALUES(?,?,?,?,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP,?,?)").WithArgs(
				"Ajay", "Action", 4.5, "2021-08-02", "Directed by Bharathi ,its base upon the war ", true).
				WillReturnResult(sqlmock.NewErrorResult(errors.New("error"))),
			},
		},
	}

	for _, tc := range testCases {
		result, notDel, err := h.Post(ctx, &tc.movie)

		if err != nil {
			if err.Error() != tc.expError.Error() {
				t.Errorf("%v %v", err.Error(), tc.expError)
			}
		} else if !reflect.DeepEqual(notDel, tc.expBool) {
			t.Errorf("got this %v expected this %v", notDel, tc.expBool)
		} else if !reflect.DeepEqual(result, *tc.expOut) {
			t.Errorf("got this %v expected this %v", result, *tc.expOut)
		}
	}
}

func TestStore_GetAll(t *testing.T) {
	currentTime := time.Now()
	mock, ctx, h := initializeTest(t)
	movies := []models.Movie{
		{
			ID:           1,
			Name:         "shivu",
			Genre:        "comedy",
			Rating:       5,
			ReleasedDate: "2019-09-05",
			UpdatedAt:    currentTime,
			CreatedAt:    currentTime,
			Plot:         "director haru based on war",
			Released:     true,
		},
	}

	testCases := []struct {
		description string
		expOut      []models.Movie
		expError    error
		mockCall    []interface{}
	}{
		{description: "Successfully get the movie details",
			expOut:   movies,
			expError: nil,
			mockCall: []interface{}{mock.ExpectQuery("select * from " + tableName +
				" where deletedAt is NULL;").WithArgs().
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "genre", "rating", "releasedDate",
					"updatedAt", "createdAt", "plot", "released", "deletedAt"}).AddRow(
					1, "shivu", "comedy", 5, "2019-09-05", currentTime, currentTime,
					"director haru based on war", true, sql.NullString{})),
			},
		},
		{description: "Error case 1",
			expOut:   []models.Movie{},
			expError: sql.ErrNoRows,
			mockCall: []interface{}{mock.ExpectQuery("select * from " + tableName +
				" where deletedAt is NULL;").WillReturnError(sql.ErrNoRows)},
		},
		{description: "Error case 2",
			expOut:   []models.Movie{},
			expError: errors.New("sql: expected 9 destination arguments in Scan, not 10"),
			mockCall: []interface{}{mock.ExpectQuery("select * from " + tableName +
				" where deletedAt is NULL;").WillReturnRows(sqlmock.NewRows([]string{"id", "genre", "rating",
				"releasedDate", "updatedAt", "createdAt", "plot", "released", "deletedAt"}).AddRow(
				1, "comedy", 5, "2019-09-05", currentTime, currentTime,
				"director haru based on war", true, sql.NullString{}))},
		},
	}

	for _, tc := range testCases {
		result, err := h.GetAll(ctx)
		if err != nil {
			if err.Error() != tc.expError.Error() {
				t.Errorf("%v %v", err.Error(), tc.expError)
			}
		} else {
			if !reflect.DeepEqual(result, tc.expOut) {
				t.Errorf("got this %v expected this %v", result, tc.expOut)
			}
		}
	}
}

func TestStore_Update(t *testing.T) {
	currentTime := time.Now()
	mock, ctx, h := initializeTest(t)
	testCases := []struct {
		description string
		id          int
		putDetails  models.UpdateMovie
		expOut      *models.Movie
		expBool     bool
		expError    error
		mockCall    []interface{}
	}{
		{description: "Successfully updated the movie details",
			id: 1,
			putDetails: models.UpdateMovie{
				Rating:       5,
				ReleasedDate: "2019-09-05",
				Plot:         "director haru based on war",
			},
			expOut: &models.Movie{ID: 1,
				Name:         "shivu",
				Genre:        "comedy",
				Rating:       5,
				ReleasedDate: "2019-09-05",
				UpdatedAt:    currentTime,
				CreatedAt:    currentTime,
				Plot:         "director haru based on war",
				Released:     true,
			},
			expBool:  false,
			expError: nil,
			mockCall: []interface{}{mock.ExpectExec(
				"UPDATE "+tableName+" SET updatedAt=CURRENT_TIMESTAMP,rating=?,plot=?,releasedDate=? WHERE id=?;").
				WithArgs(float64(5), "director haru based on war", "2019-09-05", 1).WillReturnResult(
				sqlmock.NewResult(0, 0)), mock.ExpectQuery(
				"select id,name,genre,rating,releasedDate,updatedAt,createdAt,plot,released,deletedAt " +
					"from " + tableName + " where (deletedAt is NULL AND id = ?);").WithArgs(1).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "genre", "rating", "releasedDate",
					"updatedAt", "createdAt", "plot", "released", "deletedAt"}).AddRow(
					1, "shivu", "comedy", 5, "2019-09-05", currentTime, currentTime,
					"director haru based on war", true, sql.NullString{})),
			},
		},
		{description: "Error case 1",
			id: 1,
			putDetails: models.UpdateMovie{
				Rating:       5,
				ReleasedDate: "2019-09-05",
				Plot:         "director haru based on war",
			},
			expOut:   &models.Movie{},
			expBool:  true,
			expError: sql.ErrNoRows,
			mockCall: []interface{}{mock.ExpectExec(
				"UPDATE "+tableName+" SET updatedAt=CURRENT_TIMESTAMP,rating=?,plot=?,releasedDate=? WHERE id=?;").
				WithArgs(float64(5), "director haru based on war", "2019-09-05", 1).
				WillReturnError(sql.ErrNoRows),
			},
		},
		{description: "Error case 2",
			id: 1,
			putDetails: models.UpdateMovie{
				Rating:       5,
				ReleasedDate: "2019-09-05",
				Plot:         "director haru based on war",
			},
			expOut:   &models.Movie{},
			expBool:  true,
			expError: sql.ErrNoRows,
			mockCall: []interface{}{mock.ExpectExec(
				"UPDATE "+tableName+" SET updatedAt=CURRENT_TIMESTAMP,rating=?,plot=?,releasedDate=? WHERE id=?;").
				WithArgs(float64(5), "director haru based on war", "2019-09-05", 1).
				WillReturnResult(sqlmock.NewResult(0, 0)),
				mock.ExpectQuery(
					"select id,name,genre,rating,releasedDate,updatedAt,createdAt,plot,released,deletedAt " +
						"from " + tableName + " where (deletedAt is NULL AND id = ?);").WithArgs(1).
					WillReturnError(sql.ErrNoRows),
			},
		},
	}

	for _, tc := range testCases {
		result, notDel, err := h.Update(ctx, tc.id, tc.putDetails)

		if err != nil {
			if err.Error() != tc.expError.Error() {
				t.Errorf("%v %v", err.Error(), tc.expError)
			}
		} else if !reflect.DeepEqual(notDel, tc.expBool) {
			t.Errorf("got this %v expected this %v", notDel, tc.expBool)
		} else if !reflect.DeepEqual(result, *tc.expOut) {
			t.Errorf("got this %v expected this %v", result, *tc.expOut)
		}
	}
}

func TestStore_Delete(t *testing.T) {
	mock, ctx, h := initializeTest(t)
	testCases := []struct {
		description string
		movieID     int
		expError    error
		mockCall    *sqlmock.ExpectedExec
	}{
		{description: "Successfully deleted the movie details",
			movieID:  10,
			expError: nil,
			mockCall: mock.ExpectExec("UPDATE " + tableName + " SET deletedAt=CURRENT_TIMESTAMP where " +
				"(id=? AND deletedAt is NULL);").WithArgs(10).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{description: "Error case 1",
			movieID:  10,
			expError: sql.ErrNoRows,
			mockCall: mock.ExpectExec("UPDATE " + tableName + " SET deletedAt=CURRENT_TIMESTAMP where " +
				"(id=? AND deletedAt is NULL);").WithArgs(10).WillReturnError(sql.ErrNoRows),
		},
	}

	for _, tc := range testCases {
		err := h.Delete(ctx, tc.movieID)
		if err != nil {
			if err.Error() != tc.expError.Error() {
				t.Errorf("expected %v got %v", tc.expError, err)
			}
		}
	}
}
