package server

import (
	"github.com/eternalq/project-server/internal/api/store"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	dburl = "host=localhost dbname=projectdb sslmode=disable user=postgres password=docker"
)

func Start() error {
	db, err := newDB(dburl)
	if err != nil {
		return err
	}
	defer db.Close()

	store := store.New(db)
	server := newServer(store)

	if err := server.router.Run(":8081"); err != nil {
		return err
	}

	return nil
}

func newDB(dburl string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dburl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
