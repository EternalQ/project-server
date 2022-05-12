package store

import (
	"github.com/eternalq/project-server/internal/api/models"
	"github.com/jmoiron/sqlx"
)

type CommentRepository struct {
	DB sqlx.DB
}

func (r *CommentRepository) Create(com *models.Comment) (*models.Comment, error) {
	return nil, nil
}

func (r *CommentRepository) FindByPostID(id int) ([]models.Comment, error) {
return nil, nil
}
