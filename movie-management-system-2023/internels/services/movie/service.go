package movie

import (
	"strings"

	gofrErrors "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"github.com/go-training/movie-management-system-2023/internels/models"
	"github.com/go-training/movie-management-system-2023/internels/stores"
)

type Service struct {
	store stores.MovieStorer
}

func New(store stores.MovieStorer) *Service {
	return &Service{store: store}
}

func validatePostBody(movie *models.Movie) (err []error) {
	if movie.Rating < 0 || movie.Rating > 5 {
		err = append(err, gofrErrors.MissingParam{Param: []string{"invalid rating"}})
	}

	if strings.TrimSpace(movie.Name) == "" {
		err = append(err, gofrErrors.MissingParam{Param: []string{"empty name"}})
	}

	return err
}

func (sr *Service) GetAll(ctx *gofr.Context) ([]models.Movie, error) {
	res, err := sr.store.GetAll(ctx)

	if err != nil {
		return nil, gofrErrors.Error("error while getting from store layer")
	}

	return res, nil
}

func (sr *Service) Get(ctx *gofr.Context, movieID int) (models.Movie, bool, error) {
	if movieID <= 0 {
		return models.Movie{}, true, &gofrErrors.InvalidParam{Param: []string{"Invalid id"}}
	}

	res, deleted, err := sr.store.Get(ctx, movieID)

	if err != nil {
		return models.Movie{}, true, gofrErrors.Error("error while getting from store layer")
	}

	return res, deleted, nil
}

func (sr *Service) Post(ctx *gofr.Context, mov *models.Movie) (models.Movie, bool, error) {
	invalid := validatePostBody(mov)

	if invalid != nil {
		return models.Movie{}, true, gofrErrors.MultipleErrors{Errors: invalid}
	}

	res, deleted, err := sr.store.Post(ctx, mov)

	if err != nil {
		return models.Movie{}, true, gofrErrors.Error("error while posting in store layer")
	}

	return res, deleted, nil
}

func (sr *Service) Update(ctx *gofr.Context, movieID int, mov *models.UpdateMovie) (models.Movie, bool, error) {
	if mov.Rating < 0 || mov.Rating > 5 {
		return models.Movie{}, true, gofrErrors.Error("invalid while posting in store layer")
	}

	res, deleted, err := sr.store.Update(ctx, movieID, *mov)

	if err != nil {
		return models.Movie{}, true, gofrErrors.Error("error while updating in store layer")
	}

	return res, deleted, nil
}

func (sr *Service) Delete(ctx *gofr.Context, movieID int) (bool, error) {
	err := sr.store.Delete(ctx, movieID)

	if err != nil {
		return false, gofrErrors.Error("error while deleting in store layer")
	}

	return true, nil
}
