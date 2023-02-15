package services

import (
	"Assignment_4/models"
	"context"
)

type MovieServicer interface {
	Get(ctx context.Context, movieId int) ([]models.Movie, error)
	Post(ctx context.Context, o models.Movie) (*models.Movie, error)
}
