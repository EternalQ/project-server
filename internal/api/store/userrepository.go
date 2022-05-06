package store

import (
	"database/sql"
	"errors"

	"github.com/eternalq/project-server/internal/api/models"
)

type UserRepository struct {
	store *Store
}

//TODO: change to procedure calling
func (r *UserRepository) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"insert into users (email, encrypted_password) values ($1, $2) returning id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

//TODO: change to procedure calling
func (r *UserRepository) Find(id int) (*models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow(
		"select id, email, encrypted_password from users where id = $1",
		id,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err != nil {
			return nil, errors.New("record not found")
		}

		return nil, err
	}

	return u, nil
}

//TODO: change to procedure calling
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow(
		"select id, email, encrypted_password from users where email = $1",
		email,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record not found")
		}

		return nil, err
	}

	return u, nil
}
