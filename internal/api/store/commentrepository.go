package store

import (
	"github.com/eternalq/project-server/internal/api/models"
	"github.com/jmoiron/sqlx"
)

type CommentRepository struct {
	DB sqlx.DB
}

const (
	CREATE_COMMENT = `select * from create_comment($1, $2, $3, $4)`
	FIND_COMMENTS  = `select * from find_comments($1)`
)

func (r *CommentRepository) Create(c *models.Comment) error {
	if err := r.DB.Get(c, CREATE_COMMENT, c.Comment, c.CreatedAt, c.PostID, c.UserID); err != nil {
		return err
	}

	return nil
}

func (r *CommentRepository) FindByPostID(id int) ([]models.Comment, error) {
	cc := []models.Comment{}
	if err := r.DB.Select(&cc, FIND_COMMENTS, id); err != nil {
		return nil, err
	}

	return cc, nil
}
