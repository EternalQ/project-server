package store

import (
	"github.com/eternalq/project-server/internal/api/models"
	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	DB sqlx.DB
}

const (
	CREATE_POST   = `SELECT * from create_post($1, $2, $3, $4);`
	DELETE_POST   = `SELECT * from delete_post($1);`
	FIND_POST     = `SELECT * from find_post($1);`        //by tag
	GET_POSTS     = `SELECT * from get_posts($1, $2);`    //page size and number
	ADD_POST_TAGS = `SELECT * from add_post_tags($1, $2)` //post id and tags string
)

// also inserts post tags
func (r *PostRepository) Create(p *models.Post) error {
	if err := r.DB.Get(p, CREATE_POST, p.Text, p.CreatedAt, p.PhotoURL, p.UserID); err != nil {
		return err
	}

	if _, err := r.DB.Exec(ADD_POST_TAGS, p.ID, p.TagsCSV); err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) Delete(id int) ([]models.Post, error) {
	pp := []models.Post{}
	if err := r.DB.Select(&pp, DELETE_POST, id); err != nil {
		return nil, err
	}

	return pp, nil
}

func (r *PostRepository) FindByTag(tag string) ([]models.Post, error) {
	pp := []models.Post{}
	if err := r.DB.Select(&pp, FIND_POST, tag); err != nil {
		return nil, err
	}

	return pp, nil
}

func (r *PostRepository) GetLasts(pageSize, pageNum int) ([]models.Post, error) {
	pp := []models.Post{}
	if err := r.DB.Select(&pp, GET_POSTS, pageSize, (pageNum-1)*pageSize); err != nil {
		return nil, err
	}

	return pp, nil
}
