package movie

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"developer.zopsmart.com/go/gofr/pkg/gofr/types"

	"github.com/go-training/movie-management-system-2023/internels/models"
	"github.com/go-training/movie-management-system-2023/internels/services"
	"github.com/golang/mock/gomock"
)

func TestHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceMock := services.NewMockMovieManager(ctrl)
	mockHandler := New(serviceMock)
	currentTime := time.Now().UTC()
	app := gofr.New()

	movie := models.Movie{ID: 1,
		Name: "RRR", Genre: "action", Rating: 5,
		ReleasedDate: "2021-05-27",
		UpdatedAt:    currentTime, CreatedAt: currentTime,
		Plot: "Love about nation and friendship", Released: true}
	tests := []struct {
		input  []byte
		output interface{}
		expErr error
		mock   *gomock.Call
	}{
		{
			input: []byte(`{
		   "name": "RRR",
		   "genre": "Action",
		   "rating": 5,
		   "releasedDate": "2021-05-27",
		   "plot": "Love about nation and friendship",
		   "released": true
		}`),
			output: movie,
			expErr: nil,
			mock: serviceMock.EXPECT().Post(gomock.Any(), &models.Movie{
				Name:         "RRR",
				Genre:        "Action",
				Rating:       5,
				ReleasedDate: "2021-05-27",
				Plot:         "Love about nation and friendship",
				Released:     true,
			}).Return(movie, false, nil),
		},
		{
			input: []byte(`{
		"name": "RRR",
		"genre": "Action",
		"rating": 5,
		"releasedDate": "2021-05-27",
		"plot": "Love about nation and friendship",
		"released": true
		}`),
			output: models.Movie{},
			expErr: errors.Error("bad request"),
			mock: serviceMock.EXPECT().Post(gomock.Any(), &models.Movie{
				Name:         "RRR",
				Genre:        "Action",
				Rating:       5,
				ReleasedDate: "2021-05-27",
				Plot:         "Love about nation and friendship",
				Released:     true,
			}).Return(models.Movie{}, true, errors.Error("error")),
		},
	}

	for _, tc := range tests {
		w := httptest.NewRecorder()

		r := httptest.NewRequest(http.MethodPost, "/movies", bytes.NewReader(tc.input))
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)
		resp, err := mockHandler.Create(ctx)

		if err != nil {
			if !reflect.DeepEqual(err.Error(), tc.expErr.Error()) {
				t.Errorf("expecting %v but we get %v", err.Error(), tc.expErr.Error())
			}
		} else if !reflect.DeepEqual(resp, types.Response{Data: tc.output}) {
			t.Errorf("expecting %v but we get %v", resp, types.Response{Data: tc.output})
		}
	}
}

func TestHandler_Read(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceMock := services.NewMockMovieManager(ctrl)
	mockHandler := New(serviceMock)
	app := gofr.New()
	movie := models.Movie{
		ID: 1, Name: "shivu",
		Genre:        "comedy",
		Rating:       5,
		ReleasedDate: "2019-09-05",
		UpdatedAt:    time.Now().UTC(),
		CreatedAt:    time.Now().UTC(),
		Plot:         "director haru based on war",
		Released:     true,
	}
	testcases := []struct {
		input    string
		expOut   interface{}
		expErr   error
		expBool  bool
		mockCall *gomock.Call
	}{

		{
			input:    "1",
			expOut:   movie,
			expErr:   nil,
			expBool:  false,
			mockCall: serviceMock.EXPECT().Get(gomock.Any(), 1).Return(movie, false, nil),
		},
		{
			input:  "1",
			expOut: models.Movie{},
			expErr: errors.Error("error while getting from service getByID"),
			mockCall: serviceMock.EXPECT().Get(gomock.Any(), 1).Return(models.Movie{},
				true, errors.Error("error from the service layer")),
		},
		{
			input:    "one",
			expOut:   models.Movie{},
			expErr:   errors.Error("param id is invalid"),
			mockCall: nil,
		},
	}

	for _, tc := range testcases {
		w := httptest.NewRecorder()

		r := httptest.NewRequest(http.MethodGet, "/movies", nil)
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)

		ctx.SetPathParams(map[string]string{
			"id": tc.input,
		})

		resp, err := mockHandler.Read(ctx)
		if err != nil {
			if !reflect.DeepEqual(err.Error(), tc.expErr.Error()) {
				t.Errorf("expecting %v but we get %v", err.Error(), tc.expErr.Error())
			}
		} else if !reflect.DeepEqual(resp, types.Response{Data: tc.expOut}) {
			t.Errorf("expecting %v but we get %v", resp, tc.expOut)
		}
	}
}

func TestHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceMock := services.NewMockMovieManager(ctrl)
	mockHandler := New(serviceMock)
	app := gofr.New()
	testcases := []struct {
		id     string
		expOut interface{}
		expErr error
		mock   *gomock.Call
	}{
		{
			id:     "23",
			expOut: http.StatusNoContent,
			expErr: nil,
			mock:   serviceMock.EXPECT().Delete(gomock.Any(), 23).Return(true, nil),
		},
		{
			id:     "23",
			expOut: nil,
			expErr: errors.Error("requested id is not present"),
			mock:   serviceMock.EXPECT().Delete(gomock.Any(), 23).Return(false, errors.Error("error")),
		},
		{
			id:     "one",
			expOut: nil,
			expErr: errors.Error("param id is invalid"),
			mock:   nil,
		},
	}

	for _, tc := range testcases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/movies", nil)
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)
		ctx.SetPathParams(map[string]string{
			"id": tc.id,
		})

		resp, err := mockHandler.Delete(ctx)
		if err != nil {
			if !reflect.DeepEqual(err.Error(), tc.expErr.Error()) {
				t.Errorf("expecting %v but we get %v", err.Error(), tc.expErr.Error())
			}
		} else if !reflect.DeepEqual(resp, types.Response{Data: tc.expOut}) {
			t.Errorf("expecting %v but we get %v", resp, tc.expOut)
		}
	}
}

func TestHandler_Index(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceMock := services.NewMockMovieManager(ctrl)
	mockHandler := New(serviceMock)
	app := gofr.New()
	movie := models.Movie{
		ID: 1, Name: "shiv",
		Genre:        "comedy",
		Rating:       5,
		ReleasedDate: "2019-09-05",
		UpdatedAt:    time.Now().UTC(),
		CreatedAt:    time.Now().UTC(),
		Plot:         "director har based on war",
		Released:     true,
	}

	var movies []models.Movie

	movies = append(movies, movie)
	testcases := []struct {
		expOut   interface{}
		expErr   error
		mockCall *gomock.Call
	}{

		{
			expOut:   movies,
			expErr:   nil,
			mockCall: serviceMock.EXPECT().GetAll(gomock.Any()).Return(movies, nil),
		},
		{
			expOut: []models.Movie{},
			expErr: errors.Error("error while getting from service get"),
			mockCall: serviceMock.EXPECT().GetAll(gomock.Any()).Return([]models.Movie{},
				errors.Error("error from service layer")),
		},
	}

	for _, tc := range testcases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/movies", nil)
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)

		resp, err := mockHandler.Index(ctx)
		if err != nil {
			if !reflect.DeepEqual(err.Error(), tc.expErr.Error()) {
				t.Errorf("expecting %v but we get %v", err.Error(), tc.expErr.Error())
			}
		} else if !reflect.DeepEqual(resp, types.Response{Data: tc.expOut}) {
			t.Errorf("expecting %v but we get %v", resp, tc.expOut)
		}
	}
}

func TestHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceMock := services.NewMockMovieManager(ctrl)
	mockHandler := New(serviceMock)
	app := gofr.New()
	currentTime := time.Now().UTC()
	movie := models.Movie{ID: 1,
		Name: "RRR", Genre: "action", Rating: 4.5,
		ReleasedDate: "2021-05-27",
		UpdatedAt:    currentTime, CreatedAt: currentTime,
		Plot: "Love about nation and friendship", Released: true}
	testcases := []struct {
		id      string
		body    models.UpdateMovie
		expOut  interface{}
		expBool bool
		expErr  error
		mock    *gomock.Call
	}{
		{
			id: "1",
			body: models.UpdateMovie{
				Rating:       4.5,
				ReleasedDate: "2021-05-27",
				Plot:         "Love about nation and friendship",
			},
			expOut:  movie,
			expBool: false,
			expErr:  nil,
			mock: serviceMock.EXPECT().Update(gomock.Any(), 1, &models.UpdateMovie{
				Rating:       4.5,
				ReleasedDate: "2021-05-27",
				Plot:         "Love about nation and friendship",
			}).Return(movie, false, nil),
		},
		{
			id: "1",
			body: models.UpdateMovie{
				Rating:       4.5,
				ReleasedDate: "2021-05-27",
				Plot:         "Love about nation and friendship",
			},
			expOut:  models.Movie{},
			expBool: true,
			expErr:  errors.Error("error from the service layer"),
			mock: serviceMock.EXPECT().Update(gomock.Any(), 1, &models.UpdateMovie{
				Rating:       4.5,
				ReleasedDate: "2021-05-27",
				Plot:         "Love about nation and friendship",
			}).Return(models.Movie{}, true, errors.Error("error from the service layer")),
		},
		{
			id: "one",
			body: models.UpdateMovie{
				Rating:       4.5,
				ReleasedDate: "2021-05-27",
				Plot:         "Love about nation and friendship",
			},
			expOut:  models.Movie{},
			expBool: true,
			expErr:  errors.Error("param id is invalid"),
			mock:    nil,
		},
	}

	for _, tc := range testcases {
		data, _ := json.Marshal(tc.body)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/movies", bytes.NewBuffer(data))
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)
		ctx.SetPathParams(map[string]string{
			"id": tc.id,
		})

		resp, err := mockHandler.Update(ctx)
		if err != nil {
			if !reflect.DeepEqual(err.Error(), tc.expErr.Error()) {
				t.Errorf("expecting %v but we get %v", err.Error(), tc.expErr.Error())
			}
		} else if !reflect.DeepEqual(resp, types.Response{Data: tc.expOut}) {
			t.Errorf("expecting %v but we get %v", resp, tc.expOut)
		}
	}
}
