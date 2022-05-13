package store

import (
	"github.com/eternalq/project-server/internal/api/models"
	"github.com/jmoiron/sqlx"
)

type AlbumRepository struct {
	DB sqlx.DB
}

const (
	CREATE_ALBUM = `select create_album(:name, :created_at, :user_id)`
	DELETE_ALBUM = `select delete_album(?)`
	FIND_ALBUM   = `select find_album(?)`
	ADD_PHOTO    = `select add_photo(?, ?)`
	REMOVE_PHOTO = `select remove_photo(?, ?)`
)

func (r *AlbumRepository) Create(alb *models.Album) error {
	rows, err := r.DB.NamedQuery(CREATE_ALBUM, alb)
	if err != nil {
		return err
	}

	return rows.StructScan(alb)
}

func (r *AlbumRepository) Delete(id int) ([]models.Album, error) {
	mm := []models.Album{}
	if err := r.DB.Select(&mm, DELETE_ALBUM, id); err != nil {
		return nil, err
	}

	return mm, nil
}

func (r *AlbumRepository) FindByID(id int) (*models.Album, error) {
	m := &models.Album{}
	if err := r.DB.Get(m, FIND_ALBUM, id); err != nil {
		return nil, err
	}

	return m, nil
}

func (r *AlbumRepository) AddPhoto(albumID int, photoURL string) (*models.Album, error) {
	m := &models.Album{}
	if err := r.DB.Get(m, ADD_PHOTO, albumID, photoURL); err != nil {
		return nil, err
	}

	return m, nil
}

func (r *AlbumRepository) RemovePhoto(albumID int, photoURL string) (*models.Album, error) {
	m := &models.Album{}
	if err := r.DB.Get(m, REMOVE_PHOTO, albumID, photoURL); err != nil {
		return nil, err
	}

	return m, nil
}
