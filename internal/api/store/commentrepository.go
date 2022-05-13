package store

import (
	"github.com/eternalq/project-server/internal/api/models"
	"github.com/jmoiron/sqlx"
)

type CommentRepository struct {
	DB sqlx.DB
}

const (
	CREATE_COMMENT = `select create_comment(:comment, :created_at, :post_id, :user_id)`
	FIND_COMMENTS  = `select find_comments(?)`
)

func (r *CommentRepository) Create(c *models.Comment) error {
	rows, err := r.DB.NamedQuery(CREATE_COMMENT, c)
	if err != nil {
		return err
	}

	return rows.StructScan(c)
}

func (r *CommentRepository) FindByPostID(id int) ([]models.Comment, error) {
	cc := []models.Comment{}
	if err := r.DB.Select(cc, FIND_COMMENTS, id); err != nil {
		return nil, err
	}

	return cc, nil
}
