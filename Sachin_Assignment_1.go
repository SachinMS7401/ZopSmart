package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Customers struct {
	Id          int      `json:"Id"`
	Name        string   `json:"Name"`
	Email       string   `json:"Email"`
	PhoneNumber string   `json:"PhoneNumber"`
	Address     *Address `json:"Address"`
	Orders      []Orders `json:"Orders"`
}

type Address struct {
	Street      string `json:"Street"`
	City        string `json:"City"`
	State       string `json:"State"`
	Zipcode     int    `json:"Zipcode"`
	CountryCode string `json:"CountryCode"`
}

type Orders struct {
	OrderId         int  `json:"OrderId"`
	OrderAmount     int  `json:"OrderAmount"`
	TotalItems      int  `json:"TotalItems"`
	IncludesAlcohol bool `json:"IncludesAlcohol"`
}

var c1 = Customers{
	Id:          1,
	Name:        "Sachin",
	Email:       "sachin@123",
	PhoneNumber: "1234567890",
	Address: &Address{
		Street:      "10th cross 21st main Hsr sector 1",
		City:        "Bangalore",
		State:       "Karnataka",
		Zipcode:     123,
		CountryCode: "IND",
	},
	Orders: []Orders{
		{
			OrderId:         123,
			OrderAmount:     40000,
			TotalItems:      12,
			IncludesAlcohol: false,
		},
	},
}

var c2 = Customers{
	Id:          2,
	Name:        "Tan",
	Email:       "tanu@123",
	PhoneNumber: "2345678901",
	Address: &Address{
		Street:      "11th cross 10th main HSR sector 2",
		City:        "Bangalore",
		State:       "Karnataka",
		Zipcode:     345,
		CountryCode: "IND",
	},
	Orders: []Orders{
		{
			OrderId:         345,
			OrderAmount:     7665,
			TotalItems:      21,
			IncludesAlcohol: false,
		},
	},
}

var cust = []Customers{c1, c2}

type results struct {
	Code   int       `json:"Code"`
	Status string    `json:"Status"`
	Data   Customers `json:"Data"`
}

type results2 struct {
	Code    int    `json:"Code"`
	Status  string `json:"Status"`
	Message string `json:"Message"`
}
type results3 struct {
	Code   int         `json:"Code"`
	Status string      `json:"Status"`
	Data   []Customers `json:"Data"`
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("status", "SUCCESS")
	w.WriteHeader(200)
	//json.NewEncoder(w).Encode(cust)
	res1 := results3{200, "success", cust}
	res, _ := json.Marshal(res1)
	w.Write(res)
}

func PostCustomer(w http.ResponseWriter, r *http.Request) {
	resp, _ := ioutil.ReadAll(r.Body)

	var customer Customers
	err := json.Unmarshal(resp, &customer)
	if err != nil {
		return
	}

	for _, cu := range cust {
		if cu.Id == customer.Id {
			w.Header().Add("status", "ERROR")
			w.WriteHeader(400)
			res1 := results2{400, "ERROR", "Customer id already present"}
			res, _ := json.Marshal(res1)
			w.Write(res)
			return
		}
	}

	for _, customer1 := range cust {
		fmt.Println(customer1)
		for _, customer2 := range customer1.Orders {
			for _, k := range customer.Orders {
				if customer2.OrderId == k.OrderId {
					w.Header().Add("status", "ERROR")
					w.WriteHeader(400)
					res1 := results2{400, "ERROR", "Order id already present"}
					res, _ := json.Marshal(res1)
					w.Write(res)
					return
				}
			}
		}
	}

	cust = append(cust, customer)
	w.Header().Add("status", "SUCCESS")
	w.WriteHeader(200)
	res1 := results{200, "success", customer}
	res, _ := json.Marshal(res1)
	w.Write(res)

}

func PathCustomer(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	str := strings.Split(r.URL.Path, "/")
	id := str[len(str)-1]
	id1, _ := strconv.Atoi(id)
	c := 0
	for _, cu := range cust {
		if cu.Id == id1 {
			fmt.Println(id1)
			w.Header().Add("status", "SUCCESS")
			w.WriteHeader(200)
			res1 := results{200, "success", cu}
			res, _ := json.Marshal(res1)
			w.Write(res)
			c += 1
			break
		}
	}
	if c == 0 {
		w.Header().Add("status", "ERROR")
		w.WriteHeader(400)
		res1 := results2{400, "ERROR", "ERROR message"}
		res, _ := json.Marshal(res1)
		w.Write(res)
	}
}

func QueryCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	c := 0
	id1, _ := strconv.Atoi(id)
	for _, cu1 := range cust {

		if cu1.Id == id1 && cu1.Name == name {
			fmt.Println(id1, name)
			w.Header().Add("status", "SUCCESS")
			w.WriteHeader(200)
			//json.NewEncoder(w).Encode(cu1)
			res1 := results{200, "success", cu1}
			res, _ := json.Marshal(res1)
			w.Write(res)
			c += 1
			break
		}
	}
	if c == 0 {
		w.Header().Add("status", "ERROR")
		w.WriteHeader(400)
		res1 := results2{400, "ERROR", "ERROR message"}
		res, _ := json.Marshal(res1)
		w.Write(res)
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getcustomer", GetCustomer)
	r.HandleFunc("/postcustomer", PostCustomer).Methods("POST")
	r.HandleFunc("/getcustomer/{id}", PathCustomer)
	r.HandleFunc("/querycustomer", QueryCustomer)
	err := http.ListenAndServe("localhost:8081", r)
	if err != nil {
		return
	}

}
