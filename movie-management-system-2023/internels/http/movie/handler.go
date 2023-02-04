package movie

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-training/movie-management-system-2023/internels/filters"
	"github.com/go-training/movie-management-system-2023/internels/services"
	"github.com/go-training/movie-management-system-2023/models"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type handler struct {
	svc services.MovieManager
}

func New(svc services.MovieManager) *handler {
	return &handler{svc: svc}
}

type response1 struct {
	Code   int     `json:"code"`
	Status string  `json:"status"`
	Data   details `json:"data"`
}

type details struct {
	Movies models.Movie `json:"movies"`
}

type response2 struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type response3 struct {
	Code   int      `json:"code"`
	Status string   `json:"status"`
	Data   details2 `json:"data"`
}

type details2 struct {
	Movies []models.Movie `json:"movies"`
}

func (h *handler) GetAllMovies(w http.ResponseWriter, r *http.Request) {

	ctx, _ := context.WithTimeout(r.Context(), time.Second)

	movie, err := h.svc.GetAll(ctx)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := response2{400, "ERROR", "ERROR message"}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := response3{
		Code:   200,
		Status: "Success",
		Data: details2{
			Movies: movie,
		},
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetMovie(w http.ResponseWriter, r *http.Request) {
	str := strings.Split(r.URL.Path, "/")
	mID := str[len(str)-1]

	movieId, err := strconv.Atoi(mID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := response2{400, "ERROR", "ERROR message"}
		json.NewEncoder(w).Encode(response)
		return
	}

	ctx, _ := context.WithTimeout(r.Context(), time.Second)

	movie, deleted, err := h.svc.Get(ctx, movieId)

	if err != nil || deleted == true {
		w.WriteHeader(http.StatusInternalServerError)
		response := response2{400, "ERROR", "ERROR message"}
		json.NewEncoder(w).Encode(response)
		return
	}
	mov := movie[0]

	response := response1{
		Code:   200,
		Status: "Success",
		Data: details{
			Movies: mov,
		},
	}

	json.NewEncoder(w).Encode(response)
}

func (h *handler) PostMovie(w http.ResponseWriter, r *http.Request) {

	reqbody, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := response2{400, "ERROR", "ERROR message"}
		json.NewEncoder(w).Encode(response)
		return
	}

	var m filters.Details

	json.Unmarshal(reqbody, &m)

	ctx, _ := context.WithTimeout(r.Context(), time.Second)

	data, deleted, err := h.svc.Post(ctx, m)
	fmt.Println(deleted, data)
	if err != nil || deleted == true {
		w.WriteHeader(http.StatusBadRequest)
		response := response2{400, "ERROR", "ERROR message"}
		json.NewEncoder(w).Encode(response)
		return
	}
	mov := data[0]
	response := response1{
		Code:   200,
		Status: "Success",
		Data: details{
			Movies: mov,
		},
	}
	//json.NewEncoder(w).Encode(response)
	movie, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := response2{400, "ERROR", "ERROR message"}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Write(movie)

}

func (h *handler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	str := strings.Split(r.URL.Path, "/")
	mId := str[len(str)-1]
	movieId, err := strconv.Atoi(mId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := response2{400, "ERROR", "ERROR message"}
		json.NewEncoder(w).Encode(response)
		return
	}
	req, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := response2{400, "ERROR", "ERROR message"}
		json.NewEncoder(w).Encode(response)
		return
	}

	var needToUp filters.UpdateMovie

	json.Unmarshal(req, &needToUp)

	ctx, _ := context.WithTimeout(r.Context(), time.Second)

	data, deleted, err := h.svc.Update(ctx, movieId, needToUp)
	if err != nil || deleted == true {
		w.WriteHeader(http.StatusInternalServerError)
		response := response2{400, "ERROR", "ERROR message"}
		json.NewEncoder(w).Encode(response)
		return
	}
	mov := data[0]
	response := response1{
		Code:   200,
		Status: "Success",
		Data: details{
			Movies: mov,
		},
	}
	json.NewEncoder(w).Encode(response)

}

func (h *handler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	str := strings.Split(r.URL.Path, "/")
	mId := str[len(str)-1]
	movieId, err := strconv.Atoi(mId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"param movieId is not int"}`))
		return
	}
	ctx, _ := context.WithTimeout(r.Context(), time.Second)
	data, err := h.svc.Delete(ctx, movieId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := response2{400, "ERROR", "ERROR message"}
		json.NewEncoder(w).Encode(response)
		return
	}
	message := "Movie deleted successfully."
	if data == false {
		w.WriteHeader(http.StatusInternalServerError)
		response := response2{400, "ERROR", "ERROR message"}
		json.NewEncoder(w).Encode(response)
		return
	}
	response := response2{
		Code:    200,
		Status:  "Success",
		Message: message,
	}
	json.NewEncoder(w).Encode(response)

}
