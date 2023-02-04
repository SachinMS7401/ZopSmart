package services

import (
	"context"
	"github.com/go-training/movie-management-system-2023/internels/filters"

	"github.com/go-training/movie-management-system-2023/models"
)

type MovieManager interface {
	GetAll(ctx context.Context) ([]models.Movie, error)
	Get(ctx context.Context, movieId int) ([]models.Movie, bool, error)
	Post(ctx context.Context, o filters.Details) ([]models.Movie, bool, error)
	Update(ctx context.Context, movieId int, o filters.UpdateMovie) ([]models.Movie, bool, error)
	Delete(ctx context.Context, movieId int) (bool, error)
}
