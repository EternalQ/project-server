package models

import "time"

type Album struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	PhotosURL []string  `json:"photos_url" db:"photos_url"`
	UserID    int       `json:"user_id" db:"user_id"`
}
