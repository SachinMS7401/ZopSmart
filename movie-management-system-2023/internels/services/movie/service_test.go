package movie

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	gofrerrors "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	gomock "github.com/golang/mock/gomock"

	"github.com/go-training/movie-management-system-2023/internels/models"
	"github.com/go-training/movie-management-system-2023/internels/services"
	"github.com/go-training/movie-management-system-2023/internels/stores"
)

func initializeTest(t *testing.T) (*gofr.Context, *stores.MockMovieStorer, services.MovieManager) {
	app := gofr.New()

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	ctrl := gomock.NewController(t)
	mockStore := stores.NewMockMovieStorer(ctrl)

	return ctx, mockStore, New(mockStore)
}

func TestService_Post(t *testing.T) {
	ctx, storeHandler, serviceHandler := initializeTest(t)
	movie := models.Movie{
		ID:           1,
		Name:         "shivu",
		Genre:        "comedy",
		Rating:       5,
		ReleasedDate: "2019-09-05",
		UpdatedAt:    time.Now().UTC(),
		CreatedAt:    time.Now().UTC(),
		Plot:         "director haru based on war",
		Released:     true}
	testcases := []struct {
		body     models.Movie
		expOut   models.Movie
		expBool  bool
		err      error
		mockCall *gomock.Call
	}{
		{
			body: models.Movie{Name: "shivu",
				Genre:        "comedy",
				Rating:       5,
				ReleasedDate: "2019-09-05",
				Plot:         "director haru based on war",
				Released:     true},
			expOut:   movie,
			expBool:  false,
			err:      nil,
			mockCall: storeHandler.EXPECT().Post(ctx, gomock.Any()).Return(movie, false, nil),
		},
		{
			body: models.Movie{Name: "shivu",
				Genre:        "comedy",
				Rating:       5,
				ReleasedDate: "2019-09-05",
				Plot:         "director haru based on war",
			},
			expOut:  models.Movie{},
			expBool: true,
			err:     gofrerrors.Error("error while posting in store layer"),
			mockCall: storeHandler.EXPECT().Post(ctx, gomock.Any()).Return(models.Movie{}, true,
				gofrerrors.Error("error while posting in store layer")),
		},
		{
			body: models.Movie{Name: " ",
				Genre:        "comedy",
				Rating:       5.5,
				ReleasedDate: "2019-09-05",
				Plot:         "director haru based on war",
			},
			expOut:  models.Movie{},
			expBool: true,
			err: errors.New("Parameter invalid rating is required for " +
				"this request\nParameter empty name is required for this request"),
			mockCall: nil,
		},
	}

	for _, tc := range testcases {
		actualMov, deleted, err := serviceHandler.Post(ctx, &tc.body)
		if err != nil {
			if err.Error() != tc.err.Error() {
				t.Errorf("expecting %v but we get %v", err.Error(), tc.err.Error())
			}
		} else if !reflect.DeepEqual(actualMov, tc.expOut) && tc.expBool != deleted {
			t.Errorf("expecting %v but we get %v", actualMov, tc.expOut)
		}
	}
}

func TestService_Get(t *testing.T) {
	ctx, storeHandler, serviceHandler := initializeTest(t)
	movie := models.Movie{
		ID:           1,
		Name:         "shivu",
		Genre:        "comedy",
		Rating:       5,
		ReleasedDate: "2019-09-05",
		UpdatedAt:    time.Now().UTC(),
		CreatedAt:    time.Now().UTC(),
		Plot:         "director haru based on war",
		Released:     true}
	testcases := []struct {
		id       int
		expOut   models.Movie
		expBool  bool
		err      error
		mockCall *gomock.Call
	}{
		{
			id:       1,
			expOut:   movie,
			expBool:  false,
			err:      nil,
			mockCall: storeHandler.EXPECT().Get(ctx, gomock.Any()).Return(movie, false, nil),
		},
		{
			id:      1,
			expOut:  models.Movie{},
			expBool: true,
			err:     gofrerrors.Error("error while getting from store layer"),
			mockCall: storeHandler.EXPECT().Get(ctx, gomock.Any()).Return(models.Movie{},
				true, gofrerrors.Error("error while getting from store layer")),
		},
		{
			id:       -1,
			expOut:   models.Movie{},
			expBool:  true,
			err:      gofrerrors.InvalidParam{Param: []string{"Invalid id"}},
			mockCall: nil,
		},
	}

	for _, tc := range testcases {
		actualMov, deleted, err := serviceHandler.Get(ctx, tc.id)
		if err != nil {
			if err.Error() != tc.err.Error() {
				t.Errorf("expecting %v but we get %v", err.Error(), tc.err.Error())
			}
		} else if !reflect.DeepEqual(actualMov, tc.expOut) && tc.expBool != deleted {
			t.Errorf("expecting %v but we get %v", actualMov, tc.expOut)
		}
	}
}

func TestService_Update(t *testing.T) {
	ctx, storeHandler, serviceHandler := initializeTest(t)
	movie := models.Movie{
		ID:           1,
		Name:         "shivu",
		Genre:        "comedy",
		Rating:       5,
		ReleasedDate: "2019-09-05",
		UpdatedAt:    time.Now().UTC(),
		CreatedAt:    time.Now().UTC(),
		Plot:         "director haru based on war",
		Released:     true}
	testcases := []struct {
		id       int
		body     models.UpdateMovie
		expOut   models.Movie
		expBool  bool
		err      error
		mockCall *gomock.Call
	}{
		{
			id:       1,
			body:     models.UpdateMovie{Rating: 5, ReleasedDate: "2019-09-05", Plot: "director haru based on war"},
			expOut:   movie,
			expBool:  false,
			err:      nil,
			mockCall: storeHandler.EXPECT().Update(ctx, gomock.Any(), gomock.Any()).Return(movie, false, nil),
		},
		{
			id:       1,
			body:     models.UpdateMovie{Rating: 5.5, ReleasedDate: "2019-09-05", Plot: "director haru based on war"},
			expOut:   models.Movie{},
			expBool:  true,
			err:      gofrerrors.Error("invalid while posting in store layer"),
			mockCall: nil,
		},
		{
			id:      1,
			body:    models.UpdateMovie{Rating: 5, ReleasedDate: "2019-09-05", Plot: "director haru based on war"},
			expOut:  models.Movie{},
			expBool: true,
			err:     gofrerrors.Error("error while updating in store layer"),
			mockCall: storeHandler.EXPECT().Update(ctx, gomock.Any(), gomock.Any()).Return(models.Movie{},
				true, gofrerrors.Error("error while updating in store layer")),
		},
	}

	for _, tc := range testcases {
		actualMov, deleted, err := serviceHandler.Update(ctx, tc.id, &tc.body)
		if err != nil {
			if err.Error() != tc.err.Error() {
				t.Errorf("expecting %v but we get %v", err.Error(), tc.err.Error())
			}
		} else if !reflect.DeepEqual(actualMov, tc.expOut) && tc.expBool != deleted {
			t.Errorf("expecting %v but we get %v", actualMov, tc.expOut)
		}
	}
}

func TestService_Delete(t *testing.T) {
	ctx, storeHandler, serviceHandler := initializeTest(t)
	testcases := []struct {
		id       int
		expBool  bool
		err      error
		mockCall *gomock.Call
	}{
		{
			id:       1,
			expBool:  true,
			err:      nil,
			mockCall: storeHandler.EXPECT().Delete(ctx, 1).Return(nil),
		},
		{
			id:       2,
			expBool:  false,
			err:      gofrerrors.Error("error while deleting in store layer"),
			mockCall: storeHandler.EXPECT().Delete(ctx, 2).Return(gofrerrors.Error("error while deleting in store layer")),
		},
	}

	for _, tc := range testcases {
		deleted, err := serviceHandler.Delete(ctx, tc.id)
		if err != nil {
			if err.Error() != tc.err.Error() {
				t.Errorf("expecting %v but we get %v", err.Error(), tc.err.Error())
			}
		} else if !reflect.DeepEqual(deleted, tc.expBool) {
			t.Errorf("expecting %v but we get %v", deleted, tc.expBool)
		}
	}
}

func TestService_GetAll(t *testing.T) {
	ctx, storeHandler, serviceHandler := initializeTest(t)
	movie := models.Movie{
		ID:           1,
		Name:         "shivu",
		Genre:        "comedy",
		Rating:       5,
		ReleasedDate: "2019-09-05",
		UpdatedAt:    time.Now().UTC(),
		CreatedAt:    time.Now().UTC(),
		Plot:         "director haru based on war",
		Released:     true}

	var movies []models.Movie

	movies = append(movies, movie)
	testcases := []struct {
		expOut   []models.Movie
		err      error
		mockCall *gomock.Call
	}{
		{
			expOut:   movies,
			err:      nil,
			mockCall: storeHandler.EXPECT().GetAll(ctx).Return(movies, nil),
		},
		{
			expOut:   nil,
			err:      gofrerrors.Error("error while getting from store layer"),
			mockCall: storeHandler.EXPECT().GetAll(ctx).Return(nil, gofrerrors.Error("error while getting from store layer")),
		},
	}

	for _, tc := range testcases {
		actualMovies, err := serviceHandler.GetAll(ctx)
		if err != nil {
			if err.Error() != tc.err.Error() {
				t.Errorf("expecting %v but we get %v", err.Error(), tc.err.Error())
			}
		} else if !reflect.DeepEqual(actualMovies, tc.expOut) {
			t.Errorf("expecting %v but we get %v", actualMovies, tc.expOut)
		}
	}
}
