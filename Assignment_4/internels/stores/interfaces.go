package stores

import (
	"Assignment_4/internels/filters"
	"Assignment_4/models"
	"context"
)

type MovieStorer interface {
	Get(ctx context.Context, f filters.Movie) ([]models.Movie, error)
	Post(ctx context.Context, m models.Movie) (*models.Movie, error)
}
