package models

import "time"

type Post struct {
	ID            int       `json:"id" db:"id"`
	Text          string    `json:"text" db:"text"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	PhotoURL      string    `json:"photo_url" db:"photo_url"`
	TagsCSV       string    `json:"tags_str" db:"tags_str"`
	UserEmail     string    `json:"user_email" db:"user_email"`
	UserID        int       `json:"user_id" db:"user_id"`
	CommentsCount int       `json:"comments_count" db:"comments_count"`
}
