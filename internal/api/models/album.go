package models

import "time"

type Album struct {
	ID        int
	Name      string
	CreatedAt time.Time
	PhotosURL []string
}
