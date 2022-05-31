package store

import (
	"strings"

	"github.com/eternalq/project-server/internal/api/models"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type AlbumRepository struct {
	DB sqlx.DB
}

const (
	CREATE_ALBUM = `select * from create_album($1, $2, $3)`
	DELETE_ALBUM = `select * from delete_album($1)`
	FIND_ALBUM   = `select * from find_album($1)`
	ADD_PHOTO    = `select * from add_photo($1, $2)`
	REMOVE_PHOTO = `select * from remove_photo($1, $2)`
)

func (r *AlbumRepository) Create(alb *models.Album) error {
	if err := r.DB.Get(alb, CREATE_ALBUM, alb.Name, alb.CreatedAt, alb.UserID); err != nil {
		return err
	}

	alb.PhotosURL = strings.Split(alb.PhotosStr, ",")

	return nil
}

func (r *AlbumRepository) Delete(id int) ([]models.Album, error) {
	mm := []models.Album{}
	if err := r.DB.Select(&mm, DELETE_ALBUM, id); err != nil {
		return nil, err
	}

	for _, m := range mm {
		m.PhotosURL = strings.Split(m.PhotosStr, ",")
	}

	return mm, nil
}

func (r *AlbumRepository) FindByID(id int) (*models.Album, error) {
	m := &models.Album{}
	if err := r.DB.Get(m, FIND_ALBUM, id); err != nil {
		return nil, err
	}

	m.PhotosURL = strings.Split(m.PhotosStr, ",")

	return m, nil
}

func (r *AlbumRepository) AddPhoto(albumID int, photoURL string) (*models.Album, error) {
	m := &models.Album{}
	if err := r.DB.Get(m, ADD_PHOTO, albumID, photoURL); err != nil {
		return nil, err
	}

	m.PhotosURL = strings.Split(m.PhotosStr, ",")

	return m, nil
}

func (r *AlbumRepository) RemovePhoto(albumID int, photoURL string) (*models.Album, error) {
	m := &models.Album{}
	if err := r.DB.Get(m, REMOVE_PHOTO, albumID, photoURL); err != nil {
		return nil, err
	}

	m.PhotosURL = strings.Split(m.PhotosStr, ",")

	return m, nil
}
