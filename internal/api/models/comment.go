package models

import "time"

type Comment struct {
	ID        int       `json:"id" db:"id"`
	Comment   string    `json:"comment" db:"comment"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UserEmail string    `json:"user_email" db:"user_email"`
	PostID    int       `json:"post_id" db:"post_id"`
	UserID    int       `json:"user_id" db:"user_id"`
}
