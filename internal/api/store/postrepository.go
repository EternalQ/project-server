package store

import "github.com/jmoiron/sqlx"

type PostRepository struct {
	DB sqlx.DB
}

// also insert post tags
func (r *PostRepository) Create() {

}

func (r *PostRepository) Delete() {

}

func (r *PostRepository) FindByTag() {

}

func (r *PostRepository) GetLast20() {

}
