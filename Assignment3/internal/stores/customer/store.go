package customer

import (
	"Project/internal/filters"
	"Project/internal/models"
	"context"
	"database/sql"
	"fmt"
)

const tableName = "customer_details"

type store struct {
	db *sql.DB
}

func New(db *sql.DB) *store {
	return &store{db: db}
}

func (s *store) GetCustomers(ctx context.Context, f filters.Customer) ([]models.Customer, error) {

	var details []interface{}

	if f.Id > 0 {
		details = append(details, f.Id)
	}

	query := "select * from " + tableName + " where ID = ?"

	rows, err := s.db.QueryContext(ctx, query, details...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var c models.Customer

		err := rows.Scan(&c.Id, &c.Name, &c.Email, &c.PhoneNumber)
		if err != nil {
			return nil, err
		}

		customers = append(customers, c)

	}
	fmt.Println("Data in customers is ", customers)

	return customers, nil
}

func (s *store) PostNewCustomer(ctx context.Context, o models.Customer) (*models.Customer, error) {

	var details []interface{}

	query := "INSERT INTO customer_details VALUES(?,?,?,?)"
	details = append(details, o.Id, o.Name, o.Email, o.PhoneNumber)

	row, err := s.db.ExecContext(ctx, query, details...)

	if err != nil {
		fmt.Println("Error while inserting into database")
		return nil, err
	}

	n, _ := row.RowsAffected()

	if n > 0 {
		fmt.Println(n, " number of rows affected in the database")
	}
	var c models.Customer

	c.Id = o.Id
	c.Name = o.Name
	c.Email = o.Email
	c.PhoneNumber = o.PhoneNumber

	fmt.Println(c)

	return &c, nil

}
