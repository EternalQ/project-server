package models

import "time"

type Post struct {
	ID        int
	Text      string
	CreatedAt time.Time `json:"created_at"`
	PhotoURL  string    `json:"photo_url"`
	User      User
	Comments  []Comment
}
