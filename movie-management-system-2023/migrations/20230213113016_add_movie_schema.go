package migrations

import (
	"fmt"

	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/log"
)

type K20230213113016 struct {
}

func (k K20230213113016) Up(d *datastore.DataStore, logger log.Logger) error {
	fmt.Println("Running migration up: add_movie_schema.go")

	_, err := d.DB().Exec("CREATE TABLE IF NOT EXISTS sample5(" +
		"id INT NOT NULL AUTO_INCREMENT," +
		"name  VARCHAR(10) NOT NULL," +
		"genre VARCHAR(15) NOT NULL," +
		"rating FLOAT NOT NULL," +
		"releasedDate DATE NOT NULL," +
		"updatedAt TIMESTAMP NOT NULL ," +
		"createdAt TIMESTAMP NOT NULL ," +
		"plot varchar(100) NOT NULL," +
		"released TINYINT NOT NULL," +
		"deletedAt DATETIME ," +
		"primary key (id));")

	if err != nil {
		return err
	}

	return nil
}

func (k K20230213113016) Down(d *datastore.DataStore, logger log.Logger) error {
	fmt.Println("Running migration down:  add_movie_schema.go")

	_, err := d.DB().Exec("Drop table If EXISTS sample5;")

	if err != nil {
		return err
	}

	return nil
}
