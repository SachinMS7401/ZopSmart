package movie

import (
	"net/http"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/types"

	"github.com/go-training/movie-management-system-2023/internels/models"
	"github.com/go-training/movie-management-system-2023/internels/services"
)

type Handler struct {
	svc services.MovieManager
}

func New(svc services.MovieManager) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Index(ctx *gofr.Context) (interface{}, error) {
	movie, err := h.svc.GetAll(ctx)

	if err != nil {
		errRes := &errors.Response{
			StatusCode: http.StatusBadRequest,
			Code:       "ERROR",
			Reason:     "error while getting from service get",
		}

		return nil, errRes
	}

	response := types.Response{
		Data: movie,
	}

	return response, nil
}

func (h *Handler) Read(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	movieID, err := strconv.Atoi(id)

	if err != nil {
		errRes := &errors.Response{
			StatusCode: http.StatusBadRequest,
			Code:       "ERROR",
			Reason:     "param id is invalid",
		}

		return nil, errRes
	}

	movie, deleted, err := h.svc.Get(ctx, movieID)

	if err != nil || deleted {
		errRes := &errors.Response{
			StatusCode: http.StatusBadRequest,
			Code:       "ERROR",
			Reason:     "error while getting from service getByID",
		}

		return nil, errRes
	}

	response := types.Response{
		Data: movie,
	}

	return response, nil
}

func (h *Handler) Create(ctx *gofr.Context) (interface{}, error) {
	var mov models.Movie
	err := ctx.Bind(&mov)

	if err != nil {
		errRes := &errors.Response{
			StatusCode: http.StatusBadRequest,
			Code:       "ERROR",
			Reason:     "wrong request body",
		}

		return nil, errRes
	}

	data, deleted, err := h.svc.Post(ctx, &mov)

	if err != nil || deleted {
		errRes := &errors.Response{
			StatusCode: http.StatusBadRequest,
			Code:       "ERROR",
			Reason:     "bad request",
		}

		return nil, errRes
	}

	response := types.Response{
		Data: data,
	}

	return response, nil
}

func (h *Handler) Update(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	movieID, err := strconv.Atoi(id)

	if err != nil {
		errRes := &errors.Response{
			StatusCode: http.StatusBadRequest,
			Code:       "ERROR",
			Reason:     "param id is invalid",
		}

		return nil, errRes
	}

	var updMov models.UpdateMovie
	err = ctx.Bind(&updMov)

	if err != nil {
		errRes := &errors.Response{
			StatusCode: http.StatusBadRequest,
			Code:       "ERROR",
			Reason:     "wrong request body",
		}

		return nil, errRes
	}

	data, deleted, err := h.svc.Update(ctx, movieID, &updMov)

	if err != nil || deleted {
		errRes := &errors.Response{
			StatusCode: http.StatusBadRequest,
			Code:       "ERROR",
			Reason:     "error from the service layer",
		}

		return nil, errRes
	}

	response := types.Response{
		Data: data,
	}

	return response, nil
}

func (h *Handler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	movieID, err := strconv.Atoi(id)

	if err != nil {
		errRes := &errors.Response{
			StatusCode: http.StatusBadRequest,
			Code:       "ERROR",
			Reason:     "param id is invalid",
		}

		return nil, errRes
	}

	deleted, err := h.svc.Delete(ctx, movieID)

	if err != nil || !deleted {
		errRes := &errors.Response{
			StatusCode: http.StatusBadRequest,
			Code:       "ERROR",
			Reason:     "requested id is not present",
		}

		return nil, errRes
	}

	response := types.Response{
		Data: http.StatusNoContent,
	}

	return response, nil
}
