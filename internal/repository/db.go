package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresConnection() *sql.DB {
	connStr := "host=my-postgres port=5432 user=postgres password=postgres dbname=url-shortener sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to connect DB:", err)
	}

	log.Println("Connected to PostgreSQL ✅")

	return db
}
