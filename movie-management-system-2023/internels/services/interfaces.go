package services

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"github.com/go-training/movie-management-system-2023/internels/models"
)

type MovieManager interface {
	GetAll(ctx *gofr.Context) ([]models.Movie, error)
	Get(ctx *gofr.Context, movieID int) (models.Movie, bool, error)
	Post(ctx *gofr.Context, mov *models.Movie) (models.Movie, bool, error)
	Update(ctx *gofr.Context, movieID int, upMovie *models.UpdateMovie) (models.Movie, bool, error)
	Delete(ctx *gofr.Context, movieID int) (bool, error)
}
