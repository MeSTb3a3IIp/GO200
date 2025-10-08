package apiserver

import (
	"apiserver/internal/store"
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

func Start() error {
	db, err := NewDB("user=postgres password=2205 dbname=GO200 sslmode=disable")
	if err != nil {
		return err
	}
	st := store.New(db)
	srv := NewServer(st)
	return http.ListenAndServe(":9091", srv)
}

func NewDB(nameDB string) (*sql.DB, error) {
	db, err := sql.Open("postgres", nameDB)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
