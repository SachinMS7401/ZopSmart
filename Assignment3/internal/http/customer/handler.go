package customer

import (
	"Project/internal/models"
	"Project/internal/services"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type handler struct {
	svc services.Customer
}

func New(svc services.Customer) *handler {
	return &handler{svc: svc}
}

func (h *handler) GET(w http.ResponseWriter, r *http.Request) {
	customerId := 0
	if v := r.URL.Query().Get("Id"); v != "" {
		id, err := strconv.Atoi(v)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"param customerId is not int"}`))
			return
		}
		customerId = id
	}

	ctx, _ := context.WithTimeout(r.Context(), time.Second)

	customers, err := h.svc.GetCustomer(ctx, customerId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}

	response := struct {
		code   int
		status string
		Data   struct {
			Customers []models.Customer `json:"customers"`
		} `json:"data"`
	}{
		code:   200,
		status: "Success",
		Data: struct {
			Customers []models.Customer `json:"customers"`
		}{
			Customers: customers,
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

	var c models.Customer

	json.Unmarshal(reqbody, &c)

	ctx, _ := context.WithTimeout(req.Context(), time.Second)

	fmt.Println("Reached here ")
	data, _ := h.svc.PostNewCustomer(ctx, c)

	bytes, _ := json.Marshal(data)

	w.Write(bytes)

}
