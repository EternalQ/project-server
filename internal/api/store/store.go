package store

import (
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db      *sqlx.DB
	User    *UserRepository
	Post    *PostRepository
	Album   *AlbumRepository
	Comment *CommentRepository
}

func New(db *sqlx.DB) *Store {
	return &Store{
		db:      db,
		User:    &UserRepository{DB: *db},
		Post:    &PostRepository{DB: *db},
		Album:   &AlbumRepository{DB: *db},
		Comment: &CommentRepository{DB: *db},
	}
}
