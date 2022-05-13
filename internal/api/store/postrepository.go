package store

import (
	"github.com/eternalq/project-server/internal/api/models"
	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	DB sqlx.DB
}

const (
	CREATE_POST   = `SELECT create_post(:text, :created_at, :photo_url, :user_id);`
	DELETE_POST   = `SELECT delete_post(?);`
	FIND_POST     = `SELECT find_post(?);`              //by tag
	GET_POSTS     = `SELECT get_posts(:id, :tags_str);` //page size and number
	ADD_POST_TAGS = `SELECT add_post_tags(?, ?)`        //post id and tags string
)

// also insert post tags
func (r *PostRepository) Create(p *models.Post) error {
	rows, err := r.DB.NamedQuery(CREATE_POST, p)
	if err != nil {
		return err
	}

	if _, err := r.DB.NamedExec(ADD_POST_TAGS, p); err != nil {
		return err
	}

	return rows.StructScan(p)
}

func (r *PostRepository) Delete(id int) ([]models.Post, error) {
	pp := []models.Post{}
	if err := r.DB.Select(pp, DELETE_POST, id); err != nil {
		return nil, err
	}

	return pp, nil
}

func (r *PostRepository) FindByTag(tag string) (*models.Post, error) {
	p := &models.Post{}
	if err := r.DB.Get(p, FIND_POST, tag); err != nil {
		return nil, err
	}

	return p, nil
}

func (r *PostRepository) GetLast20(pageSize, pageNum int) ([]models.Post, error) {
	pp := []models.Post{}
	if err := r.DB.Select(pp, GET_POSTS, pageSize, pageNum); err != nil {
		return nil, err
	}

	return pp, nil
}
