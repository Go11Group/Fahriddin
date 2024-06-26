package postgres

import (
	"database/sql"
	_"github.com/lib/pq"
)
func ConnectDB() (*sql.DB,error){
	connStr := "host=localhost port=5432 user=postgres dbname=nt password=0412 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
