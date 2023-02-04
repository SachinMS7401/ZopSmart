package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	handler "github.com/go-training/movie-management-system-2023/internels/http/movie"
	svc "github.com/go-training/movie-management-system-2023/internels/services/movie"
	store "github.com/go-training/movie-management-system-2023/internels/stores/movie"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

//const (
//	username = "root"
//	password = "Shivu@7401"
//	host     = "127.0.0.1"
//	port     = "3306"
//	dbname   = "assignment4"
//)

//func dsn(dbname string) string {
//	return fmt.Sprintf("%s:%s@tcp(%s)/%sparseTime=true", username, password, host, dbname)
//}
//
//func ConnectToMySql() (*sql.DB, error) {
//	db, err := sql.Open("mysql", dsn(dbname))
//	if err != nil {
//		return nil, errors.New("database is not connecting")
////	}
//	fmt.Println(db)
//	db.SetMaxOpenConns(20)
//	db.SetMaxIdleConns(20)
//	db.SetConnMaxLifetime(time.Minute * 5)

//	return db, err
//}

func main() {
	db, err := sql.Open("mysql", "root:Shivu@7401@tcp(127.0.0.1:3306)/assignment4?parseTime=true")
	//db, err := ConnectToMySql()
	if err != nil {
		fmt.Errorf("error %v connecting to the database", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)
	fmt.Println("Connected to Database")
	store := store.New(db)
	svc := svc.New(store)
	handler := handler.New(svc)

	r := mux.NewRouter()
	r.HandleFunc("/movies", handler.GetAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handler.GetMovie).Methods("GET")
	r.HandleFunc("/movies", handler.PostMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handler.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handler.DeleteMovie).Methods("DELETE")

	http.ListenAndServe("localhost:8080", r)
}
