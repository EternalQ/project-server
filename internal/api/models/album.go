package models

import "time"

type Album struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	PhotosURL []string  `json:"photos_url"`
	UserID    int       `json:"user_id"`
}
