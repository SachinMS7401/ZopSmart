package Assignment_4

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"time"
)

func main() {
	db, err := ConnectToDB()
	if err != nil {
		fmt.Errorf("error %v connecting to the database", err)
	}
	fmt.Println("Connected to Database")

	r := mux.NewRouter()
	r.HandleFunc("/movies", Get).Methods("GET")
}

const (
	username = "root"
	password = "Shivu@7401"
	host     = "127.0.0.1"
	port     = "3306"
	dbname   = "assignment4"
)

func dsn(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbname)
}

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(dbname))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)
	return db, err
}
