package Movie

import (
	"Assignment_4/internels/filters"
	"Assignment_4/internels/stores"
	"Assignment_4/models"
	"context"
	"errors"
	"fmt"
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

func (s *service) Get(ctx context.Context, MovieId int) ([]models.Movie, error) {
	if MovieId <= 0 {
		return nil, InvalidParam{Param: "MovieId"}
	}
	fmt.Println(MovieId)

	res, err := s.store.Get(ctx, filters.Movie{Id: MovieId})

	if err != nil {
		return nil, errors.New("unable to get from database")
	}
	return res, nil
}

func (s *service) Post(ctx context.Context, m models.Movie) (*models.Movie, error) {

	fmt.Println("Reached service layer ")
	res, err := s.store.Post(ctx, m)

	if err != nil {
		fmt.Println("Error in the store ", err)
		return nil, err
	}

	return res, nil
}
