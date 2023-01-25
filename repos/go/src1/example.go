package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "get login")
	case "POST":
		fmt.Fprintln(w, "post login")
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe("localhost:8080", nil)
}
