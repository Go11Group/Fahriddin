package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Baza ulanishi uchun konfiguratsiya
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0412"
	dbname   = "nt"
)

// PostgreSQL bazasiga ulanish
func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}
