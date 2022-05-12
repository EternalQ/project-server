package store

import (
	"github.com/eternalq/project-server/internal/api/models"
	"github.com/jmoiron/sqlx"
)

type AlbumRepository struct {
	DB sqlx.DB
}

func (r *AlbumRepository) Create(alb *models.Album) (*models.Album, error) {
	return nil, nil
}

func (r *AlbumRepository) Delete(id int) error {
	return nil
}

func (r *AlbumRepository) FindByID(id int) (*models.Album, error) {
	return nil, nil
}

func (r *AlbumRepository) AddPhoto(albumID int, photoURL string) error {
	return nil
}

func (r *AlbumRepository) RemovePhoto(albumID int, photoURL string) error {
	return nil
}
