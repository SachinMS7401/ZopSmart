package Movie

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-training/movie-management-system-2023/internels/filters"
	"github.com/go-training/movie-management-system-2023/internels/stores"
	"github.com/go-training/movie-management-system-2023/models"
)

type service struct {
	store stores.MovieStorer
}

func New(store stores.MovieStorer) *service {
	return &service{store: store}
}

type InvalidParam struct {
	Param string
}

func (i InvalidParam) Error() string {
	return fmt.Sprintf("param '%s' is invalid", i.Param)
}

func (sr *service) GetAll(ctx context.Context) ([]models.Movie, error) {
	//fmt.Println(MovieId)

	res, err := sr.store.GetAll(ctx)
	if err != nil {
		return nil, errors.New("unable to get from database")
	}
	return res, nil
}

func (sr *service) Get(ctx context.Context, MovieId int) ([]models.Movie, bool, error) {
	//fmt.Println(MovieId)
	if MovieId <= 0 {
		return nil, true, InvalidParam{Param: "MovieId"}
	}

	res, deleted, err := sr.store.Get(ctx, filters.Movie{Id: MovieId})
	if err != nil {
		return nil, true, errors.New("unable to get from database")
	}
	return res, deleted, nil
}

func (sr *service) Post(ctx context.Context, m filters.Details) ([]models.Movie, bool, error) {
	id, err := sr.store.Post(ctx, m)

	if err != nil {
		fmt.Println("Error in the store ", err)
		return nil, true, err
	}
	var id1 = int(id)
	res, deleted, err := sr.store.Get(ctx, filters.Movie{Id: id1})
	fmt.Println(id)
	return res, deleted, nil
}

func (sr *service) Update(ctx context.Context, movieId int, m filters.UpdateMovie) ([]models.Movie, bool, error) {
	err := sr.store.Update(ctx, movieId, m)
	if err != nil {
		fmt.Println("Error in the store ", err)
		return nil, true, err
	}
	res, deleted, err := sr.store.Get(ctx, filters.Movie{Id: movieId})
	return res, deleted, nil
}

func (sr *service) Delete(ctx context.Context, movieId int) (bool, error) {
	err := sr.store.Delete(ctx, movieId)
	if err != nil {
		fmt.Println("Error in the store ", err)
		return false, err
	}
	return true, nil
}
