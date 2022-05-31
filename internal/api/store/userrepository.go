package store

import (
	"github.com/eternalq/project-server/internal/api/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB sqlx.DB
}

const (
	CREATE_USER = `select * from create_user($1, $2, $3)`
	FIND_USER   = `select * from find_user($1)`
	ALL_USERS   = `select * from all_users()`
	USER_ALBUMS = `select * from user_albums($1)`
	USER_POSTS  = `select * from user_posts($1)`
)

func (r *UserRepository) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforCreate(); err != nil {
		return err
	}

	return r.DB.Get(u, CREATE_USER, u.Email, u.EncryptedPassword, u.CreatedAt)
}

func (r *UserRepository) Find(id int) (*models.User, error) {
	u := &models.User{}

	if err := r.DB.Get(u, FIND_USER, id); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}

	if err := r.DB.Get(u, FIND_USER, email); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	uu := []models.User{}

	if err := r.DB.Select(&uu, ALL_USERS); err != nil {
		return nil, err
	}

	return uu, nil
}

func (r *UserRepository) GetPosts(id int) ([]models.Post, error) {
	pp := []models.Post{}

	if err := r.DB.Select(&pp, USER_POSTS, id); err != nil {
		return nil, err
	}

	return pp, nil
}

func (r *UserRepository) GetAlbums(id int) ([]models.Album, error) {
	aa := []models.Album{}

	if err := r.DB.Select(&aa, USER_ALBUMS, id); err != nil {
		return nil, err
	}

	return aa, nil
}
