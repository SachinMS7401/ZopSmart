package main

import (
	http1 "Project/internal/http/customer"
	svc "Project/internal/services/customer"
	store "Project/internal/stores/customer"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	db, err := ConnectToDB()
	if err != nil {
		fmt.Println("Error Connecting to the database", err)
		return
	}
	fmt.Println("Connected To Database ")

	store := store.New(db)
	svc := svc.New(store)
	handler := http1.New(svc)

	r := mux.NewRouter()
	r.HandleFunc("/customers", handler.GET).Methods("GET")
	r.HandleFunc("/customers", handler.POST).Methods("POST")

	http.ListenAndServe("localhost:8088", r)
}

const (
	username = "root"
	password = "Shivu@7401"
	host     = "127.0.0.1"
	port     = "3306"
	dbname   = "assignment3"
)

func dsn(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbname)
}
func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn("assignment3"))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)
	return db, err
}
