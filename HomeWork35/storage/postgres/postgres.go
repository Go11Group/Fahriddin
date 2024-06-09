package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host = localhost port = 5432 name = postgres dbname = nt password = 0412")
	if err != nil {
		return &sql.DB{}, err
	}
	err = db.Ping()
	if err != nil {
		return &sql.DB{}, err
	}
	return db, nil
}
