package models

import "time"

type Comment struct {
	ID        int
	Comment   string
	CreatedAt time.Time `json:"created_at"`
	UserEmail string    `json:"user_email"`
}
