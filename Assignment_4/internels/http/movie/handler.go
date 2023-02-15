package movie

import (
	"Assignment_4/internels/services"
	"Assignment_4/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"time"
)

type handler struct {
	svc services.MovieServicer
}

func New(svc services.MovieServicer) *handler {
	return &handler{svc: svc}
}

func (h *handler) GET(w http.ResponseWriter, r *http.Request) {
	movieId := 0
	v := mux.Vars(r)
	str := v["Id"]
	if str != "" {
		id, err := strconv.Atoi(str)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"param movieId is not int"}`))
			return
		}
		movieId = id
	}

	ctx, _ := context.WithTimeout(r.Context(), time.Second)

	movie, err := h.svc.Get(ctx, movieId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}

	response := struct {
		code   int
		status string
		Data   struct {
			Movies []models.Movie `json:"movies"`
		} `json:"data"`
	}{
		code:   200,
		status: "Success",
		Data: struct {
			Movies []models.Movie `json:"movies"`
		}{
			Movies: movie,
		},
	}

	json.NewEncoder(w).Encode(response)
}

func (h *handler) POST(w http.ResponseWriter, req *http.Request) {

	reqbody, err := io.ReadAll(req.Body)

	if err != nil {
		fmt.Println("Error in the body ", err)
		return
	}

	var c models.Movie

	json.Unmarshal(reqbody, &c)

	ctx, _ := context.WithTimeout(req.Context(), time.Second)

	fmt.Println("Reached here ")
	data, _ := h.svc.Post(ctx, c)

	bytes, _ := json.Marshal(data)
	w.Write(bytes)

}

func (h *handler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	movieId := 0
	v := mux.Vars(r)
	str := v["Id"]
	if str != "" {
		id, err := strconv.Atoi(str)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"param movieId is not int"}`))
			return
		}
		movieId = id
	}
	ctx, _ := context.WithTimeout(r.Context(), time.Second)

	movies, err := h.svc.Get(ctx, movieId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}

}
