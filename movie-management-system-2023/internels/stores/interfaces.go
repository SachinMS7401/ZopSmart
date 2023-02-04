package stores

import (
	"context"

	"github.com/go-training/movie-management-system-2023/internels/filters"
	"github.com/go-training/movie-management-system-2023/models"
)

type MovieStorer interface {
	GetAll(ctx context.Context) ([]models.Movie, error)
	Get(ctx context.Context, f filters.Movie) ([]models.Movie, bool, error)
	Post(ctx context.Context, o filters.Details) (int64, error)
	Update(ctx context.Context, movieId int, o filters.UpdateMovie) error
	Delete(ctx context.Context, movieId int) error
}
