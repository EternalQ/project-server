package models

import "time"

type Comment struct {
	ID        int
	Comment   string
	CreatedAt time.Time
	User      User
}
