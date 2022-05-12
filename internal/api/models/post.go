package models

import "time"

type Post struct {
	ID            int       `json:"id"`
	Text          string    `json:"text"`
	CreatedAt     time.Time `json:"created_at"`
	PhotoURL      string    `json:"photo_url"`
	TagsCSV       string    `json:"tags_str"`
	UserEmail     string    `json:"user_email"`
	UserID        int       `json:"user_id"`
	CommentsCount int       `json:"comments_count"`
}
