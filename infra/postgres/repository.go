package postgres

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() *Repository {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return &Repository{db: db}
}
